package utils

import "regexp"

func ValidatePhone(phone string) bool {
	reg := regexp.MustCompile(`^\(?\d{2}\)?\s?\d{4,5}-?\d{4}$`)

	return reg.MatchString(phone)
}
