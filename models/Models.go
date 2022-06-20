package models

type User struct {
	ID        int     `json:"id"`
	Firstname string  `json:"firstname"`
	Lastname  string  `json:"lastname"`
	Username  string  `json:"username"`
	Email     string  `json:"email"`
	Password  *string `json:"password,omitempty"`
	Age       int     `json:"age"`
	AvatarUrl string  `json:"avatarUrl"`
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
	ClientId       int    `json:"clientID"`
	ClientUsername string `json:"clientUserName"`
	Rating         int    `json:"Rating"`
	Comment        string `json:"comment"`
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
	ID          int
	Description string
	Avatar      string
}
