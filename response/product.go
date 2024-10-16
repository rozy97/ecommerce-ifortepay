package response

import "github.com/shopspring/decimal"

type Product struct {
	ID       uint64          `json:"id"`
	SKU      string          `json:"sku"`
	Name     string          `json:"name"`
	Price    decimal.Decimal `json:"price"`
	Quantity *uint           `json:"quantity,omitempty"`
}
