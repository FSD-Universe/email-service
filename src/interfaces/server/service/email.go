// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

// Package service
package service

import (
	DTO "email-service/src/interfaces/server/dto"

	"half-nothing.cn/service-core/interfaces/http/dto"
)

var (
	ErrSendEmailCode = dto.NewApiStatus("EMAIL_SEND_FAILED", "邮件发送失败", dto.HttpCodeInternalError)
)

type EmailInterface interface {
	SendEmailCode(form *DTO.SendEmailCode) *dto.ApiResponse[DTO.SendEmailCodeResponse]
}
