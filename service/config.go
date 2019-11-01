package service

import (
	"github.com/sirupsen/logrus"
)

type Config struct {
	Name    string `envconfig:"-"`
	Version string `envconfig:"-"`

	MQURI     string `default:"amqp://guest:guest@127.0.0.1:5672" envconfig:"MQURI"`
	Templates string `default:"templates" envconfig:"TEMPLATES"`
	RPCAddr   string `required:"true" envconfig:"RPC_ADDR" default:"127.0.0.1:5005"`
	FromEmail string `default:"support@videocoin.network"`

	Logger *logrus.Entry `envconfig:"-"`
}
