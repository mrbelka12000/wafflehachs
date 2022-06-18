package validate

import (
	"errors"
	"regexp"
)

func ValidatingEmail(email string) error {
	Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !Re.MatchString(email) {
		return errors.New("Invaild email address")
	}
	return nil
}

// ValidatingPassword ..
func ValidatingPassword(password string) error {
	var checked bool
	for _, l := range password {
		if l < 32 || l > 126 {
			checked = true
		}
	}
	if checked {
		return errors.New("Пожалуйста, используйте только латинский алфавит")
	}
	if len(password) < 8 {
		return errors.New("Пароль должен содержать более 8 символов")
	}
	var numbers, bigletters, smallletters, specialsymbols bool
	for _, v := range password {
		if v >= 'a' && v <= 'z' {
			smallletters = true
		} else if v >= 'A' && v <= 'Z' {
			bigletters = true
		} else if v >= '0' && v <= '9' {
			numbers = true
		} else {
			specialsymbols = true
		}
	}
	if !numbers {
		return errors.New("Пароль должен содержать цифры")
	}
	if !bigletters {
		return errors.New("Пароль должен содержать заглавные буквы")
	}
	if !smallletters {
		return errors.New("Пароль должен содержать строчные буквы")
	}
	if !specialsymbols {
		return errors.New("Пароль должен содержать символы")
	}
	return nil
}
