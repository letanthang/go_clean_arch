package usecase

import (
	"app/internal/domain/entity"
	"app/internal/domain/repository"
	"app/internal/infrastructure/db/query_builder"
	"context"
)


type productListUsecase struct {
	productRepo repository.IProductRepo
}

func NewProductListUsecase(productRepo repository.IProductRepo) ProductListUsecase {
	return &productListUsecase{
		productRepo: productRepo,
	}
}

func (uc *productListUsecase) Do(ctx context.Context) ([]entity.Product, error) {
	return uc.productRepo.List(ctx, query_builder.NewProductAllQuery())
}
