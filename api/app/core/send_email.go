package core

import (
	"context"
	"crypto/tls"
	"errors"
	"github.com/Ovsienko023/reporter/app/domain"
	"gopkg.in/gomail.v2"
)

// SendEmail возвращает следующие ошибки:
// ErrUnauthorized
// ErrInternal
func (c *Core) SendEmail(ctx context.Context, msg *domain.SendEmailRequest) error {
	_, err := c.authorize(msg.Token)
	if err != nil {
		return err
	}

	err = Send(
		&Config{
			Host:     "mail.trueconf.com",
			Port:     465,
			Username: msg.Email,
			Password: msg.Password,
		},
		&Letter{
			From:    msg.Email,
			To:      msg.Recipients,
			Subject: msg.Subject,
			Body:    msg.Body,
		},
	)
	if err != nil {
		return ErrInternal
	}

	return nil
}

type Config struct {
	Host     string `json:"host,omitempty"`
	Port     int    `json:"port,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type Letter struct {
	From    string   `json:"from,omitempty"`
	To      []string `json:"to,omitempty"`
	Subject string   `json:"subject,omitempty"`
	Body    string   `json:"body,omitempty"`
}

func Send(cfg *Config, msg *Letter) error {
	d := gomail.NewDialer(cfg.Host, cfg.Port, cfg.Username, cfg.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send emails using d.
	m := gomail.NewMessage()
	m.SetHeader("From", msg.From)

	if len(msg.To) == 0 {
		return errors.New("not to") // todo
	}
	m.SetHeader("To", msg.To...)

	m.SetHeader("Subject", msg.Subject)
	m.SetBody("text/plain", msg.Body)

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		return ErrInternal // todo
	}

	return nil
}
