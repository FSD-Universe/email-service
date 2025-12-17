// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

// Package email
package email

import (
	"email-service/src/interfaces/config"
	"email-service/src/interfaces/email"
	"fmt"
	"strings"
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

func (c *CodeManager) GenerateEmailCode(target string, cid int) (*email.VerifyCodeEmail, time.Duration, error) {
	target = strings.ToLower(target)
	if val, ok := c.sendCache.Get(target); ok {
		return nil, val.Add(c.config.VerifyIntervalDuration).Sub(time.Now()), email.ErrEmailCodeCooldown
	}
	code := randstr.String(6)
	now := time.Now()
	expireAt := now.Add(c.config.VerifyExpireDuration)
	c.cache.Set(target, &email.CodeData{Cid: cid, Code: code}, expireAt)
	c.sendCache.SetWithTTL(target, time.Now(), c.config.VerifyIntervalDuration)
	return &email.VerifyCodeEmail{
		Cid:       fmt.Sprintf("%04d", cid),
		Code:      code,
		Expired:   fmt.Sprintf("%.0f", c.config.VerifyExpireDuration.Minutes()),
		ExpiredAt: expireAt.Format(time.RFC3339),
	}, time.Duration(0), nil
}

func (c *CodeManager) VerifyEmailCode(target string, cid int, code string) error {
	target = strings.ToLower(target)
	c.logger.Infof("verifying email code for %s with cid %d and code %s", target, cid, code)
	val, ok := c.cache.Get(target)
	if !ok {
		c.logger.Warnf("email code for %s expired", target)
		return email.ErrEmailCodeExpired
	}
	if val.Code != code || val.Cid != cid {
		c.logger.Warnf("email code for %s invalid", target)
		return email.ErrEmailCodeInvalid
	}
	c.cache.Del(target)
	return nil
}
