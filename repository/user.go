package repository

import (
	"context"

	"github.com/rozy97/ecommerce-ifortepay/model"
)

func (ur *UserRepository) CountUserByEmail(ctx context.Context, email string) (total int, err error) {
	err = ur.db.GetContext(ctx, &total, "SELECT COUNT(id) total FROM users WHERE email = $1 AND deleted_at IS NULL", email)
	return
}

func (ur *UserRepository) CreateUser(ctx context.Context, user *model.User) (id uint64, err error) {
	query := `INSERT INTO users(email, password, created_at, updated_at, deleted_at)
		VALUES ($1, $2, $3, $4, $5) RETURNING id;
	`
	err = ur.db.GetContext(ctx, &id, query, user.Email, user.Password, user.CreatedAt, user.UpdatedAt, user.DeletedAt)
	if err != nil {
		return 0, err
	}

	return
}

func (ur *UserRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	query := `
		SELECT 
			id, email, password
		FROM users
		WHERE email = $1
		AND deleted_at IS NULL 
	`
	var user model.User
	err := ur.db.GetContext(ctx, &user, query, email)
	return &user, err
}
