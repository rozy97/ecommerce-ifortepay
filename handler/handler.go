package handler

import (
	"context"

	"github.com/rozy97/ecommerce-ifortepay/request"
	"github.com/rozy97/ecommerce-ifortepay/response"
)

type UserUsecase interface {
	Register(ctx context.Context, req *request.Register) (*response.Register, error)
	Login(ctx context.Context, req *request.Login) (*response.Login, error)
}

type UserHandler struct {
	uu UserUsecase
}

func NewUserHandler(uu UserUsecase) *UserHandler {
	return &UserHandler{uu: uu}
}
