package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rozy97/ecommerce-ifortepay/config"
	"github.com/rozy97/ecommerce-ifortepay/model"
	"github.com/rozy97/ecommerce-ifortepay/request"
	"github.com/rozy97/ecommerce-ifortepay/response"
)

func (u *UserUsecase) Register(ctx context.Context, req *request.Register) (*response.Register, error) {
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

	_, err = u.userRepository.CreateUser(ctx, &model.User{
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

func (u *UserUsecase) Login(ctx context.Context, req *request.Login) (*response.Login, error) {
	user, err := u.userRepository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if config.GetMD5Hash(req.Password+u.env.SecretPassword) != user.Password {
		return nil, fmt.Errorf("password incorrect")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id": user.ID,
			"nbf":     time.Now().Unix(),
			"exp":     time.Now().Add(time.Hour * 5).Unix(),
		})

	tokenString, err := token.SignedString([]byte(u.env.SecretKey))
	if err != nil {
		return nil, err
	}

	return &response.Login{AccessToken: tokenString}, nil
}
