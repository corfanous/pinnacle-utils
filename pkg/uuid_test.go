package pkg_test

import (
	"testing"

	. "github.com/corfanous/pinnacle-utils/pkg"
)

func notString(cond bool) string {
	if cond == false {
		return "not a"
	} /*else {
		return "is a"
	}*/
	return "a"
}
func TestUUIDString(t *testing.T) {
	suits := []struct {
		Name    string
		UUID    string
		Expects bool
	}{
		{
			Name:    "Number",
			UUID:    "20000",
			Expects: false,
		},
		{
			Name:    "String",
			UUID:    "HELLO.WORLD.GH",
			Expects: false,
		},
		{
			Name:    "UUID v4",
			UUID:    NewUUID(),
			Expects: true,
		},
	}
	for _, suit := range suits {
		if UUIDString(suit.UUID) == suit.Expects {
			t.Logf("%s: is %s uuid string", suit.UUID, notString(suit.Expects))
		} else {
			t.Errorf("Invalid test result: %v", suit)
		}
	}
}
