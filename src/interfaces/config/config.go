// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

// Package config
package config

import "half-nothing.cn/service-core/interfaces/config"

type Config struct {
	GlobalConfig    *GlobalConfig           `yaml:"global"`
	EmailConfig     *EmailConfig            `yaml:"email"`
	ServerConfig    *config.ServerConfig    `yaml:"server"`
	TelemetryConfig *config.TelemetryConfig `yaml:"telemetry"`
}

func (c *Config) InitDefaults() {
	c.GlobalConfig = &GlobalConfig{}
	c.GlobalConfig.InitDefaults()
	c.EmailConfig = &EmailConfig{}
	c.EmailConfig.InitDefaults()
	c.ServerConfig = &config.ServerConfig{}
	c.ServerConfig.InitDefaults()
	c.TelemetryConfig = &config.TelemetryConfig{}
	c.TelemetryConfig.InitDefaults()
}

func (c *Config) Verify() (bool, error) {
	if ok, err := c.GlobalConfig.Verify(); !ok {
		return ok, err
	}
	if ok, err := c.EmailConfig.Verify(); !ok {
		return ok, err
	}
	if ok, err := c.ServerConfig.Verify(); !ok {
		return ok, err
	}
	if ok, err := c.TelemetryConfig.Verify(); !ok {
		return ok, err
	}
	return true, nil
}
