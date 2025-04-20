package service

import "tdidf/internal/config"

type ServiceImpl struct {
	config *config.Config
}

type Option func(*ServiceImpl)

func New(conf *config.Config, opts ...Option) *ServiceImpl {
	s := ServiceImpl{
		config: conf,
	}
	for _, opt := range opts {
		opt(&s)
	}
	return &s
}
