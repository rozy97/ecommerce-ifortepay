package response

type CartItem struct {
	ID       uint64  `json:"id"`
	Product  Product `json:"product"`
	Quantity uint    `json:"quantity"`
}
