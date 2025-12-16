// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

// Package email
package email

import (
	"email-service/src/interfaces/config"
	"email-service/src/interfaces/email"
	"time"

	"github.com/thanhpk/randstr"
	"half-nothing.cn/service-core/interfaces/cache"
	"half-nothing.cn/service-core/interfaces/logger"
)

type CodeManager struct {
	logger    logger.Interface
	config    *config.EmailConfig
	cache     cache.Interface[string, *email.CodeData]
	sendCache cache.Interface[string, time.Time]
}

func NewCodeManager(
	lg logger.Interface,
	config *config.EmailConfig,
	cache cache.Interface[string, *email.CodeData],
	sendCache cache.Interface[string, time.Time],
) *CodeManager {
	return &CodeManager{
		logger:    logger.NewLoggerAdapter(lg, "code-manager"),
		config:    config,
		cache:     cache,
		sendCache: sendCache,
	}
}

func (c *CodeManager) GenerateEmailCode(target string, cid int) (string, time.Duration, error) {
	if val, ok := c.sendCache.Get(target); ok {
		return "", val.Add(c.config.VerifyIntervalDuration).Sub(time.Now()), email.ErrEmailCodeCooldown
	}
	code := randstr.String(6)
	c.cache.SetWithTTL(target, &email.CodeData{Cid: cid, Code: code}, c.config.VerifyExpireDuration)
	c.sendCache.SetWithTTL(target, time.Now(), c.config.VerifyIntervalDuration)
	return code, time.Duration(0), nil
}

func (c *CodeManager) VerifyEmailCode(target string, cid int, code string) error {
	val, ok := c.cache.Get(target)
	if !ok {
		return email.ErrEmailCodeExpired
	}
	if val.Code != code || val.Cid != cid {
		return email.ErrEmailCodeInvalid
	}
	c.cache.Del(target)
	return nil
}
