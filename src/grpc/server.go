// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

// Package grpc
package grpc

import (
	"email-service/src/interfaces/content"
	"email-service/src/interfaces/email"
	pb "email-service/src/interfaces/grpc"

	"half-nothing.cn/service-core/interfaces/logger"
)

type EmailServer struct {
	pb.UnimplementedEmailServer
	logger  logger.Interface
	sender  email.SenderInterface
	manager email.CodeManagerInterface
}

func NewEmailServer(content *content.ApplicationContent) *EmailServer {
	return &EmailServer{
		logger:  logger.NewLoggerAdapter(content.Logger(), "grpc-server"),
		sender:  content.EmailSender(),
		manager: content.CodeManager(),
	}
}
