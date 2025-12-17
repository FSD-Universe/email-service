// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

// Package email
package email

import (
	"email-service/src/interfaces/config"
	pb "email-service/src/interfaces/grpc"
	"errors"

	"half-nothing.cn/service-core/utils"
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

func (a *ActivityAtcJoinEmail) FromGRPCMessage(message *pb.ActivityAtcJoin) *ActivityAtcJoinEmail {
	if message == nil {
		return a
	}
	a.Cid = utils.GetPointerData(message.Cid)
	a.ActivityName = utils.GetPointerData(message.ActivityName)
	a.ActivityTime = utils.GetPointerData(message.ActivityTime)
	a.Facility = utils.GetPointerData(message.Facility)
	a.Frequency = utils.GetPointerData(message.Frequency)
	return a
}

type ActivityAtcLeaveEmail struct {
	Cid          string
	ActivityName string
}

func (a *ActivityAtcLeaveEmail) FromGRPCMessage(message *pb.ActivityAtcLeave) *ActivityAtcLeaveEmail {
	if message == nil {
		return a
	}
	a.Cid = utils.GetPointerData(message.Cid)
	a.ActivityName = utils.GetPointerData(message.ActivityName)
	return a
}

type ActivityPilotJoinEmail struct {
	Cid          string
	ActivityName string
	ActivityTime string
	Callsign     string
	Aircraft     string
}

func (a *ActivityPilotJoinEmail) FromGRPCMessage(message *pb.ActivityPilotJoin) *ActivityPilotJoinEmail {
	if message == nil {
		return a
	}
	a.Cid = utils.GetPointerData(message.Cid)
	a.ActivityName = utils.GetPointerData(message.ActivityName)
	a.ActivityTime = utils.GetPointerData(message.ActivityTime)
	a.Callsign = utils.GetPointerData(message.Callsign)
	a.Aircraft = utils.GetPointerData(message.Aircraft)
	return a
}

type ActivityPilotLeaveEmail struct {
	Cid          string
	ActivityName string
}

func (a *ActivityPilotLeaveEmail) FromGRPCMessage(message *pb.ActivityPilotLeave) *ActivityPilotLeaveEmail {
	if message == nil {
		return a
	}
	a.Cid = utils.GetPointerData(message.Cid)
	a.ActivityName = utils.GetPointerData(message.ActivityName)
	return a
}

type ApplicationPassedEmail struct {
	Cid      string
	Operator string
	Message  string
	Contact  string
}

func (a *ApplicationPassedEmail) FromGRPCMessage(message *pb.ApplicationPassed) *ApplicationPassedEmail {
	if message == nil {
		return a
	}
	a.Cid = utils.GetPointerData(message.Cid)
	a.Operator = utils.GetPointerData(message.Operator)
	a.Message = utils.GetPointerData(message.Message)
	a.Contact = utils.GetPointerData(message.Contact)
	return a
}

type ApplicationProcessingEmail struct {
	Cid     string
	Time    string
	Contact string
}

func (a *ApplicationProcessingEmail) FromGRPCMessage(message *pb.ApplicationProcessing) *ApplicationProcessingEmail {
	if message == nil {
		return a
	}
	a.Cid = utils.GetPointerData(message.Cid)
	a.Time = utils.GetPointerData(message.Time)
	a.Contact = utils.GetPointerData(message.Contact)
	return a
}

type ApplicationRejectedEmail struct {
	Cid      string
	Operator string
	Reason   string
	Contact  string
}

func (a *ApplicationRejectedEmail) FromGRPCMessage(message *pb.ApplicationRejected) *ApplicationRejectedEmail {
	if message == nil {
		return a
	}
	a.Cid = utils.GetPointerData(message.Cid)
	a.Operator = utils.GetPointerData(message.Operator)
	a.Reason = utils.GetPointerData(message.Reason)
	a.Contact = utils.GetPointerData(message.Contact)
	return a
}

type AtcRatingChangeEmail struct {
	Cid      string
	NewValue string
	OldValue string
	Operator string
	Contact  string
}

func (a *AtcRatingChangeEmail) FromGRPCMessage(message *pb.AtcRatingChange) *AtcRatingChangeEmail {
	if message == nil {
		return a
	}
	a.Cid = utils.GetPointerData(message.Cid)
	a.NewValue = utils.GetPointerData(message.NewValue)
	a.OldValue = utils.GetPointerData(message.OldValue)
	a.Operator = utils.GetPointerData(message.Operator)
	a.Contact = utils.GetPointerData(message.Contact)
	return a
}

type BannedEmail struct {
	Cid      string
	Reason   string
	Time     string
	Operator string
	Contact  string
}

func (b *BannedEmail) FromGRPCMessage(message *pb.Banned) *BannedEmail {
	if message == nil {
		return b
	}
	b.Cid = utils.GetPointerData(message.Cid)
	b.Reason = utils.GetPointerData(message.Reason)
	b.Time = utils.GetPointerData(message.Time)
	b.Operator = utils.GetPointerData(message.Operator)
	b.Contact = utils.GetPointerData(message.Contact)
	return b
}

type InstructorChangeEmail struct {
	Cid        string
	Reason     string
	Instructor string
	Operator   string
	Contact    string
}

func (i *InstructorChangeEmail) FromGRPCMessage(message *pb.InstructorChange) *InstructorChangeEmail {
	if message == nil {
		return i
	}
	i.Cid = utils.GetPointerData(message.Cid)
	i.Reason = utils.GetPointerData(message.Reason)
	i.Instructor = utils.GetPointerData(message.Instructor)
	i.Operator = utils.GetPointerData(message.Operator)
	i.Contact = utils.GetPointerData(message.Contact)
	return i
}

type KickedFromServerEmail struct {
	Cid      string
	Reason   string
	Time     string
	Operator string
	Contact  string
}

func (k *KickedFromServerEmail) FromGRPCMessage(message *pb.KickedFromServer) *KickedFromServerEmail {
	if message == nil {
		return k
	}
	k.Cid = utils.GetPointerData(message.Cid)
	k.Reason = utils.GetPointerData(message.Reason)
	k.Time = utils.GetPointerData(message.Time)
	k.Operator = utils.GetPointerData(message.Operator)
	k.Contact = utils.GetPointerData(message.Contact)
	return k
}

type PasswordChangeEmail struct {
	Cid       string
	Time      string
	IP        string
	UserAgent string
}

func (p *PasswordChangeEmail) FromGRPCMessage(message *pb.PasswordChange) *PasswordChangeEmail {
	if message == nil {
		return p
	}
	p.Cid = utils.GetPointerData(message.Cid)
	p.Time = utils.GetPointerData(message.Time)
	p.IP = utils.GetPointerData(message.Ip)
	p.UserAgent = utils.GetPointerData(message.UserAgent)
	return p
}

type PasswordResetEmail struct {
	Cid       string
	Time      string
	IP        string
	UserAgent string
}

func (p *PasswordResetEmail) FromGRPCMessage(message *pb.PasswordReset) *PasswordResetEmail {
	if message == nil {
		return p
	}
	p.Cid = utils.GetPointerData(message.Cid)
	p.Time = utils.GetPointerData(message.Time)
	p.IP = utils.GetPointerData(message.Ip)
	p.UserAgent = utils.GetPointerData(message.UserAgent)
	return p
}

type PermissionChangeEmail struct {
	Cid         string
	Permissions string
	Operator    string
	Contact     string
}

func (p *PermissionChangeEmail) FromGRPCMessage(message *pb.PermissionChange) *PermissionChangeEmail {
	if message == nil {
		return p
	}
	p.Cid = utils.GetPointerData(message.Cid)
	p.Permissions = utils.GetPointerData(message.Permissions)
	p.Operator = utils.GetPointerData(message.Operator)
	p.Contact = utils.GetPointerData(message.Contact)
	return p
}

type RoleChangeEmail struct {
	Cid      string
	Roles    string
	Operator string
	Contact  string
}

func (r *RoleChangeEmail) FromGRPCMessage(message *pb.RoleChange) *RoleChangeEmail {
	if message == nil {
		return r
	}
	r.Cid = utils.GetPointerData(message.Cid)
	r.Roles = utils.GetPointerData(message.Roles)
	r.Operator = utils.GetPointerData(message.Operator)
	r.Contact = utils.GetPointerData(message.Contact)
	return r
}

type TicketReplyEmail struct {
	Cid   string
	Title string
	Reply string
}

func (t *TicketReplyEmail) FromGRPCMessage(message *pb.TicketReply) *TicketReplyEmail {
	if message == nil {
		return t
	}
	t.Cid = utils.GetPointerData(message.Cid)
	t.Title = utils.GetPointerData(message.Title)
	t.Reply = utils.GetPointerData(message.Reply)
	return t
}

type WelcomeEmail struct {
	Cid string
}

func (w *WelcomeEmail) FromGRPCMessage(message *pb.Welcome) *WelcomeEmail {
	if message == nil {
		return w
	}
	w.Cid = utils.GetPointerData(message.Cid)
	return w
}

type VerifyCodeEmail struct {
	Cid       string
	Code      string
	ExpiredAt string
	Expired   string
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
