package request

type CartItem struct {
	ID        uint64 `json:"id"`
	ProductID uint64 `json:"product_id"`
	Quantity  uint   `json:"quantity"`
}
