package models

type User struct {
	ID          int     `json:"id"`
	Firstname   string  `json:"firstname"`
	Lastname    string  `json:"lastname"`
	Username    string  `json:"username"`
	Email       string  `json:"email"`
	Password    *string `json:"password,omitempty"`
	Description string  `json:"description"`
	Age         int     `json:"age"`
	AvatarUrl   string  `json:"avatarUrl"`
}

type Client struct {
	User
}

type Psychologist struct {
	User
	BusyMode string   `json:"busyMode"`
	Rate     float64  `json:"rate"`
	Reviews  []Review `json:"reviews,omitempty"`
}

type Review struct {
	User
	Anonym  bool   `json:"anonym"`
	Rating  int    `json:"Rating"`
	Comment string `json:"comment"`
}

type ErrorResponse struct {
	ErrorMessage string `json:"errorMessage"`
	ErrorCode    int    `json:"errorCode"`
}

type SessionResponse struct {
	ID     int    `json:"id"`
	Cookie string `json:"cookie"`
}

type ContinueSignUp struct {
	UserID      int
	Description string
	Avatar      string
}

type Room struct {
	Id       string `json:"roomId"`
	ClientId int    `json:"clientId"`
	PsychoId int    `json:"psychoId"`
}
