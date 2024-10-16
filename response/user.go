package response

const (
	ErrCodeEmailAlreadyRegistered = 123

	ErrMessageEmailAlreadyRegistered = "email already registered"
)

type Register struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Login struct {
	AccessToken string `json:"access_token"`
}
