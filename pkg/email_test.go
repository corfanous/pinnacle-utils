package pkg_test

import (
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
