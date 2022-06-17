package service

type Psychologyst interface {
	SignUp()
}

type Client interface {
	SignUp()
}

type Service struct {
	Psychologyst
	Client
}
