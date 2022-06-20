package request

import "wafflehacks/models"

type UserSignUpContinueRequest struct {
	Description string
	Avatar      string
}

func (uc *UserSignUpContinueRequest) Build(userid int) *models.ContinueSignUp {
	return &models.ContinueSignUp{
		UserID:      userid,
		Description: uc.Description,
		Avatar:      uc.Avatar,
	}
}

//Handle стоит ли идти в базу или нет
func (uc *UserSignUpContinueRequest) Handle() bool {
	return uc.Avatar == "" && uc.Description == ""
}
