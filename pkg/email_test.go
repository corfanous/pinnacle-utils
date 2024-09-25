package pkg_test

import (
	"strings"
	"testing"

	. "github.com/corfanous/pinnacle-utils/pkg"
)

func TestEmailIsValid(t *testing.T) {
	suits := []struct {
		Name    string
		Email   string
		Expects bool
	}{{
		Name:    "Empty email",
		Email:   "",
		Expects: false,
	}, {
		Name:    "Name with empty address",
		Email:   "Kofi <>",
		Expects: false,
	}, {
		Name:    "Name with address",
		Email:   "kofi@gmail.com",
		Expects: true,
	}, {
		Name:    "List of valid addresses",
		Email:   "g@gmal.com,x@yahoo.com,lorna@stx.com",
		Expects: true,
	}}
	for _, suit := range suits {
		if EmailIsValid(suit.Email) == suit.Expects {
			t.Logf("%s: %s: %s", suit.Name, suit.Email, "valid")
		} else {
			t.Errorf("%s: %s: %s", suit.Name, suit.Email, "invalid")
		}
	}
}

func TestHeader(t *testing.T) {
	suits := []struct {
		Name    string
		Header  Header
		Expects bool
	}{
		{
			Name: "FROM",
			Header: Header{
				"from": []string{"hello@hello.com"},
			},
			Expects: true,
		},
		{
			Name: "TO",
			Header: Header{
				"to": []string{"a@test.com", "b@test.com"},
			},
			Expects: true,
		},
	}
	for _, suit := range suits {
		if vals, ok := suit.Header[strings.ToLower(suit.Name)]; ok == suit.Expects {
			t.Logf("%s:%s\n", suit.Name, strings.Join(vals, ","))
		} else {
			t.Errorf("%s:%s\n", suit.Name, strings.Join(vals, ","))
		}
	}
}
func TestSend(t *testing.T) {
	// auth := pkg.NewPlainAuth("", "test@test.com", "test")
	client := NewMailClient("", 2525, nil)
	msg := &MailMessage{}
	msg.SetHeader("From", "sender1@test.com")
	msg.SetHeader("To", "reciever1@test.com")
	msg.SetHeader("Subject", "testing")
	msg.SetStringBody("test2 message!")
	message, _ := msg.Build()

	err := client.Send(msg.GetHeader("From"), msg.GetHeaderValues("To"), msg.GetHeader("Subject"), message)
	if err != nil {
		t.Error(err)
	} else {
		t.Log("Mail successfully sent")
	}
}
