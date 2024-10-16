package usecase

import (
	"context"

	"github.com/rozy97/ecommerce-ifortepay/model"
	"github.com/rozy97/ecommerce-ifortepay/response"
)

func (p *ProductUsecase) CreateProduct(ctx context.Context, product *model.Product) error {
	_, err := p.productRepository.CreateProduct(ctx, product)
	return err
}

func (p *ProductUsecase) GetProducts(ctx context.Context, page, size uint) ([]response.Product, error) {
	offset := (page - 1) * size
	result := make([]response.Product, 0)
	products, err := p.productRepository.GetProducts(ctx, size, offset)
	if err != nil {
		return result, err
	}

	for _, product := range products {
		result = append(result, response.Product{
			ID:       product.ID,
			SKU:      product.SKU,
			Name:     product.Name,
			Price:    product.Price,
			Quantity: &product.Quantity,
		})
	}

	return result, nil
}
