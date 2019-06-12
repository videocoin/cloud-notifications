package service

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"

	v1 "github.com/videocoin/cloud-api/notifications/v1"
	"github.com/vanng822/go-premailer/premailer"
	"gopkg.in/yaml.v2"
)

var (
	ErrTemplateNotFound = errors.New("template not found")
)

type Template struct {
	Subject string `yaml:"subject"`
}

type TemplateStore struct {
	Email map[string]Template `yaml:"email"`
	Web   map[string]Template `yaml:"web"`

	path string
}

func NewTemplateStore(path string) (*TemplateStore, error) {
	f, err := os.Open(path + "/templates.yaml")
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

	store.path = path

	return store, nil
}

func (ts *TemplateStore) renderTemplate(name string, params map[string]string) (string, error) {
	t, err := template.ParseFiles(fmt.Sprintf("%s/%s.html", ts.path, name), ts.path+"/style.css")
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	err = t.ExecuteTemplate(&buf, name+".html", params)
	if err != nil {
		return "", err
	}

	prem, err := premailer.NewPremailerFromString(
		buf.String(), premailer.NewOptions())
	if err != nil {
		return "", err
	}

	html, err := prem.Transform()
	if err != nil {
		return "", err
	}

	return html, nil
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
