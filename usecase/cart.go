package usecase

import (
	"context"
	"database/sql"
	"errors"

	"github.com/rozy97/ecommerce-ifortepay/config"
	"github.com/rozy97/ecommerce-ifortepay/model"
	"github.com/rozy97/ecommerce-ifortepay/request"
	"github.com/rozy97/ecommerce-ifortepay/response"
)

func (cu *CartUsecase) GetUserCartItems(ctx context.Context, userID uint64, page, size uint) ([]response.CartItem, error) {
	cart, err := cu.cartRepository.GetCartByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	if size == 0 {
		size = config.PAGINATION_DEFAULT_SIZE
	}

	offset := (page - 1) * size
	result := make([]response.CartItem, 0)
	cartItems, err := cu.cartRepository.GetCardItemsByCardID(ctx, cart.ID, size, offset)
	if err != nil {
		return result, err
	}

	if len(cartItems) == 0 {
		return result, nil
	}

	products, err := cu.productRepository.GetProductsByIDs(ctx, cartItems.GetProductIDs())
	if err != nil {
		return nil, err
	}

	for _, cartItem := range cartItems {
		product := products.GetProductByID(cartItem.ProductID)
		result = append(result, response.CartItem{
			ID: cartItem.ID,
			Product: response.Product{
				ID:    product.ID,
				SKU:   product.SKU,
				Name:  product.Name,
				Price: product.Price,
			},
			Quantity: cartItem.Quantity,
		})
	}

	return result, nil
}

func (cu *CartUsecase) CreateCartItem(ctx context.Context, userID uint64, cartItem request.CartItem) error {
	_, err := cu.productRepository.GetProductByID(ctx, cartItem.ProductID)
	if err != nil {
		return err
	}

	cart, err := cu.cartRepository.GetCartByUserID(ctx, userID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	currentTimeUTC := cu.clock.Now().UTC()
	if cart == nil {
		cartID, err := cu.cartRepository.CreateCart(ctx, &model.Cart{
			UserID:    userID,
			CreatedAt: currentTimeUTC,
			UpdatedAt: currentTimeUTC,
		})
		if err != nil {
			return err
		}

		cart = &model.Cart{ID: cartID}
	}

	_, err = cu.cartRepository.CreateCartItem(ctx, &model.CartItem{
		CartID:    cart.ID,
		ProductID: cartItem.ProductID,
		Quantity:  cartItem.Quantity,
		CreatedAt: currentTimeUTC,
		UpdatedAt: currentTimeUTC,
	})

	return err
}
