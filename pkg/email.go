package pkg

import (
	"bytes"
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

type Header map[string][]string

// Retrieves the first element of the Header Map
func (h Header) Get(key string) string {
	v := h[key]
	if len(v) == 0 {
		return ""
	}
	return v[0]
}

// Retrieves the list of values associated with a Key
func (h Header) Values(key string) []string {
	if h == nil {
		return nil
	}
	return h[key]
}

// Sets the key value per of a header
func (h Header) Set(key, value string) {
	h[key] = append(h[key], value)
}

type MailMessage struct {
	headers Header
	// content string
	body []byte
}

// Retrieves a MailMessage header value for a given key
func (mm *MailMessage) GetHeader(key string) string {
	return mm.headers.Get(key)
}
func (mm *MailMessage) GetHeaderValues(key string) []string {
	return mm.headers.Values(key)
}

// Add a header
func (mm *MailMessage) SetHeader(key string, value string) {
	if mm.headers == nil {
		mm.headers = make(Header)
	}
	mm.headers.Set(key, value)
}

// Set the content value
func (mm *MailMessage) SetStringBody(msg string) {
	if mm.body == nil {
		mm.body = []byte(msg)
	} else {
		mm.body = append(mm.body, []byte(msg)...)
	}
}
func (mm *MailMessage) SetByteBody(msg []byte) {
	if mm.body == nil {
		mm.body = msg
	} else {
		mm.body = append(mm.body, msg...)
	}
}

// Converts a mail to a corresponding
func (mm *MailMessage) Build() ([]byte, error) {
	var buf bytes.Buffer
	for key, vals := range mm.headers {
		_, err := buf.WriteString(fmt.Sprintf("%s: %s\n", key, strings.Join(vals, ",")))
		if err != nil {
			return nil, err
		}
	}
	//write the body
	//start
	_, err := buf.WriteString("\n")
	if err != nil {
		return nil, err
	}
	_, err = buf.Write(mm.body)
	if err != nil {
		return nil, err
	}
	//end
	buf.WriteString("\n")
	return buf.Bytes(), nil
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
