package usecase

import (
	"context"

	"github.com/rozy97/ecommerce-ifortepay/config"
	"github.com/rozy97/ecommerce-ifortepay/model"
	"github.com/rozy97/ecommerce-ifortepay/request"
	"github.com/rozy97/ecommerce-ifortepay/response"
)

func (u *UserUsecase) Register(ctx context.Context, req request.Register) (*response.Register, error) {
	total, err := u.userRepository.CountUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if total > 0 {
		return &response.Register{
			Code:    response.ErrCodeEmailAlreadyRegistered,
			Message: response.ErrMessageEmailAlreadyRegistered,
		}, nil
	}

	currentTimeUTC := u.clock.Now().UTC()

	err = u.userRepository.CreateUser(ctx, &model.User{
		Email:     req.Email,
		Password:  config.GetMD5Hash(req.Password + u.env.SecretPassword),
		CreatedAt: currentTimeUTC,
		UpdatedAt: currentTimeUTC,
		DeletedAt: nil,
	})

	if err != nil {
		return nil, err
	}

	return &response.Register{
		Message: "register success",
	}, nil
}
