package usecase

import (
	"context"
	"time"

	"github.com/rozy97/ecommerce-ifortepay/config"
	"github.com/rozy97/ecommerce-ifortepay/model"
)

type UserRepository interface {
	CountUserByEmail(ctx context.Context, email string) (int, error)
	CreateUser(ctx context.Context, user *model.User) (uint64, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
}

type ProductRepository interface {
	CreateProduct(ctx context.Context, product *model.Product) (uint64, error)
	GetProducts(ctx context.Context, limit, offset uint) (model.Products, error)
	GetProductsByIDs(ctx context.Context, IDs []uint64) (model.Products, error)
	GetProductByID(ctx context.Context, ID uint64) (*model.Product, error)
	GetProductPromotionsByProductIDs(ctx context.Context, ProductID []uint64) ([]model.ProductPromotion, error)
}

type CartRepository interface {
	CreateCart(ctx context.Context, cart *model.Cart) (uint64, error)
	GetCartByUserID(ctx context.Context, userID uint64) (*model.Cart, error)
	CreateCartItem(ctx context.Context, cartItem *model.CartItem) (uint64, error)
	GetCardItemsByCardID(ctx context.Context, cartID uint64, limit, offset uint) (model.CartItems, error)
}

type OrderRepository interface {
	CreateOrder(ctx context.Context, order *model.Order) (uint64, error)
	GetOrderByID(ctx context.Context, ID uint64) (*model.Order, error)
	GetOrdersByUserID(ctx context.Context, userID uint64) ([]model.Order, error)
	CreateOrderItem(ctx context.Context, orderItem *model.OrderItem)
	GetOrderItemsByOrderID(ctx context.Context, orderID uint64) ([]model.OrderItem, error)
}

type Clock interface {
	Now() time.Time
}

type UserUsecase struct {
	userRepository UserRepository
	clock          Clock
	env            *config.Environment
}

func NewUserUsecase(ur UserRepository, clock Clock, env *config.Environment) *UserUsecase {
	return &UserUsecase{userRepository: ur, clock: clock, env: env}
}

type ProductUsecase struct {
	productRepository ProductRepository
	clock             Clock
	env               config.Environment
}

func NewProductUsecase(pr ProductRepository, clock Clock, env config.Environment) *ProductUsecase {
	return &ProductUsecase{productRepository: pr, clock: clock, env: env}
}

type CartUsecase struct {
	cartRepository    CartRepository
	productRepository ProductRepository
	clock             Clock
	env               config.Environment
}

func NewCartUsecase(cr CartRepository, pr ProductRepository, clock Clock, env config.Environment) *CartUsecase {
	return &CartUsecase{cartRepository: cr, productRepository: pr, clock: clock, env: env}
}
