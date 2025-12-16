// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

// Package config
package config

import (
	"email-service/src/interfaces/global"
	"errors"
	"fmt"
	"html/template"
	"net/url"
	"path"
	"time"

	"golang.org/x/sync/errgroup"
	"gopkg.in/gomail.v2"
	"half-nothing.cn/service-core/interfaces/config"
	"half-nothing.cn/service-core/utils"
)

type Template struct {
	Enable   bool   `yaml:"enable"`
	FileName string `yaml:"file_name"`
	Subject  string `yaml:"subject"`
	// 内部字段
	LocalPath string `yaml:"-"`
	Type      Email  `yaml:"-"`
}

func (t *Template) Verify() (bool, error) {
	if !t.Enable {
		t.Type.Data.Enable = false
		return true, nil
	}
	remoteFileUrl, err := url.JoinPath(*global.DownloadPrefix, t.Type.Data.RemotePath)
	if err != nil {
		return false, fmt.Errorf("failed to get remote path: %v", err)
	}
	localFilePath := path.Join(t.LocalPath, t.FileName)
	data, err := config.ReadOrDownloadFile(localFilePath, remoteFileUrl)
	if err != nil {
		return false, fmt.Errorf("failed to read or download file: %v", err)
	}
	parsedTemplate, err := template.New(t.Type.Value).Parse(string(data))
	if err != nil {
		return false, fmt.Errorf("failed to parse template: %v", err)
	}
	t.Type.Data.Template = parsedTemplate
	t.Type.Data.Subject = t.Subject
	t.Type.Data.Enable = true
	return true, nil
}

type TemplateConfig struct {
	VerifyCodeEmail            *Template `yaml:"verify_code_email"`
	WelcomeEmail               *Template `yaml:"welcome_email"`
	RatingChangeEmail          *Template `yaml:"rating_change_email"`
	KickedFromServerEmail      *Template `yaml:"kicked_from_server_email"`
	PasswordChangeEmail        *Template `yaml:"password_change_email"`
	PasswordResetEmail         *Template `yaml:"password_reset_email"`
	ApplicationPassedEmail     *Template `yaml:"application_passed_email"`
	ApplicationRejectedEmail   *Template `yaml:"application_rejected_email"`
	ApplicationProcessingEmail *Template `yaml:"application_processing_email"`
	TicketReplyEmail           *Template `yaml:"ticket_reply_email"`
	ActivityPilotJoinEmail     *Template `yaml:"activity_pilot_join_email"`
	ActivityPilotLeaveEmail    *Template `yaml:"activity_pilot_leave_email"`
	ActivityAtcJoinEmail       *Template `yaml:"activity_atc_join_email"`
	ActivityAtcLeaveEmail      *Template `yaml:"activity_atc_leave_email"`
	InstructorChangeEmail      *Template `yaml:"instructor_change_email"`
	BannedEmail                *Template `yaml:"banned_email"`
	RoleChangeEmail            *Template `yaml:"role_change_email"`
	PermissionChangeEmail      *Template `yaml:"permission_change_email"`
	// 内部字段
	LocalPath string `yaml:"-"`
}

