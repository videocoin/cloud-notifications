package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/centrifugal/gocent"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	v1 "github.com/videocoin/cloud-api/notifications/v1"
	"github.com/videocoin/cloud-pkg/mqmux"
)

var (
	ErrUnknownTarget    = errors.New("unknown target")
	ErrUnknownRecipient = errors.New("unknown recipient")
	ErrUnknownEvent     = errors.New("unknown event")
)

type CoreOption struct {
	FromEmail      string
	InternalEmails []string
	Env            string
	Logger         *logrus.Entry
}

type Core struct {
	opts   *CoreOption
	logger *logrus.Entry

	mq    *mqmux.WorkerMux
	store *TemplateStore

	email *sendgrid.Client
	cent  *gocent.Client
}

func NewCore(mq *mqmux.WorkerMux, store *TemplateStore, opts *CoreOption) (*Core, error) {
	emailCli := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))

	centCli := gocent.New(gocent.Config{
		Addr: os.Getenv("CENT_API_ADDR"),
		Key:  os.Getenv("CENT_API_KEY"),
	})

	return &Core{
		opts:   opts,
		logger: opts.Logger,
		mq:     mq,
		store:  store,
		email:  emailCli,
		cent:   centCli,
	}, nil
}

func (c *Core) Start() error {
	err := c.mq.Consumer("notifications.send", 5, false, c.performMessage)
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

	var span opentracing.Span
	tracer := opentracing.GlobalTracer()
	spanCtx, err := tracer.Extract(opentracing.TextMap, mqmux.RMQHeaderCarrier(msg.Headers))

	if err != nil {
		span = tracer.StartSpan("performMessage")
	} else {
		span = tracer.StartSpan("performMessage", ext.RPCServerOption(spanCtx))
	}

	defer span.Finish()

	notification := &v1.Notification{}
	err = json.Unmarshal(msg.Body, notification)
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

	n.Params["subject"] = nt.Subject
	if c.opts.Env == "everest" {
		n.Params["domain"] = fmt.Sprintf("console.videocoin.network")
	} else {
		n.Params["domain"] = fmt.Sprintf("console.%s.videocoin.network", c.opts.Env)
	}

	html, err := c.store.renderTemplate(n.Template, n.Params)
	if err != nil {
		logger.Error(err)
		return err
	}

	_, ok := n.Params["internal"]
	if ok {
		if c.opts.Env != "dev" && c.opts.Env != "snb" {
			for _, to := range c.opts.InternalEmails {
				err = c.sendEmail(to, nt.Subject, html)
				if err != nil {
					logger.Error(err)
					return err
				}
			}
		}
	} else {
		to, ok := n.Params["to"]
		if !ok {
			return ErrUnknownRecipient
		}

		err = c.sendEmail(to, nt.Subject, html)
		if err != nil {
			logger.Error(err)
			return err
		}
	}

	return nil
}

func (c *Core) sendEmail(to, subject, content string) error {
	c.logger.WithField("to", to).Info("sending email")

	message := mail.NewSingleEmail(
		mail.NewEmail("VideoCoin", c.opts.FromEmail), subject, mail.NewEmail("", to), " ", content)
	resp, err := c.email.Send(message)
	if err != nil {
		c.logger.Error(err)
		return err
	}

	if resp.StatusCode >= 400 && resp.StatusCode < 600 {
		respErr := fmt.Errorf("failed to send email: %s", resp.Body)
		c.logger.Error(respErr)
		return respErr
	}

	return nil
}

func (c *Core) performWebNotification(n *v1.Notification) error {
	logger := c.logger.WithFields(logrus.Fields{
		"target": n.Target.String(),
	})

	toChannel := fmt.Sprintf("users#%s", n.Params["user_id"])

	logger.WithField("to channel", toChannel).WithField("params", n.Params).Info("sending push")

	payload, err := json.Marshal(n.Params)
	if err != nil {
		logger.Error(err)
		return err
	}

	err = c.cent.Publish(context.Background(), toChannel, payload)
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}
