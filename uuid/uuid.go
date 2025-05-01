package uuid

import (
	"github.com/google/uuid"
)

// New creates a new uuid string
func New() string {
	return uuid.New().String()
}

// UUID validates a uuid string
func UUID(value string) bool {
	_, err := uuid.Parse(value)
	return err == nil
}
