package response

import "fmt"

const (
	ErrCodeEmailAlreadyRegistered = 123

	ErrMessageEmailAlreadyRegistered = "email already registered"
)

var (
	ErrEmailAlreadyRegistered = fmt.Errorf("%s", ErrMessageEmailAlreadyRegistered)
)

type Register struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Login struct {
	AccessToken string `json:"access_token"`
}
