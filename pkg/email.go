package pkg

import (
	"errors"
	"fmt"
	"net/mail"
	"net/smtp"
	"strings"
)

const DEFAULT_SMPT_PORT int = 587

func EmailIsValid(email string) bool {
	_, err := mail.ParseAddressList(email)
	return err == nil
}

// Creates a new SMPT Plan Auth object
func NewPlainAuth(hostname string, sender, password string) smtp.Auth {
	return smtp.PlainAuth("", sender, password, hostname)
}

type MailMessage struct {
	msg []byte
}

func (mm *MailMessage) SetHeader(key, value string) {
	header := fmt.Sprintf("%s: %s\n", key, value)
	if mm.msg == nil {
		mm.msg = []byte(header)
	} else {
		mm.msg = append(mm.msg, []byte(header)...)
	}
}
func (mm *MailMessage) SetContent(msg string) {
	msg = fmt.Sprintf("\n%s\n", msg)
	if mm.msg == nil {
		mm.msg = []byte(msg)
	} else {
		mm.msg = append(mm.msg, []byte(msg)...)
	}
}
func (mm *MailMessage) Build() ([]byte, error) {
	if mm.msg == nil {
		return nil, errors.New("empty message")
	}
	return mm.msg, nil
}

type MailClient struct {
	port     int
	hostName string
	auth     smtp.Auth
}

// Creates a new MailClient
func NewMailClient(hostname string, port int, auth smtp.Auth) *MailClient {
	return &MailClient{hostName: hostname, port: port, auth: auth}
}

// Sends an email to server using smpt
func (e *MailClient) Send(from string, to []string, subject string, msg []byte) error {
	if strings.TrimSpace(from) == "" {
		return errors.New("sender email address cannot be empty")
	}
	if len(to) == 0 {
		return errors.New("reciever email addresses cannot be empty")
	}
	if len(msg) == 0 {
		return errors.New("message cannot be empty")
	}
	addr := fmt.Sprintf("%s:%d", e.hostName, e.port)
	return smtp.SendMail(addr, e.auth, from, to, msg)
}

func NewMailMessage(from string, to []string, subject string, body []byte) ([]byte, error) {
	if !EmailIsValid(from) {
		return nil, fmt.Errorf("invalid sender email address: %s", from)
	}
	from = fmt.Sprintf("From: %s\n", from)
	recievers := strings.Join(to, ",")
	if !EmailIsValid(recievers) {
		return nil, fmt.Errorf("invalid reciever email address: %s", recievers)
	}
	recievers = fmt.Sprintf("To: %s\n", recievers)
	subject = fmt.Sprintf("Subject: %s\n", subject)
	header := fmt.Sprintf("%s%s%s\n\n", from, recievers, subject)

	header_byte := []byte(header)
	message := append(header_byte, body...)

	return message, nil
}
