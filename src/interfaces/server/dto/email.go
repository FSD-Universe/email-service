// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

// Package dto
package dto

type SendEmailCode struct {
	Cid   int    `json:"cid" valid:"required,min=0"`
	Email string `json:"email" valid:"required,regex=^[\\w-]+@[\\w-]+(\\.[\\w-]+)+$"`
}

type SendEmailCodeResponse = bool
