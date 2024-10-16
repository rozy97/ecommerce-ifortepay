package middleware

import "github.com/rozy97/ecommerce-ifortepay/config"

type middleware struct {
	env *config.Environment
}

func NewMiddleware(env *config.Environment) *middleware {
	return &middleware{env: env}
}