func (t *TemplateConfig) InitDefaults() {
	t.VerifyCodeEmail = &Template{Enable: true, FileName: "verify_code.template", Subject: "邮箱验证码", Type: EmailVerifyCode}
	t.WelcomeEmail = &Template{Enable: true, FileName: "welcome.template", Subject: "欢迎注册", Type: EmailWelcome}
	t.RatingChangeEmail = &Template{Enable: true, FileName: "atc_rating_change.template", Subject: "管制权限变更通知", Type: EmailRatingChange}
	t.KickedFromServerEmail = &Template{Enable: true, FileName: "kicked_from_server.template", Subject: "您已被踢出服务器", Type: EmailKickedFromServer}
	t.PasswordChangeEmail = &Template{Enable: true, FileName: "password_change.template", Subject: "飞控密码更改通知", Type: EmailPasswordChange}
	t.PasswordResetEmail = &Template{Enable: true, FileName: "password_reset.template", Subject: "飞控密码重置通知", Type: EmailPasswordReset}
	t.ApplicationPassedEmail = &Template{Enable: true, FileName: "application_passed.template", Subject: "管制员申请通过", Type: EmailApplicationPassed}
	t.ApplicationRejectedEmail = &Template{Enable: true, FileName: "application_rejected.template", Subject: "管制员申请被拒", Type: EmailApplicationRejected}
	t.ApplicationProcessingEmail = &Template{Enable: true, FileName: "application_processing.template", Subject: "管制面试通知", Type: EmailApplicationProcessing}
	t.ApplicationRejectedEmail = &Template{Enable: true, FileName: "application_rejected.template", Subject: "管制员申请被拒", Type: EmailApplicationRejected}
	t.ApplicationProcessingEmail = &Template{Enable: true, FileName: "application_processing.template", Subject: "管制面试通知", Type: EmailApplicationProcessing}
	t.TicketReplyEmail = &Template{Enable: true, FileName: "ticket_reply.template", Subject: "工单回复通知", Type: EmailTicketReply}
	t.TicketReplyEmail = &Template{Enable: true, FileName: "ticket_reply.template", Subject: "工单回复通知", Type: EmailTicketReply}
	t.ActivityPilotJoinEmail = &Template{Enable: true, FileName: "activity_pilot_join.template", Subject: "活动报名成功", Type: EmailActivityPilotJoin}
	t.ActivityPilotLeaveEmail = &Template{Enable: true, FileName: "activity_pilot_leave.template", Subject: "退出活动成功", Type: EmailActivityPilotLeave}
	t.ActivityAtcJoinEmail = &Template{Enable: true, FileName: "activity_atc_join.template", Subject: "活动报名成功", Type: EmailActivityAtcJoin}
	t.ActivityAtcLeaveEmail = &Template{Enable: true, FileName: "activity_atc_leave.template", Subject: "退出活动成功", Type: EmailActivityAtcLeave}
	t.InstructorChangeEmail = &Template{Enable: true, FileName: "instructor_change.template", Subject: "教员变更通知", Type: EmailInstructorChange}
	t.BannedEmail = &Template{Enable: true, FileName: "banned.template", Subject: "您已被封禁", Type: EmailBanned}
	t.RoleChangeEmail = &Template{Enable: true, FileName: "role_change.template", Subject: "飞控角色变更通知", Type: EmailRoleChange}
	t.PermissionChangeEmail = &Template{Enable: true, FileName: "permission_change.template", Subject: "飞控权限变更通知", Type: EmailPermissionChange}
}

func (t *TemplateConfig) Verify() (bool, error) {
	fields := []*Template{t.VerifyCodeEmail, t.WelcomeEmail, t.RatingChangeEmail, t.KickedFromServerEmail,
		t.PasswordChangeEmail, t.PasswordResetEmail, t.ApplicationPassedEmail, t.ApplicationRejectedEmail,
		t.ApplicationProcessingEmail, t.TicketReplyEmail, t.ActivityPilotJoinEmail, t.ActivityPilotLeaveEmail,
		t.ActivityAtcJoinEmail, t.ActivityAtcLeaveEmail, t.InstructorChangeEmail, t.BannedEmail, t.RoleChangeEmail,
		t.PermissionChangeEmail}
	utils.ForEach(fields, func(_ int, field *Template) {
		field.LocalPath = t.LocalPath
	})
	eg := errgroup.Group{}
	for _, field := range fields {
		eg.Go(func() error {
			ok, err := field.Verify()
			if ok {
				return nil
			}
			return err
		})
	}
	if err := eg.Wait(); err != nil {
		return false, err
	}
	return true, nil
}

type TemplatesConfig struct {
	LocalPath string          `yaml:"local_path"`
	Templates *TemplateConfig `yaml:"templates"`
}

func (t *TemplatesConfig) InitDefaults() {
	t.LocalPath = "data/templates"
	t.Templates = &TemplateConfig{}
	t.Templates.InitDefaults()
}

func (t *TemplatesConfig) Verify() (bool, error) {
	t.Templates.LocalPath = t.LocalPath
	return t.Templates.Verify()
}

type EmailConfig struct {
	Host           string           `yaml:"host"`
	Port           int              `yaml:"port"`
	Username       string           `yaml:"username"`
	Password       string           `yaml:"password"`
	VerifyExpire   string           `yaml:"verify_expire"`
	VerifyInterval string           `yaml:"verify_interval"`
	Template       *TemplatesConfig `yaml:"template"`
	// 内部字段
	VerifyExpireDuration   time.Duration     `yaml:"-"`
	VerifyIntervalDuration time.Duration     `yaml:"-"`
	Server                 *gomail.Dialer    `json:"-"`
	Closer                 gomail.SendCloser `json:"-"`
}

func (e *EmailConfig) InitDefaults() {
	e.Host = "smtp.example.com"
	e.Port = 587
	e.Username = ""
	e.Password = ""
	e.VerifyExpire = "5m"
	e.VerifyInterval = "1m"
	e.Template = &TemplatesConfig{}
	e.Template.InitDefaults()
}

