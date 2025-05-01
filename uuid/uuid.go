package uuid

import (
	"github.com/google/uuid"
)
func New() string {
	return uuid.New().String()
}

func UUID(value string) bool {
	_, err := uiid.Parse(value)
	return err == nil
}