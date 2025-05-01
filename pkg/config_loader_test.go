package pkg_test

import (
	"testing"

	. "github.com/corfanous/pinnacle-utils/pkg"
)

func TestLoadConfig(t *testing.T) {
	DSN := "localhost"
	cfg := &struct {
		Dsn string `mapstructure:"DSN"`
	}{}
	err := LoadConfig("./data", cfg)
	if err == nil && cfg.Dsn == DSN {
		t.Logf("LoadConfig successfully loads config: %v", cfg)
	} else {
		t.Errorf("LoadConfig failure: %v\n", err)
	}
}
