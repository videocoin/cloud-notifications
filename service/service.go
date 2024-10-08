package service

import (
	"strings"

	"github.com/videocoin/cloud-pkg/mqmux"
)

type Service struct {
	cfg  *Config
	core *Core
	rpc  *RPCServer
}

func NewService(cfg *Config) (*Service, error) {
	store, err := NewTemplateStore(cfg.Templates)
	if err != nil {
		return nil, err
	}

	coreOpts := &CoreOption{
		FromEmail:      cfg.FromEmail,
		InternalEmails: strings.Fields(cfg.InternalEmails),
		Logger:         cfg.Logger,
		Env:            cfg.Env,
	}

	mq, err := mqmux.NewWorkerMux(cfg.MQURI, cfg.Name)
	if err != nil {
		return nil, err
	}
	mq.Logger = cfg.Logger.WithField("system", "mq")

	core, err := NewCore(mq, store, coreOpts)
	if err != nil {
		return nil, err
	}

	rpcConfig := &RPCServerOpts{
		Logger: cfg.Logger,
		Addr:   cfg.RPCAddr,
	}

	rpc, err := NewRPCServer(rpcConfig)
	if err != nil {
		return nil, err
	}

	svc := &Service{
		cfg:  cfg,
		rpc:  rpc,
		core: core,
	}

	return svc, nil
}

func (s *Service) Start(errCh chan error) {
	go func() {
		errCh <- s.rpc.Start()
	}()

	err := s.core.Start()
	if err != nil {
		errCh <- err
	}
}

func (s *Service) Stop() error {
	err := s.core.Stop()
	return err
}
