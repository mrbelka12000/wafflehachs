package request

import (
	"errors"
	"net/http"
	"strconv"
	"wafflehacks/models"
)

type UpdateUserRequest struct {
	ID          int
	Firstname   string
	Lastname    string
	Username    string
	Description string
	Age         int
	AvatarUrl   string
	HaveAvater  bool
}

func (uur *UpdateUserRequest) Build() *models.User {
	return &models.User{
		ID:          uur.ID,
		Firstname:   uur.Firstname,
		Lastname:    uur.Lastname,
		Username:    uur.Username,
		Description: uur.Description,
		Age:         uur.Age,
		AvatarUrl:   uur.AvatarUrl,
	}
}

func (uur *UpdateUserRequest) BuildRequest(id int, r *http.Request) (*UpdateUserRequest, error) {
	firstname := r.FormValue("firstname")
	lastname := r.FormValue("lastname")
	username := r.FormValue("username")
	description := r.FormValue("description")
	age, err := strconv.Atoi(r.FormValue("age"))
	if err != nil {
		return nil, errors.New("укажите числовой возраст")
	}
	avatarurl := r.FormValue("avatarUrl")

	if firstname == "" {
		return nil, errors.New("пустое имя")
	}
	if lastname == "" {
		return nil, errors.New("пустая фамилия")
	}
	if username == "" {
		return nil, errors.New("пустой никнейм")
	}
	if age < 16 {
		return nil, errors.New("слишком молодой")
	}
	if avatarurl != "" {
		uur.HaveAvater = true
	}
	return &UpdateUserRequest{
		ID:          id,
		Firstname:   firstname,
		Lastname:    lastname,
		Username:    username,
		Description: description,
		Age:         age,
		AvatarUrl:   avatarurl,
	}, nil
}
