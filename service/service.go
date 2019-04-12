package service

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

	core, err := NewCore(cfg.MQURI, store, coreOpts)
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
