package service

import (
	"github.com/sirupsen/logrus"
)

type Config struct {
	Name    string `envconfig:"-"`
	Version string `envconfig:"-"`

	Addr      string `default:"0.0.0.0:5020"`
	MQURI     string `default:"amqp://guest:guest@127.0.0.1:5672"`
	Templates string `default:"./templates.yaml"`

	FromEmail string `default:"support@videocoin.network"`

	Logger *logrus.Entry `envconfig:"-"`
}
