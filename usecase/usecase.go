package usecase

import (
	"context"
	"time"

	"github.com/rozy97/ecommerce-ifortepay/config"
	"github.com/rozy97/ecommerce-ifortepay/model"
)

type UserRepository interface {
	CountUserByEmail(ctx context.Context, email string) (int, error)
	CreateUser(ctx context.Context, user *model.User) error
}

type Clock interface {
	Now() time.Time
}

type UserUsecase struct {
	userRepository UserRepository
	clock          Clock
	env            config.Environment
}

func NewUserUsecase(ur UserRepository, clock Clock, env config.Environment) *UserUsecase {
	return &UserUsecase{userRepository: ur, clock: clock, env: env}
}
