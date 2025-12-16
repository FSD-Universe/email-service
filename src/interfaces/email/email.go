// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

// Package email
package email

import (
	"email-service/src/interfaces/config"
	"errors"
)

var (
	ErrEmailNotRegistered = errors.New("email not registered")
	ErrEmailNotEnabled    = errors.New("email not enabled")
	ErrEmailDataInvalid   = errors.New("email data invalid")
)

type SenderInterface interface {
	SendEmail(emailType config.Email, target string, data interface{}) error
}

type DataValidator func(data interface{}) bool

type ActivityAtcJoinEmail struct {
	Cid          string
	ActivityName string
	ActivityTime string
	Facility     string
	Frequency    string
}

type ActivityAtcLeaveEmail struct {
	Cid          string
	ActivityName string
}

type ActivityPilotJoinEmail struct {
	Cid          string
	ActivityName string
	ActivityTime string
	Callsign     string
	Aircraft     string
}

type ActivityPilotLeaveEmail struct {
	Cid          string
	ActivityName string
}

type ApplicationPassedEmail struct {
	Cid      string
	Operator string
	Message  string
	Contact  string
}

type ApplicationProcessingEmail struct {
	Cid     string
	Time    string
	Contact string
}

type ApplicationRejectedEmail struct {
	Cid      string
	Operator string
	Reason   string
	Contact  string
}

type AtcRatingChangeEmail struct {
	Cid      string
	NewValue string
	OldValue string
	Operator string
	Contact  string
}

type BannedEmail struct {
	Cid      string
	Reason   string
	Time     string
	Operator string
	Contact  string
}

type InstructorChangeEmail struct {
	Cid        string
	Reason     string
	Instructor string
	Operator   string
	Contact    string
}

type KickedFromServerEmail struct {
	Cid      string
	Reason   string
	Time     string
	Operator string
	Contact  string
}

type PasswordChangeEmail struct {
	Cid       string
	Time      string
	IP        string
	UserAgent string
}

type PasswordResetEmail struct {
	Cid       string
	Time      string
	IP        string
	UserAgent string
}

type PermissionChangeEmail struct {
	Cid         string
	Permissions string
	Operator    string
	Contact     string
}

type RoleChangeEmail struct {
	Cid      string
	Roles    string
	Operator string
	Contact  string
}

type TicketReplyEmail struct {
	Cid   string
	Title string
	Reply string
}

type VerifyCodeEmail struct {
	Cid       string
	Code      string
	ExpiredAt string
	Expired   string
}

type WelcomeEmail struct {
	Cid string
}

var Validators = map[config.Email]DataValidator{
	config.EmailVerifyCode:            func(data interface{}) bool { _, ok := data.(*VerifyCodeEmail); return ok },
	config.EmailWelcome:               func(data interface{}) bool { _, ok := data.(*WelcomeEmail); return ok },
	config.EmailRatingChange:          func(data interface{}) bool { _, ok := data.(*AtcRatingChangeEmail); return ok },
	config.EmailKickedFromServer:      func(data interface{}) bool { _, ok := data.(*KickedFromServerEmail); return ok },
	config.EmailPasswordChange:        func(data interface{}) bool { _, ok := data.(*PasswordChangeEmail); return ok },
	config.EmailPasswordReset:         func(data interface{}) bool { _, ok := data.(*PasswordResetEmail); return ok },
	config.EmailApplicationPassed:     func(data interface{}) bool { _, ok := data.(*ApplicationPassedEmail); return ok },
	config.EmailApplicationRejected:   func(data interface{}) bool { _, ok := data.(*ApplicationRejectedEmail); return ok },
	config.EmailApplicationProcessing: func(data interface{}) bool { _, ok := data.(*ApplicationProcessingEmail); return ok },
	config.EmailTicketReply:           func(data interface{}) bool { _, ok := data.(*TicketReplyEmail); return ok },
	config.EmailActivityPilotJoin:     func(data interface{}) bool { _, ok := data.(*ActivityPilotJoinEmail); return ok },
	config.EmailActivityPilotLeave:    func(data interface{}) bool { _, ok := data.(*ActivityPilotLeaveEmail); return ok },
	config.EmailActivityAtcJoin:       func(data interface{}) bool { _, ok := data.(*ActivityAtcJoinEmail); return ok },
	config.EmailActivityAtcLeave:      func(data interface{}) bool { _, ok := data.(*ActivityAtcLeaveEmail); return ok },
	config.EmailInstructorChange:      func(data interface{}) bool { _, ok := data.(*InstructorChangeEmail); return ok },
	config.EmailBanned:                func(data interface{}) bool { _, ok := data.(*BannedEmail); return ok },
	config.EmailRoleChange:            func(data interface{}) bool { _, ok := data.(*RoleChangeEmail); return ok },
	config.EmailPermissionChange:      func(data interface{}) bool { _, ok := data.(*PermissionChangeEmail); return ok },
}
