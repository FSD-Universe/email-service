// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

// Package grpc
package grpc

import (
	pb "email-service/src/interfaces/grpc"

	"half-nothing.cn/service-core/interfaces/logger"
)

type EmailServer struct {
	pb.UnimplementedEmailServer
	logger logger.Interface
}

func NewEmailServer(lg logger.Interface) *EmailServer {
	return &EmailServer{
		logger: logger.NewLoggerAdapter(lg, "grpc-server"),
	}
}