func (e *EmailConfig) Verify() (bool, error) {
	if e.Host == "" {
		return false, errors.New("smtp server host cannot be empty")
	}
	if e.Port <= 0 {
		return false, errors.New("smtp server port cannot be less than or equal to 0")
	}
	if e.Port >= 65535 {
		return false, errors.New("smtp server port cannot be greater than 65535")
	}
	if e.Username == "" {
		return false, errors.New("smtp server username cannot be empty")
	}
	if e.Password == "" {
		return false, errors.New("smtp server password cannot be empty")
	}

	e.Server = gomail.NewDialer(e.Host, e.Port, e.Username, e.Password)
	dial, err := e.Server.Dial()
	if err != nil {
		return false, fmt.Errorf("connect to smtp server fail, %v", err)
	}
	e.Closer = dial

	if e.VerifyExpire == "" {
		return false, errors.New("verify expire cannot be empty")
	}
	if duration, err := time.ParseDuration(e.VerifyExpire); err != nil {
		return false, err
	} else {
		e.VerifyExpireDuration = duration
	}
	if e.VerifyInterval == "" {
		return false, errors.New("verify interval cannot be empty")
	}
	if duration, err := time.ParseDuration(e.VerifyInterval); err != nil {
		return false, err
	} else {
		e.VerifyIntervalDuration = duration
	}
	return e.Template.Verify()
}

type EmailData struct {
	Enable     bool
	Template   *template.Template
	RemotePath string
	Subject    string
}

type Email *utils.Enum[string, *EmailData]

var (
	EmailVerifyCode            = utils.NewEnum("verify_code", &EmailData{Enable: true, RemotePath: "/docker/data/template/verify_code.template"})
	EmailWelcome               = utils.NewEnum("welcome", &EmailData{Enable: true, RemotePath: "/docker/data/template/welcome.template"})
	EmailRatingChange          = utils.NewEnum("rating_change", &EmailData{Enable: true, RemotePath: "/docker/data/template/atc_rating_change.template"})
	EmailKickedFromServer      = utils.NewEnum("kicked_from_server", &EmailData{Enable: true, RemotePath: "/docker/data/template/kicked_from_server.template"})
	EmailPasswordChange        = utils.NewEnum("password_change", &EmailData{Enable: true, RemotePath: "/docker/data/template/password_change.template"})
	EmailPasswordReset         = utils.NewEnum("password_reset", &EmailData{Enable: true, RemotePath: "/docker/data/template/password_reset.template"})
	EmailApplicationPassed     = utils.NewEnum("application_passed", &EmailData{Enable: true, RemotePath: "/docker/data/template/application_passed.template"})
	EmailApplicationRejected   = utils.NewEnum("application_rejected", &EmailData{Enable: true, RemotePath: "/docker/data/template/application_rejected.template"})
	EmailApplicationProcessing = utils.NewEnum("application_processing", &EmailData{Enable: true, RemotePath: "/docker/data/template/application_processing.template"})
	EmailTicketReply           = utils.NewEnum("ticket_reply", &EmailData{Enable: true, RemotePath: "/docker/data/template/ticket_reply.template"})
	EmailActivityPilotJoin     = utils.NewEnum("activity_pilot_join", &EmailData{Enable: true, RemotePath: "/docker/data/template/activity_pilot_join.template"})
	EmailActivityPilotLeave    = utils.NewEnum("activity_pilot_leave", &EmailData{Enable: true, RemotePath: "/docker/data/template/activity_pilot_leave.template"})
	EmailActivityAtcJoin       = utils.NewEnum("activity_atc_join", &EmailData{Enable: true, RemotePath: "/docker/data/template/activity_atc_join.template"})
	EmailActivityAtcLeave      = utils.NewEnum("activity_atc_leave", &EmailData{Enable: true, RemotePath: "/docker/data/template/activity_atc_leave.template"})
	EmailInstructorChange      = utils.NewEnum("instructor_change", &EmailData{Enable: true, RemotePath: "/docker/data/template/instructor_change.template"})
	EmailBanned                = utils.NewEnum("banned", &EmailData{Enable: true, RemotePath: "/docker/data/template/banned.template"})
	EmailRoleChange            = utils.NewEnum("role_change", &EmailData{Enable: true, RemotePath: "/docker/data/template/role_change.template"})
	EmailPermissionChange      = utils.NewEnum("permission_change", &EmailData{Enable: true, RemotePath: "/docker/data/template/permission_change.template"})
)

var Emails = utils.NewEnums(
	EmailVerifyCode,
	EmailWelcome,
	EmailRatingChange,
	EmailKickedFromServer,
	EmailPasswordChange,
	EmailPasswordReset,
	EmailApplicationPassed,
	EmailApplicationRejected,
	EmailApplicationProcessing,
	EmailTicketReply,
	EmailActivityPilotJoin,
	EmailActivityPilotLeave,
	EmailActivityAtcJoin,
	EmailActivityAtcLeave,
	EmailInstructorChange,
	EmailBanned,
	EmailRoleChange,
	EmailPermissionChange,
)
