// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

// Package service
package service

import (
	"email-service/src/interfaces/config"
	"email-service/src/interfaces/email"
	DTO "email-service/src/interfaces/server/dto"
	"email-service/src/interfaces/server/service"
	"errors"
	"fmt"

	"half-nothing.cn/service-core/interfaces/http/dto"
	"half-nothing.cn/service-core/interfaces/logger"
)

type EmailService struct {
	logger  logger.Interface
	sender  email.SenderInterface
	manager email.CodeManagerInterface
}

func NewEmailService(
	lg logger.Interface,
	sender email.SenderInterface,
	manager email.CodeManagerInterface,
) *EmailService {
	return &EmailService{
		logger:  logger.NewLoggerAdapter(lg, "email-service"),
		sender:  sender,
		manager: manager,
	}
}

func (e *EmailService) SendEmailCode(form *DTO.SendEmailCode) *dto.ApiResponse[DTO.SendEmailCodeResponse] {
	emailData, duration, err := e.manager.GenerateEmailCode(form.Email, form.Cid)
	if err != nil {
		if errors.Is(err, email.ErrEmailCodeCooldown) {
			return dto.NewApiResponse[DTO.SendEmailCodeResponse](
				dto.NewApiStatus(
					"EMAIL_SEND_INTERVAL",
					fmt.Sprintf("邮件已发送, 请在%.0f秒后重试", duration.Seconds()),
					dto.HttpCodeBadRequest,
				),
				false,
			)
		}
		return dto.NewApiResponse[DTO.SendEmailCodeResponse](dto.ErrServerError, false)
	}

	err = e.sender.SendEmail(config.EmailVerifyCode, form.Email, emailData)
	if err != nil {
		return dto.NewApiResponse[DTO.SendEmailCodeResponse](service.ErrSendEmailCode, false)
	}
	return dto.NewApiResponse[DTO.SendEmailCodeResponse](dto.SuccessHandleRequest, true)
}
