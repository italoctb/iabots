package validators

import (
	"errors"
	"regexp"
)

var hhmmRegex = regexp.MustCompile(`^(2[0-3]|[01][0-9]):([0-5][0-9])$`)

// ValidateHHMM checks if the input string is in "HH:MM" format (24-hour)
func ValidateHHMM(input string) error {
	if input == "" {
		return nil // permite vazio
	}
	if !hhmmRegex.MatchString(input) {
		return errors.New("Horário inválido. Formato esperado: HH:MM (24 horas)")
	}
	return nil
}
