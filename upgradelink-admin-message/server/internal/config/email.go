package config

import (
	"crypto/tls"
	"fmt"
	"net/smtp"

	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type EmailConf struct {
	AuthType  string
	EmailAddr string
	Password  string
	HostName  string
	Identify  string
	Secret    string
	Port      int
	TLS       bool
}

// NewAuth creates the auth from config
func (e EmailConf) NewAuth() *smtp.Auth {
	var auth smtp.Auth
	switch e.AuthType {
	case "plain":
		auth = smtp.PlainAuth(e.Identify, e.EmailAddr, e.Password, e.HostName)
	case "CRAMMD5":
		auth = smtp.CRAMMD5Auth(e.EmailAddr, e.Secret)
	}
	return &auth
}

func (e EmailConf) NewClient() (*smtp.Client, error) {
	hostAddress := fmt.Sprintf("%s:%d", e.HostName, e.Port)
	if e.TLS == true {
		tlsconfig := &tls.Config{
			InsecureSkipVerify: true,
			ServerName:         hostAddress,
		}

		// initialize connection
		conn, err := tls.Dial("tcp", hostAddress, tlsconfig)
		if err != nil {
			logx.Error("failed to dial connection to email server, please check the host and tls config", logx.Field("config", e), logx.Field("detail", err))
			return nil, errorx.NewInvalidArgumentError("failed to connect email server, please check the host and tls config")
		}

		// get client
		c, err := smtp.NewClient(conn, e.HostName)
		if err != nil {
			logx.Error("failed to create smtp client, please check the host and tls config", logx.Field("config", e), logx.Field("detail", err))
			return nil, errorx.NewInvalidArgumentError("failed to connect email server, please check the host and tls config")
		}

		err = c.Auth(*e.NewAuth())
		if err != nil {
			logx.Error("failed to get the auth of smtp server, please check the identify and secret", logx.Field("config", e), logx.Field("detail", err))
			return nil, errorx.NewInvalidArgumentError("failed to authorize the server, please check the identify and secret")
		}

		return c, nil
	} else {
		// initialize connection
		c, err := smtp.Dial(hostAddress)
		if err != nil {
			logx.Error("failed to dial connection to email server, please check the host and tls config", logx.Field("config", e), logx.Field("detail", err))
			return nil, errorx.NewInvalidArgumentError("failed to connect email server, please check the host and tls config")
		}

		err = c.Auth(*e.NewAuth())
		if err != nil {
			logx.Error("failed to get the auth of smtp server, please check the identify and secret", logx.Field("config", e), logx.Field("detail", err))
			return nil, errorx.NewInvalidArgumentError("failed to authorize the server, please check the identify and secret")
		}

		return c, nil
	}
}
