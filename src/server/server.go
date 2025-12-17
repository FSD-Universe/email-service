// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

// Package server
package server

import (
	"context"
	"email-service/src/interfaces/content"
	"email-service/src/server/controller"
	"email-service/src/server/service"
	"io"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"half-nothing.cn/service-core/http"
	"half-nothing.cn/service-core/interfaces/logger"
)

func StartHttpServer(content *content.ApplicationContent) {
	c := content.ConfigManager().GetConfig()
	lg := logger.NewLoggerAdapter(content.Logger(), "http-server")

	lg.Info("Http server initializing...")
	e := echo.New()
	e.Logger.SetOutput(io.Discard)
	e.Logger.SetLevel(log.OFF)

	http.SetEchoConfig(lg, e, c.ServerConfig.HttpServerConfig, 30*time.Second)

	if c.TelemetryConfig.HttpServerTrace {
		http.SetTelemetry(e, c.TelemetryConfig)
	}

	emailController := controller.NewEmailController(
		lg,
		service.NewEmailService(
			lg,
			content.EmailSender(),
			content.CodeManager(),
		),
	)

	apiGroup := e.Group("/api/v1")
	emailGroup := apiGroup.Group("/emails")
	emailGroup.POST("/code", emailController.SendEmailCode)

	content.Cleaner().Add("http-server", func(ctx context.Context) error {
		return e.Shutdown(ctx)
	})

	http.Serve(lg, e, c.ServerConfig.HttpServerConfig)
}
