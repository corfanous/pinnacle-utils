package pkg_test

import (
	"testing"

	. "github.com/corfanous/pinnacle-utils/pkg"
)

func TestEmailIsValid(t *testing.T) {
	suits := []struct {
		Name  string
		Email string
		Valid bool
	}{{
		Name:  "Empty email",
		Email: "",
		Valid: false,
	}, {
		Name:  "Name with empty address",
		Email: "Kofi <>",
		Valid: false,
	}, {
		Name:  "Name with address",
		Email: "kofi@gmail.com",
		Valid: true,
	}}
	for _, suit := range suits {
		if EmailIsValid(suit.Email) == suit.Valid {
			t.Logf("%s: %s: %s", suit.Name, suit.Email, "valid")
		} else {
			t.Errorf("%s: %s: %s", suit.Name, suit.Email, "invalid")
		}
	}
}
