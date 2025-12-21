// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

// Package grpc
package grpc

import (
	"context"
	"email-service/src/interfaces/config"
	"email-service/src/interfaces/email"
	pb "email-service/src/interfaces/grpc"
	"errors"
	"reflect"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"half-nothing.cn/service-core/interfaces/logger"
)

var (
	SuccessResponse = &pb.SendResponse{Success: true}
	FailedResponse  = &pb.SendResponse{Success: false}
)

type EmailServer struct {
	pb.UnimplementedEmailServer
	logger  logger.Interface
	sender  email.SenderInterface
	manager email.CodeManagerInterface
}

func NewEmailServer(
	lg logger.Interface,
	sender email.SenderInterface,
	manager email.CodeManagerInterface,
) *EmailServer {
	return &EmailServer{
		logger:  logger.NewLoggerAdapter(lg, "grpc-server"),
		sender:  sender,
		manager: manager,
	}
}

func (e *EmailServer) handleSendError(err error) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, email.ErrEmailNotEnabled) {
		return status.Error(codes.Unavailable, "this type of email is not enabled")
	}
	if errors.Is(err, email.ErrEmailNotRegistered) {
		return status.Error(codes.InvalidArgument, "invalid email type")
	}
	if errors.Is(err, email.ErrEmailDataInvalid) {
		return status.Error(codes.Internal, "internal server error")
	}
	return status.Error(codes.Internal, "internal server error")
}

func (e *EmailServer) extractAndValidateFields(msg interface{}) bool {
	val := reflect.ValueOf(msg).Elem()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if field.Kind() == reflect.String {
			if field.IsZero() {
				return false
			}
		}
	}
	return true
}

func (e *EmailServer) sendEmailTemplate(emailType config.Email, targetEmail string, data interface{}) (*pb.SendResponse, error) {
	e.logger.Infof("send %s email to %s with arguments %#v", emailType.Value, targetEmail, data)
	if err := e.sender.SendEmail(emailType, targetEmail, data); err != nil {
		return FailedResponse, e.handleSendError(err)
	}
	return SuccessResponse, nil
}

func (e *EmailServer) SendActivityAtcJoin(_ context.Context, d *pb.ActivityAtcJoin) (*pb.SendResponse, error) {
	if !e.extractAndValidateFields(d) {
		return nil, status.Error(codes.InvalidArgument, "missing required argument")
	}
	data := &email.ActivityAtcJoinEmail{
		Cid:          d.Cid,
		ActivityName: d.ActivityName,
		ActivityTime: d.ActivityTime,
		Facility:     d.Facility,
		Frequency:    d.Frequency,
	}
	return e.sendEmailTemplate(config.EmailActivityAtcJoin, d.TargetEmail, data)
}

func (e *EmailServer) SendActivityAtcLeave(_ context.Context, d *pb.ActivityAtcLeave) (*pb.SendResponse, error) {
	if !e.extractAndValidateFields(d) {
		return nil, status.Error(codes.InvalidArgument, "missing required argument")
	}
	data := &email.ActivityAtcLeaveEmail{
		Cid:          d.Cid,
		ActivityName: d.ActivityName,
	}
	return e.sendEmailTemplate(config.EmailActivityAtcLeave, d.TargetEmail, data)
}

func (e *EmailServer) SendActivityPilotJoin(_ context.Context, d *pb.ActivityPilotJoin) (*pb.SendResponse, error) {
	if !e.extractAndValidateFields(d) {
		return nil, status.Error(codes.InvalidArgument, "missing required argument")
	}
	data := &email.ActivityPilotJoinEmail{
		Cid:          d.Cid,
		ActivityName: d.ActivityName,
		ActivityTime: d.ActivityTime,
		Aircraft:     d.Aircraft,
		Callsign:     d.Callsign,
	}
	return e.sendEmailTemplate(config.EmailActivityPilotJoin, d.TargetEmail, data)
}

func (e *EmailServer) SendActivityPilotLeave(_ context.Context, d *pb.ActivityPilotLeave) (*pb.SendResponse, error) {
	if !e.extractAndValidateFields(d) {
		return nil, status.Error(codes.InvalidArgument, "missing required argument")
	}
	data := &email.ActivityPilotLeaveEmail{
		Cid:          d.Cid,
		ActivityName: d.ActivityName,
	}
	return e.sendEmailTemplate(config.EmailActivityPilotLeave, d.TargetEmail, data)
}

