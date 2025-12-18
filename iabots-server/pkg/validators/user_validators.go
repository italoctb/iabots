package validators

import (
	"errors"
	"regexp"
	"unicode"
)

// emailRegex valida emails simples (não suporta todos os edge cases do RFC, mas cobre 99% dos casos reais)
var emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)

// ValidateEmail retorna true se o email for considerado válido.
func ValidateEmail(email string) error {
	if !emailRegex.MatchString(email) {
		return errors.New("invalid email format")
	}
	return nil
}

// ValidatePassword verifica se a senha tem pelo menos:
// - 6 caracteres
// - uma letra maiúscula
// - uma letra minúscula
// - um número
func ValidatePassword(password string) error {
	if len(password) < 6 {
		return errors.New("a senha deve ter pelo menos 6 caracteres")
	}

	var hasUpper, hasLower, hasDigit bool
	for _, c := range password {
		switch {
		case unicode.IsUpper(c):
			hasUpper = true
		case unicode.IsLower(c):
			hasLower = true
		case unicode.IsDigit(c):
			hasDigit = true
		}
	}

	if !hasUpper {
		return errors.New("a senha deve conter ao menos uma letra maiúscula")
	}
	if !hasLower {
		return errors.New("a senha deve conter ao menos uma letra minúscula")
	}
	if !hasDigit {
		return errors.New("a senha deve conter ao menos um número")
	}

	return nil
}
