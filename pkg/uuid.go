package pkg

import (
	"github.com/google/uuid"
)

// Generate UUID v4 string
func NewUUID() string {
	return uuid.NewString()
}

// Test a given string for uuid
func UUIDString(s string) bool {
	_, err := uuid.Parse(s)
	return err == nil
}