func (e *EmailServer) SendApplicationPassed(_ context.Context, d *pb.ApplicationPassed) (*pb.SendResponse, error) {
	if !e.extractAndValidateFields(d) {
		return nil, status.Error(codes.InvalidArgument, "missing required argument")
	}
	data := &email.ApplicationPassedEmail{
		Cid:      d.Cid,
		Contact:  d.Contact,
		Message:  d.Message,
		Operator: d.Operator,
	}
	return e.sendEmailTemplate(config.EmailApplicationPassed, d.TargetEmail, data)
}

func (e *EmailServer) SendApplicationProcessing(_ context.Context, d *pb.ApplicationProcessing) (*pb.SendResponse, error) {
	if !e.extractAndValidateFields(d) {
		return nil, status.Error(codes.InvalidArgument, "missing required argument")
	}
	data := &email.ApplicationProcessingEmail{
		Cid:     d.Cid,
		Contact: d.Contact,
		Time:    d.Time,
	}
	return e.sendEmailTemplate(config.EmailApplicationProcessing, d.TargetEmail, data)
}

func (e *EmailServer) SendApplicationRejected(_ context.Context, d *pb.ApplicationRejected) (*pb.SendResponse, error) {
	if !e.extractAndValidateFields(d) {
		return nil, status.Error(codes.InvalidArgument, "missing required argument")
	}
	data := &email.ApplicationRejectedEmail{
		Cid:      d.Cid,
		Contact:  d.Contact,
		Operator: d.Operator,
		Reason:   d.Reason,
	}
	return e.sendEmailTemplate(config.EmailApplicationRejected, d.TargetEmail, data)
}

func (e *EmailServer) SendAtcRatingChange(_ context.Context, d *pb.AtcRatingChange) (*pb.SendResponse, error) {
	if !e.extractAndValidateFields(d) {
		return nil, status.Error(codes.InvalidArgument, "missing required argument")
	}
	data := &email.AtcRatingChangeEmail{
		Cid:      d.Cid,
		Contact:  d.Contact,
		NewValue: d.NewValue,
		OldValue: d.OldValue,
		Operator: d.Operator,
	}
	return e.sendEmailTemplate(config.EmailRatingChange, d.TargetEmail, data)
}

func (e *EmailServer) SendBanned(_ context.Context, d *pb.Banned) (*pb.SendResponse, error) {
	if !e.extractAndValidateFields(d) {
		return nil, status.Error(codes.InvalidArgument, "missing required argument")
	}
	data := &email.BannedEmail{
		Cid:      d.Cid,
		Contact:  d.Contact,
		Operator: d.Operator,
		Reason:   d.Reason,
		Time:     d.Time,
	}
	return e.sendEmailTemplate(config.EmailBanned, d.TargetEmail, data)
}

func (e *EmailServer) SendInstructorChange(_ context.Context, d *pb.InstructorChange) (*pb.SendResponse, error) {
	if !e.extractAndValidateFields(d) {
		return nil, status.Error(codes.InvalidArgument, "missing required argument")
	}
	data := &email.InstructorChangeEmail{
		Cid:        d.Cid,
		Contact:    d.Contact,
		Reason:     d.Reason,
		Instructor: d.Instructor,
		Operator:   d.Operator,
	}
	return e.sendEmailTemplate(config.EmailInstructorChange, d.TargetEmail, data)
}

func (e *EmailServer) SendKickedFromServer(_ context.Context, d *pb.KickedFromServer) (*pb.SendResponse, error) {
	if !e.extractAndValidateFields(d) {
		return nil, status.Error(codes.InvalidArgument, "missing required argument")
	}
	data := &email.KickedFromServerEmail{
		Cid:      d.Cid,
		Contact:  d.Contact,
		Operator: d.Operator,
		Reason:   d.Reason,
		Time:     d.Time,
	}
	return e.sendEmailTemplate(config.EmailKickedFromServer, d.TargetEmail, data)
}

func (e *EmailServer) SendPasswordChange(_ context.Context, d *pb.PasswordChange) (*pb.SendResponse, error) {
	if !e.extractAndValidateFields(d) {
		return nil, status.Error(codes.InvalidArgument, "missing required argument")
	}
	data := &email.PasswordChangeEmail{
		Cid:       d.Cid,
		Time:      d.Time,
		IP:        d.Ip,
		UserAgent: d.UserAgent,
	}
	return e.sendEmailTemplate(config.EmailPasswordChange, d.TargetEmail, data)
}

