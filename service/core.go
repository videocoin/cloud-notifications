package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"text/template"

	v1 "github.com/VideoCoin/cloud-api/notifications/v1"
	"github.com/VideoCoin/cloud-pkg/mqmux"
	"github.com/pusher/pusher-http-go"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

var (
	ErrUnknownTarget    = errors.New("unknown target")
	ErrUnknownRecipient = errors.New("unknown recipient")
	ErrUnknownEvent     = errors.New("unknown event")
)

type CoreOption struct {
	FromEmail string
	Logger    *logrus.Entry
}

type Core struct {
	opts   *CoreOption
	logger *logrus.Entry

	mq    *mqmux.WorkerMux
	store *TemplateStore

	email  *sendgrid.Client
	pusher *pusher.Client
}

func NewCore(mq *mqmux.WorkerMux, store *TemplateStore, opts *CoreOption) (*Core, error) {
	emailCli := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	pucherCli := &pusher.Client{
		AppId:   os.Getenv("PUSHER_API_APPID"),
		Key:     os.Getenv("PUSHER_API_KEY"),
		Secret:  os.Getenv("PUSHER_API_SECRET"),
		Cluster: "us2",
		Secure:  true,
	}

	return &Core{
		opts:   opts,
		logger: opts.Logger,
		mq:     mq,
		store:  store,
		email:  emailCli,
		pusher: pucherCli,
	}, nil
}

func (c *Core) Start() error {
	err := c.mq.Consumer("notifications/send", 5, false, c.performMessage)
	if err != nil {
		return err
	}
	return c.mq.Run()
}

func (c *Core) Stop() error {
	return c.mq.Close()
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
	case v1.NotificationTarget_WEB:
		return c.performWebNotification(n)
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

func (c *Core) performWebNotification(n *v1.Notification) error {
	logger := c.logger.WithFields(logrus.Fields{
		"target": n.Target.String(),
	})

	toChannel := fmt.Sprintf("user-%s", n.Params["user_id"])

	event, ok := n.Params["event"]
	if !ok {
		return ErrUnknownEvent
	}

	logger.WithField("to channel", toChannel).WithField("event", event).Info("sending push")

	_, err := c.pusher.Trigger(toChannel, event, n.Params)

	if err != nil {
		logger.Error(err)
		return err
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
