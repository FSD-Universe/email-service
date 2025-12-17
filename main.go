// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

// Package main
package main

import (
	"context"
	"email-service/src/email"
	grpcImpl "email-service/src/grpc"
	c "email-service/src/interfaces/config"
	"email-service/src/interfaces/content"
	e "email-service/src/interfaces/email"
	g "email-service/src/interfaces/global"
	pb "email-service/src/interfaces/grpc"
	"email-service/src/server"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"half-nothing.cn/service-core/cache"
	"half-nothing.cn/service-core/cleaner"
	"half-nothing.cn/service-core/config"
	"half-nothing.cn/service-core/discovery"
	grpcUtils "half-nothing.cn/service-core/grpc"
	"half-nothing.cn/service-core/interfaces/global"
	"half-nothing.cn/service-core/logger"
	"half-nothing.cn/service-core/telemetry"
	"half-nothing.cn/service-core/utils"
)

func main() {
	global.CheckFlags()

	utils.CheckStringEnv(g.EnvDownloadPrefix, g.DownloadPrefix)

	configManager := config.NewManager[*c.Config]()
	if err := configManager.Init(); err != nil {
		fmt.Printf("fail to initialize configuration file: %v", err)
		return
	}

	applicationConfig := configManager.GetConfig()
	lg := logger.NewLogger()
	lg.Init(
		global.LogName,
		applicationConfig.GlobalConfig.LogConfig,
	)

	lg.Info(" _____           _ _ _____             _")
	lg.Info("|   __|_____ ___|_| |   __|___ ___ _ _|_|___ ___")
	lg.Info("|   __|     | .'| | |__   | -_|  _| | | |  _| -_|")
	lg.Info("|_____|_|_|_|__,|_|_|_____|___|_|  \\_/|_|___|___|")
	lg.Info(fmt.Sprintf("%49s", fmt.Sprintf("Copyright © %d-%d Half_nothing", global.BeginYear, time.Now().Year())))
	lg.Info(fmt.Sprintf("%49s", fmt.Sprintf("EmailService v%s", g.AppVersion)))

	cl := cleaner.NewCleaner(lg)
	cl.Init()

	cl.Add("EmailSender", func(_ context.Context) error {
		return applicationConfig.EmailConfig.Closer.Close()
	})

	if applicationConfig.TelemetryConfig.Enable {
		sdk := telemetry.NewSDK(lg, applicationConfig.TelemetryConfig)
		shutdown, err := sdk.SetupOTelSDK(context.Background())
		if err != nil {
			lg.Fatalf("fail to initialize telemetry: %v", err)
			return
		}
		cl.Add("Telemetry", shutdown)
	}

	codeCache := cache.NewMemoryCache[string, *e.CodeData](applicationConfig.EmailConfig.VerifyExpireDuration)
	sendCache := cache.NewMemoryCache[string, time.Time](applicationConfig.EmailConfig.VerifyIntervalDuration)
	cl.Add("Cache", func(_ context.Context) error { codeCache.Close(); sendCache.Close(); return nil })

	applicationContent := content.NewApplicationContentBuilder().
		SetConfigManager(configManager).
		SetCleaner(cl).
		SetLogger(lg).
		SetEmailSender(email.NewSender(lg, applicationConfig.EmailConfig)).
		SetCodeManager(email.NewCodeManager(lg, applicationConfig.EmailConfig, codeCache, sendCache)).
		Build()

	go server.StartHttpServer(applicationContent)

	if applicationConfig.ServerConfig.GrpcServerConfig.Enable {
		started := make(chan bool)
		go grpcUtils.StartGrpcServer(lg, cl, applicationConfig.ServerConfig.GrpcServerConfig, started, func(s *grpc.Server) {
			grpcServer := grpcImpl.NewEmailServer(applicationContent)
			pb.RegisterEmailServer(s, grpcServer)
		})
		go discovery.StartServiceDiscovery(lg, cl, started, utils.NewVersion(g.AppVersion),
			"email-service", applicationConfig.ServerConfig.GrpcServerConfig.Port)
	}

	cl.Wait()
}