func (e *EmailServer) SendPasswordReset(_ context.Context, d *pb.PasswordReset) (*pb.SendResponse, error) {
	if !e.extractAndValidateFields(d) {
		return nil, status.Error(codes.InvalidArgument, "missing required argument")
	}
	data := &email.PasswordResetEmail{
		Cid:       d.Cid,
		Time:      d.Time,
		IP:        d.Ip,
		UserAgent: d.UserAgent,
	}
	return e.sendEmailTemplate(config.EmailPasswordReset, d.TargetEmail, data)
}

func (e *EmailServer) SendPermissionChange(_ context.Context, d *pb.PermissionChange) (*pb.SendResponse, error) {
	if !e.extractAndValidateFields(d) {
		return nil, status.Error(codes.InvalidArgument, "missing required argument")
	}
	data := &email.PermissionChangeEmail{
		Cid:         d.Cid,
		Permissions: d.Permissions,
		Operator:    d.Operator,
		Contact:     d.Contact,
	}
	return e.sendEmailTemplate(config.EmailPermissionChange, d.TargetEmail, data)
}

func (e *EmailServer) SendRoleChange(_ context.Context, d *pb.RoleChange) (*pb.SendResponse, error) {
	if !e.extractAndValidateFields(d) {
		return nil, status.Error(codes.InvalidArgument, "missing required argument")
	}
	data := &email.RoleChangeEmail{
		Cid:      d.Cid,
		Roles:    d.Roles,
		Operator: d.Operator,
		Contact:  d.Contact,
	}
	return e.sendEmailTemplate(config.EmailRoleChange, d.TargetEmail, data)
}

func (e *EmailServer) SendTicketReply(_ context.Context, d *pb.TicketReply) (*pb.SendResponse, error) {
	if !e.extractAndValidateFields(d) {
		return nil, status.Error(codes.InvalidArgument, "missing required argument")
	}
	data := &email.TicketReplyEmail{
		Cid:   d.Cid,
		Reply: d.Reply,
		Title: d.Title,
	}
	return e.sendEmailTemplate(config.EmailTicketReply, d.TargetEmail, data)
}

func (e *EmailServer) SendWelcome(_ context.Context, d *pb.Welcome) (*pb.SendResponse, error) {
	if !e.extractAndValidateFields(d) {
		return nil, status.Error(codes.InvalidArgument, "missing required argument")
	}
	data := &email.WelcomeEmail{
		Cid: d.Cid,
	}
	return e.sendEmailTemplate(config.EmailWelcome, d.TargetEmail, data)
}

const (
	VerifySuccess int32 = iota
	VerifyExpired
	VerifyInvalid
	VerifyUnknown
)

func (e *EmailServer) VerifyEmailCode(_ context.Context, d *pb.VerifyCode) (*pb.VerifyResponse, error) {
	if !e.extractAndValidateFields(d) {
		return nil, status.Error(codes.InvalidArgument, "missing required argument")
	}
	err := e.manager.VerifyEmailCode(d.Email, d.Code)
	if err == nil {
		return &pb.VerifyResponse{Success: true, Code: VerifySuccess}, nil
	}
	if errors.Is(err, email.ErrEmailCodeExpired) {
		return &pb.VerifyResponse{Success: false, Code: VerifyExpired}, nil
	}
	if errors.Is(err, email.ErrEmailCodeInvalid) {
		return &pb.VerifyResponse{Success: false, Code: VerifyInvalid}, nil
	}
	return &pb.VerifyResponse{Success: false, Code: VerifyUnknown}, status.Error(codes.Internal, "failed to verify email d")
}

func (e *EmailServer) RemoveEmailCode(_ context.Context, d *pb.RemoveVerifyCode) (*pb.RemoveVerifyCodeResponse, error) {
	if !e.extractAndValidateFields(d) {
		return nil, status.Error(codes.InvalidArgument, "missing required argument")
	}
	e.logger.Infof("remove %s email code", d.Email)
	e.manager.RemoveEmailCode(d.Email)
	return &pb.RemoveVerifyCodeResponse{Success: true}, nil
}
