package request

import (
	"errors"
	"strings"
	"wafflehacks/models"
)

type ClientSignUpRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	UserName  string `json:"userName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Age       int    `json:"age"`
}

func (sur *ClientSignUpRequest) Build() *models.Client {
	user := models.User{
		Firstname: sur.FirstName,
		Lastname:  sur.LastName,
		Username:  sur.UserName,
		Email:     sur.Email,
		Password:  &sur.Password,
		Age:       sur.Age,
	}
	return &models.Client{User: user}
}

func (sur *ClientSignUpRequest) Validate() error {
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
