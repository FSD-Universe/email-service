// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

// Package global
package global

import "flag"

var (
	DownloadPrefix = flag.String("download_prefix", "https://raw.githubusercontent.com/FSD-Universe/email-service/refs/heads/main", "auto download prefix")
)

const (
	AppVersion    = "0.1.0"
	ConfigVersion = "0.1.0"

	ServiceName = "email-service"

	EnvDownloadPrefix = "DOWNLOAD_PREFIX"
)
