package service

import "github.com/VideoCoin/cloud-pkg/mqmux"

type Service struct {
	cfg  *Config
	core *Core
}

func NewService(cfg *Config) (*Service, error) {
	store, err := NewTemplateStore(cfg.Templates)
	if err != nil {
		return nil, err
	}

	coreOpts := &CoreOption{
		FromEmail: cfg.FromEmail,
		Logger:    cfg.Logger,
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

	svc := &Service{
		cfg:  cfg,
		core: core,
	}

	return svc, nil
}

func (s *Service) Start() error {
	go s.core.Start()
	return nil
}

func (s *Service) Stop() error {
	s.core.Stop()
	return nil
}
