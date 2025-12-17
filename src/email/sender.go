// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

// Package email
package email

import (
	"email-service/src/interfaces/config"
	"email-service/src/interfaces/email"
	"html/template"
	"strings"

	"gopkg.in/gomail.v2"
	"half-nothing.cn/service-core/interfaces/logger"
)

type Sender struct {
	logger    logger.Interface
	config    *config.EmailConfig
	templates *config.TemplateConfig
}

func NewSender(
	lg logger.Interface,
	c *config.EmailConfig,
) *Sender {
	sender := &Sender{
		logger:    logger.NewLoggerAdapter(lg, "email-sender"),
		config:    c,
		templates: c.Template.Templates,
	}
	return sender
}

func (sender *Sender) SendEmail(emailType config.Email, target string, data interface{}) error {
	if !emailType.Data.Enable {
		return email.ErrEmailNotEnabled
	}
	validator, exist := email.Validators[emailType]
	if !exist {
		return email.ErrEmailNotRegistered
	}
	if !validator(data) {
		return email.ErrEmailDataInvalid
	}

	target = strings.ToLower(target)

	m, err := sender.generateEmail(target, emailType, data)
	if err != nil {
		sender.logger.Errorf("failed to generate %s email: %s", emailType.Value, err.Error())
		return err
	}

	sender.logger.Infof("sending %s email to %s with args: %#v", emailType.Value, target, data)

	if err := gomail.Send(sender.config.Closer, m); err != nil {
		sender.logger.Errorf("failed to send %s email: %s", emailType.Value, err.Error())
		return err
	}
	return nil
}

func (sender *Sender) renderTemplate(template *template.Template, data interface{}) (string, error) {
	var sb strings.Builder
	if err := template.Execute(&sb, data); err != nil {
		return "", err
	}
	return sb.String(), nil
}

func (sender *Sender) generateEmail(email string, emailType config.Email, data interface{}) (*gomail.Message, error) {
	content, err := sender.renderTemplate(emailType.Data.Template, data)
	if err != nil {
		return nil, err
	}

	m := gomail.NewMessage()
	m.SetHeader("From", sender.config.Username)
	m.SetHeader("To", email)
	m.SetHeader("Subject", emailType.Data.Subject)
	m.SetBody("text/html", content)

	return m, nil
}
