package service

import (
	"context"

	"github.com/betawulan/synapsis/model"
	"github.com/betawulan/synapsis/repository"
)

type onlineStoreService struct {
	onlineStoreRepo repository.OnlineStoreRepository
}

func (o onlineStoreService) Fetch(ctx context.Context, filter model.ProductCategoryFilter) (model.ProductCategoryResponse, error) {
	productCategories, err := o.onlineStoreRepo.Fetch(ctx, filter)
	if err != nil {
		return model.ProductCategoryResponse{}, err
	}

	return model.ProductCategoryResponse{Products: productCategories}, nil
}

func NewOnlineStoreService(onlineStoreRepo repository.OnlineStoreRepository) OnlineStoreService {
	return onlineStoreService{
		onlineStoreRepo: onlineStoreRepo,
	}
}
