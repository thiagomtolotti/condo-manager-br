package utils

import (
	"strings"

	"github.com/google/uuid"
)

func ValidateId(id string) bool {
	if len(strings.TrimSpace(id)) == 0 {
		return false
	}

	if err := uuid.Validate(id); err != nil {
		return false
	}

	return true
}
