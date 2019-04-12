package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"text/template"
	"time"

	v1 "github.com/VideoCoin/cloud-api/notifications/v1"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

var (
	ErrUnknownTarget    = errors.New("unknown target")
	ErrUnknownRecipient = errors.New("unknown recipient")
)

type CoreOption struct {
	FromEmail string
	Logger    *logrus.Entry
}

type Core struct {
	opts   *CoreOption
	logger *logrus.Entry

	conn  *amqp.Connection
	q     *amqp.Queue
	ch    *amqp.Channel
	m     <-chan amqp.Delivery
	store *TemplateStore

	email *sendgrid.Client
}

func NewCore(uri string, store *TemplateStore, opts *CoreOption) (*Core, error) {
	hn, _ := os.Hostname()

	emailCli := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))

	config := amqp.Config{
		Heartbeat: 10 * time.Second,
		Locale:    "en_US",
		Properties: amqp.Table{
			"connection_name": hn,
		},
	}

	conn, err := amqp.DialConfig(uri, config)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	q, err := ch.QueueDeclare(
		"notifications",
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return nil, err
	}

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	if err != nil {
		return nil, err
	}

	messages, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		return nil, err
	}

	return &Core{
		opts:   opts,
		logger: opts.Logger,
		conn:   conn,
		q:      &q,
		ch:     ch,
		m:      messages,
		store:  store,
		email:  emailCli,
	}, nil
}

func (c *Core) Start() error {
	for msg := range c.m {
		err := c.performMessage(msg)
		if err != nil {
			msg.Reject(false)
		}
		msg.Ack(false)
	}

	return nil
}

func (c *Core) Stop() error {
	c.ch.Close()
	c.conn.Close()
	return nil
}

func (c *Core) performMessage(msg amqp.Delivery) error {
	c.logger.Debugf("received a message: %s", msg.Body)

	notification := &v1.Notification{}
	err := json.Unmarshal(msg.Body, notification)
	if err != nil {
		c.logger.Error(err)
		return err
	}

	err = c.performNotification(notification)
	if err != nil {
		c.logger.Error(err)
		return err
	}

	return nil
}

func (c *Core) performNotification(n *v1.Notification) error {
	switch n.GetTarget() {
	case v1.NotificationTarget_EMAIL:
		return c.performEmailNotification(n)
	default:
		return ErrUnknownTarget
	}
}

func (c *Core) performEmailNotification(n *v1.Notification) error {
	logger := c.logger.WithFields(logrus.Fields{
		"target":   n.Target.String(),
		"template": n.Template,
	})

	nt, err := c.store.GetTemplate(v1.NotificationTarget_EMAIL, n.Template)
	if err != nil {
		logger.Error(err)
		return err
	}

	toEmail, ok := n.Params["to"]
	if !ok {
		return ErrUnknownRecipient
	}

	buf := bytes.NewBuffer(nil)
	err = applyTemplate(buf, n.Template, nt.Content, n.Params)
	if err != nil {
		logger.Error(err)
		return err
	}

	logger.WithField("to", toEmail).Info("sending email")

	from := mail.NewEmail("", c.opts.FromEmail)
	to := mail.NewEmail("", toEmail)
	message := mail.NewSingleEmail(from, nt.Subject, to, " ", buf.String())
	resp, err := c.email.Send(message)
	if err != nil {
		logger.Error(err)
		return err
	}

	if resp.StatusCode >= 400 && resp.StatusCode < 600 {
		respErr := fmt.Errorf("failed to send email: %s", resp.Body)
		logger.Error(respErr)
		return respErr
	}

	return nil
}

func applyTemplate(wr io.Writer, name, content string, params interface{}) error {
	t, err := template.New(name).Parse(content)
	if err != nil {
		return err
	}

	err = t.Execute(wr, params)
	if err != nil {
		return err
	}

	return nil
}
