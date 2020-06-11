package service

import (
	"github.com/sirupsen/logrus"
)

type Config struct {
	Name    string        `envconfig:"-"`
	Version string        `envconfig:"-"`
	Logger  *logrus.Entry `envconfig:"-"`

	MQURI          string `envconfig:"MQURI" default:"amqp://guest:guest@127.0.0.1:5672"`
	Templates      string `envconfig:"TEMPLATES" default:"templates"`
	Env            string `envconfig:"ENV" default:"dev"`
	RPCAddr        string `envconfig:"RPC_ADDR" default:"127.0.0.1:5005"`
	FromEmail      string `default:"support@videocoin.network"`
	InternalEmails string `envconfig:"INTERNAL_EMAILS" default:"dmitry@liveplanet.net adidenko@liveplanet.net"`
}
