package request

import "wafflehacks/models"

type ClientSignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (cli *ClientSignInRequest) Build() *models.User {
	return &models.User{
		Email:    cli.Email,
		Password: &cli.Password,
	}
}

// func (cli *ClientSignInRequest) Validate() error {

// }
