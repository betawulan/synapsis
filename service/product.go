package service

import (
	"context"

	"github.com/betawulan/synapsis/model"
	"github.com/betawulan/synapsis/repository"
)

type productService struct {
	productRepo repository.ProductRepository
}

func (p productService) Fetch(ctx context.Context, filter model.ProductCategoryFilter) ([]model.ProductCategory, error) {
	productCategories, err := p.productRepo.Fetch(ctx, filter)
	if err != nil {
		return nil, err
	}

	return productCategories, nil
}

func NewProductService(productRepo repository.ProductRepository) ProductService {
	return productService{
		productRepo: productRepo,
	}
}
