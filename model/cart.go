package model

import "time"

type Cart struct {
	ID        uint64     `db:"id"`
	UserID    uint64     `db:"user_id"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}

type CartItem struct {
	ID        uint64     `db:"id"`
	CartID    uint64     `db:"cart_id"`
	ProductID uint64     `db:"product_id"`
	Quantity  uint       `db:"quantity"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}

type CartItems []CartItem

func (ci CartItems) GetProductIDs() []uint64 {
	productIDs := make([]uint64, 0)
	for _, i := range ci {
		productIDs = append(productIDs, i.ProductID)
	}
	return productIDs
}
