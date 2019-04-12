package service

import (
	"errors"
	"io/ioutil"
	"os"

	v1 "github.com/VideoCoin/cloud-api/notifications/v1"
	"gopkg.in/yaml.v2"
)

var (
	ErrTemplateNotFound = errors.New("template not found")
)

type Template struct {
	Subject string `yaml:"subject"`
	Content string `yaml:"content"`
	Type    string `yaml:"type"`
}

type TemplateStore struct {
	Email map[string]Template `yaml:"email"`
	Web   map[string]Template `yaml:"web"`
}

func NewTemplateStore(path string) (*TemplateStore, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	store := &TemplateStore{}

	err = yaml.Unmarshal(b, store)
	if err != nil {
		return nil, err
	}

	return store, nil
}

func (s *TemplateStore) GetTemplate(t v1.NotificationTarget, name string) (*Template, error) {
	switch t {
	case v1.NotificationTarget_EMAIL:
		template, ok := s.Email[name]
		if !ok {
			return nil, ErrTemplateNotFound
		}
		return &template, nil
	case v1.NotificationTarget_WEB:
		template, ok := s.Web[name]
		if !ok {
			return nil, ErrTemplateNotFound
		}
		return &template, nil
	default:
		return nil, ErrTemplateNotFound
	}
}
