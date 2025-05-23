package cpf

import (
	"errors"
	"strings"
	"unicode"
	"unicode/utf8"
)

type CPF struct {
	Value string
}

func New(s string) (CPF, error) {
	clean := cleanCPF(s)

	if !isValidCpf(clean) {
		return CPF{}, ErrInvalidCPF
	}

	return CPF{Value: clean}, nil
}

var ErrInvalidCPF = errors.New("Invalid cpf")

func cleanCPF(s string) string {
	// TODO: Use regex instead of ReplaceAll
	s = strings.ReplaceAll(s, ".", "")
	s = strings.ReplaceAll(s, "-", "")

	return s
}

// https://www.macoratti.net/alg_cpf.htm
func isValidCpf(cpf string) bool {
	if len(cpf) != 11 {
		return false
	}

	if areAllDigitsEqual(cpf) {
		return false
	}

	var digits [11]int
	for i, digit := range cpf {
		if !unicode.IsDigit(digit) {
			return false
		}

		digits[i] = int(digit - '0')
	}

	firstVerificationDigit := int(cpf[9] - '0')
	secondVerificationDigit := int(cpf[10] - '0')

	if firstVerificationDigit != getFirstVerificationDigit(digits) {
		return false
	}

	if secondVerificationDigit != getSecondVerificationDigit(digits) {
		return false
	}

	return true
}

func areAllDigitsEqual(val string) bool {
	for index, digit := range val {
		if index == 0 {
			continue
		}

		prevDigit, _ := utf8.DecodeRune([]byte{val[index-1]})
		if prevDigit != digit {
			return false
		}
	}

	return true
}

func getFirstVerificationDigit(cpf [11]int) int {
	var multiplyResult int

	for index, curr := range cpf {
		if index > 8 {
			break
		}

		multiplyResult += curr * (10 - index)
	}

	verificationDigit := (multiplyResult * 10) % 11

	if verificationDigit == 10 {
		return 0
	}

	return verificationDigit
}

func getSecondVerificationDigit(cpf [11]int) int {
	var multiplyResult int

	for index, curr := range cpf {
		if index > 9 {
			break
		}

		multiplyResult += curr * (11 - index)
	}

	verificationDigit := (multiplyResult * 10) % 11

	if verificationDigit == 10 {
		return 0
	}

	return verificationDigit
}
