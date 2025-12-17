// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

// Package controller
package controller

import (
	DTO "email-service/src/interfaces/server/dto"
	"email-service/src/interfaces/server/service"

	"github.com/labstack/echo/v4"
	"half-nothing.cn/service-core/interfaces/http/dto"
	"half-nothing.cn/service-core/interfaces/logger"
)

type EmailController struct {
	logger  logger.Interface
	service service.EmailInterface
}

func NewEmailController(
	lg logger.Interface,
	service service.EmailInterface,
) *EmailController {
	return &EmailController{
		logger:  logger.NewLoggerAdapter(lg, "email-controller"),
		service: service,
	}
}

func (controller *EmailController) SendEmailCode(ctx echo.Context) error {
	data := &DTO.SendEmailCode{}
	if err := ctx.Bind(data); err != nil {
		controller.logger.Errorf("SendEmailCode handle fail, parse argument fail, %v", err)
		return dto.ErrorResponse(ctx, dto.ErrErrorParam)
	}
	controller.logger.Debugf("SendEmailCode with argument %#v", data)
	res, err := dto.ValidStruct(data)
	if err != nil {
		controller.logger.Errorf("SendEmailCode handle fail, validate err, %v", err)
		return dto.ErrorResponse(ctx, dto.ErrServerError)
	}
	if res != nil {
		controller.logger.Errorf("SendEmailCode handle fail, validate argument fail, %v", res)
		return dto.ErrorResponse(ctx, res)
	}
	return controller.service.SendEmailCode(data).Response(ctx)
}
