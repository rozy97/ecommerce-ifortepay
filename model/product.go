package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type Product struct {
	ID        uint64          `db:"id"`
	SKU       string          `db:"sku"`
	Name      string          `db:"name"`
	Price     decimal.Decimal `db:"price"`
	Quantity  uint            `db:"quantity"`
	CreatedAt time.Time       `db:"created_at"`
	UpdatedAt time.Time       `db:"updated_at"`
	DeletedAt *time.Time      `db:"deleted_at"`
}

type ProductPromotion struct {
	ID        uint64     `db:"product_id"`
	ProductID uint64     `db:"product_id"`
	Type      int        `db:"type"`
	Metadata  []byte     `db:"metadata"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}

type Products []Product

func (ps Products) GetProductByID(ID uint64) Product {
	for _, p := range ps {
		if p.ID == ID {
			return p
		}
	}

	return Product{}
}
