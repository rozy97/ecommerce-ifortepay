package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type Order struct {
	ID        uint64          `db:"id"`
	UserID    uint64          `db:"user_id"`
	Amount    decimal.Decimal `db:"amount"`
	Status    int             `db:"status"`
	CreatedAt time.Time       `db:"created_at"`
	UpdatedAt time.Time       `db:"updated_at"`
	DeletedAt *time.Time      `db:"deleted_at"`
}

type OrderItem struct {
	ID        uint64     `db:"id"`
	OrderID   uint64     `db:"order_id"`
	ProductID uint64     `db:"product_id"`
	Quantity  uint       `db:"quantity"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}
