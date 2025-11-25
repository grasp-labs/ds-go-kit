package fakes

import (
	"fmt"

	"github.com/google/uuid"
)

func Caller(id uuid.UUID) error {
	return SubCall(id)
}

func SubCall(id uuid.UUID) error {
	return CauseError(id)
}

func CauseError(id uuid.UUID) error {
	return fmt.Errorf("New err with id: %s", id.String())
}
