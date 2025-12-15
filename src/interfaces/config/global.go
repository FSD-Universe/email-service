// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

// Package config
package config

import (
	"email-service/src/interfaces/global"

	"half-nothing.cn/service-core/interfaces/config"
)

type GlobalConfig struct {
	*config.GlobalConfig `yaml:",inline"`
}

func (g *GlobalConfig) InitDefaults() {
	g.GlobalConfig = &config.GlobalConfig{}
	g.GlobalConfig.InitDefaults()
	g.Name = "email-service"
	g.Version = global.ConfigVersion
	g.ConfigVersion = global.ConfigVersion
}
