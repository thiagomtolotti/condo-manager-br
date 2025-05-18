package moradorService

import (
	"strings"
	"unicode"
)

// https://www.macoratti.net/alg_cpf.htm
func ValidateCPF(cpf string) bool {
	// TODO: Sanitize cpf
	cpf = strings.ReplaceAll(cpf, ".", "")
	cpf = strings.ReplaceAll(cpf, "-", "")

	if len(cpf) != 11 {
		return false
	}

	var multiplyResult int
	for index, character := range cpf {
		if index > 8 {
			break
		}
		if !unicode.IsDigit(character) {
			return false
		}

		number := int(character - '0')

		multiplyResult += number * (10 - index)
	}
	var verificationDigits [2]int
	verificationDigits[0] = ((multiplyResult * 10) % 11) % 10

	if verificationDigits[0] != int(cpf[9]-'0') {
		return false
	}

	var newMultiplyResult int
	for index, character := range cpf {
		if index > 9 {
			break
		}

		number := int(character - '0')
		newMultiplyResult += number * (11 - index)
	}

	verificationDigits[1] = ((newMultiplyResult * 10) % 11) % 10

	if verificationDigits[1] != int(cpf[10]-'0') {
		return false
	}

	return true
}

func isDigit(r rune) bool {
	return unicode.IsDigit(r)
}
