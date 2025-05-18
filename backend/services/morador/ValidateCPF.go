package moradorService

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// https://www.macoratti.net/alg_cpf.htm
func ValidateCPF(cpf string) bool {
	cpf = strings.ReplaceAll(cpf, ".", "")
	cpf = strings.ReplaceAll(cpf, "-", "")

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
