package usecase

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/rozy97/ecommerce-ifortepay/config"
	mocks "github.com/rozy97/ecommerce-ifortepay/mocks/usecase"
	"github.com/rozy97/ecommerce-ifortepay/model"
	"github.com/rozy97/ecommerce-ifortepay/request"
	"github.com/rozy97/ecommerce-ifortepay/response"
	"github.com/stretchr/testify/assert"
)

func Test_Register(t *testing.T) {
	mockUserRepository := mocks.NewUserRepository(t)
	mockClock := mocks.NewClock(t)
	ctx := context.Background()
	email := "john@mail.com"
	password := "12345"
	request := request.Register{Email: email, Password: password}

	env := config.Environment{
		SecretPassword: "secret",
	}

	currentTime := time.Now()

	t.Run("failed test case, error while count email in database", func(t *testing.T) {
		usecase := NewUserUsecase(mockUserRepository, mockClock, env)
		expectedErrorMessage := "error while execute query"
		mockUserRepository.EXPECT().CountUserByEmail(ctx, email).Return(0, fmt.Errorf("%s", expectedErrorMessage)).Once()
		result, err := usecase.Register(ctx, request)

		assert.Nil(t, result)
		assert.Equal(t, expectedErrorMessage, err.Error())
	})

	t.Run("failed test case, error email already registered", func(t *testing.T) {
		usecase := NewUserUsecase(mockUserRepository, mockClock, env)
		mockUserRepository.EXPECT().CountUserByEmail(ctx, email).Return(1, nil).Once()
		result, err := usecase.Register(ctx, request)

		assert.Equal(t, response.ErrCodeEmailAlreadyRegistered, result.Code)
		assert.Equal(t, response.ErrMessageEmailAlreadyRegistered, result.Message)
		assert.Equal(t, nil, err)
	})

	t.Run("failed test case, error while insert user", func(t *testing.T) {
		usecase := NewUserUsecase(mockUserRepository, mockClock, env)
		mockUserRepository.EXPECT().CountUserByEmail(ctx, email).Return(0, nil).Once()
		mockClock.EXPECT().Now().Return(currentTime).Once()
		expectedErrorMessage := "error while execute query"
		mockUserRepository.EXPECT().CreateUser(ctx, &model.User{
			Email:     email,
			Password:  config.GetMD5Hash(password + env.SecretPassword),
			CreatedAt: currentTime.UTC(),
			UpdatedAt: currentTime.UTC(),
			DeletedAt: nil,
		}).Return(fmt.Errorf("%s", expectedErrorMessage)).Once()
		result, err := usecase.Register(ctx, request)

		assert.Nil(t, result)
		assert.Equal(t, expectedErrorMessage, err.Error())
	})

	t.Run("success test case", func(t *testing.T) {
		usecase := NewUserUsecase(mockUserRepository, mockClock, env)
		mockUserRepository.EXPECT().CountUserByEmail(ctx, email).Return(0, nil).Once()
		mockClock.EXPECT().Now().Return(currentTime).Once()
		mockUserRepository.EXPECT().CreateUser(ctx, &model.User{
			Email:     email,
			Password:  config.GetMD5Hash(password + env.SecretPassword),
			CreatedAt: currentTime.UTC(),
			UpdatedAt: currentTime.UTC(),
			DeletedAt: nil,
		}).Return(nil).Once()
		_, err := usecase.Register(ctx, request)

		assert.Nil(t, err)
	})

}
