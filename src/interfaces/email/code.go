// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

// Package email
package email

import (
	"errors"
	"time"
)

type CodeData struct {
	Cid  int
	Code string
}

var (
	ErrEmailCodeCooldown = errors.New("email code cool down")
	ErrEmailCodeExpired  = errors.New("email code expired")
	ErrEmailCodeInvalid  = errors.New("email code invalid")
)

type CodeManagerInterface interface {
	GenerateEmailCode(target string, cid int) (*VerifyCodeEmail, time.Duration, error)
	VerifyEmailCode(target string, cid int, code string) error
}
