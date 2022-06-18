package request

import (
	"errors"
	"strings"
	"wafflehacks/entities/busymode"
	"wafflehacks/models"
)

type PsychoSignUpRequest struct {
	CV string `json:"cv"`
	ClientSignUpRequest
}

func (sur *PsychoSignUpRequest) Build() *models.Psychologist {
	user := models.User{
		Firstname: sur.FirstName,
		Lastname:  sur.LastName,
		Username:  sur.UserName,
		Email:     sur.Email,
		Password:  &sur.Password,
		Age:       sur.Age,
	}

	return &models.Psychologist{
		User:     user,
		BusyMode: busymode.ActiveMode,
	}
}

func (sur *PsychoSignUpRequest) Validate() error {
	text := ""
	text = strings.Replace(sur.FirstName, "\r\n", " ", -1)
	text = strings.Trim(text, " ")
	if text == "" {
		return errors.New("пустое имя")
	}

	text = strings.Replace(sur.LastName, "\r\n", " ", -1)
	text = strings.Trim(text, " ")
	if text == "" {
		return errors.New("пустая фамилия")
	}

	text = strings.Replace(sur.UserName, "\r\n", " ", -1)
	text = strings.Trim(text, " ")
	if text == "" {
		return errors.New("пустой никнейм")
	}

	text = strings.Replace(sur.Email, "\r\n", " ", -1)
	text = strings.Trim(text, " ")
	if text == "" {
		return errors.New("пустой email")
	}

	text = strings.Replace(sur.Password, "\r\n", " ", -1)
	text = strings.Trim(text, " ")
	if text == "" {
		return errors.New("пустой пароль")
	}

	return nil
}
