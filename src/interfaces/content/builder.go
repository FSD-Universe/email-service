// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

// Package content
package content

import (
	c "email-service/src/interfaces/config"
	"email-service/src/interfaces/email"

	"half-nothing.cn/service-core/interfaces/cleaner"
	"half-nothing.cn/service-core/interfaces/config"
	"half-nothing.cn/service-core/interfaces/logger"
)

type ApplicationContentBuilder struct {
	content *ApplicationContent
}

func NewApplicationContentBuilder() *ApplicationContentBuilder {
	return &ApplicationContentBuilder{
		content: &ApplicationContent{},
	}
}

func (builder *ApplicationContentBuilder) SetConfigManager(configManager config.ManagerInterface[*c.Config]) *ApplicationContentBuilder {
	builder.content.configManager = configManager
	return builder
}

func (builder *ApplicationContentBuilder) SetCleaner(cleaner cleaner.Interface) *ApplicationContentBuilder {
	builder.content.cleaner = cleaner
	return builder
}

func (builder *ApplicationContentBuilder) SetLogger(logger logger.Interface) *ApplicationContentBuilder {
	builder.content.logger = logger
	return builder
}

func (builder *ApplicationContentBuilder) SetEmailSender(emailSender email.SenderInterface) *ApplicationContentBuilder {
	builder.content.emailSender = emailSender
	return builder
}

func (builder *ApplicationContentBuilder) SetCodeManager(codeManager email.CodeManagerInterface) *ApplicationContentBuilder {
	builder.content.codeManager = codeManager
	return builder
}
func (builder *ApplicationContentBuilder) Build() *ApplicationContent {
	return builder.content
}
