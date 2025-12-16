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

// ApplicationContent 应用程序上下文结构体，包含所有核心组件的接口
type ApplicationContent struct {
	configManager config.ManagerInterface[*c.Config] // 配置管理器
	cleaner       cleaner.Interface                  // 清理器
	logger        logger.Interface                   // 日志
	emailSender   email.SenderInterface              // 邮件发送器
	codeManager   email.CodeManagerInterface         // 邮件验证码管理器
}

func (app *ApplicationContent) ConfigManager() config.ManagerInterface[*c.Config] {
	return app.configManager
}

func (app *ApplicationContent) Cleaner() cleaner.Interface { return app.cleaner }

func (app *ApplicationContent) Logger() logger.Interface { return app.logger }

func (app *ApplicationContent) EmailSender() email.SenderInterface { return app.emailSender }

func (app *ApplicationContent) CodeManager() email.CodeManagerInterface { return app.codeManager }
