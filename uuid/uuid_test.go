package uuid_test

import (
	. "github.com/corfanous/pinnacle-utils/uuid"
	"testing"
)

func notString(cond bool) string {
	if cond == false {
		return "not a"
	}
	return "a"
}
func TestUUID(t *testing.T) {
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
			UUID:    New(),
			Expects: true,
		},
	}
	for _, suit := range suits {
		if UUID(suit.UUID) == suit.Expects {
			t.Logf("%s: is %s uuid string", suit.UUID, notString(suit.Expects))
		} else {
			t.Errorf("Invalid test result: %v", suit)
		}
	}
}
