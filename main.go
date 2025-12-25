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

	defer cl.Clean()

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

	emailSender := email.NewSender(lg, applicationConfig.EmailConfig)
	emailManager := email.NewCodeManager(lg, applicationConfig.EmailConfig, codeCache, sendCache)

	contentBuilder := content.NewApplicationContentBuilder().
		SetConfigManager(configManager).
		SetCleaner(cl).
		SetLogger(lg).
		SetEmailSender(emailSender).
		SetCodeManager(emailManager)

	started := make(chan bool)
	initFunc := func(s *grpc.Server) {
		grpcServer := grpcImpl.NewEmailServer(lg, emailSender, emailManager)
		pb.RegisterEmailServer(s, grpcServer)
	}
	if applicationConfig.TelemetryConfig.Enable && applicationConfig.TelemetryConfig.GrpcServerTrace {
		go grpcUtils.StartGrpcServerWithTrace(lg, cl, applicationConfig.ServerConfig.GrpcServerConfig, started, initFunc)
	} else {
		go grpcUtils.StartGrpcServer(lg, cl, applicationConfig.ServerConfig.GrpcServerConfig, started, initFunc)
	}

	consulClient := discovery.NewConsulClient(lg, applicationConfig.GlobalConfig.Discovery, g.AppVersion)

	if err := consulClient.RegisterServer(); err != nil {
		lg.Fatalf("fail to register server: %v", err)
		return
	}

	cl.Add("Discovery", consulClient.UnregisterServer)

	go func() {
		for {
			<-consulClient.EventChan
		}
	}()

	go server.StartHttpServer(contentBuilder.Build())

	cl.Wait()
}
