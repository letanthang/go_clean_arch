package usecase

import (
	"app/internal/domain/entity"
	"app/internal/domain/repository"
	"app/internal/infrastructure/db/query_builder"
	"context"
)

type productUsecase struct {
	productRepo repository.IProductRepo
}

func NewProductGetUsecase(productRepo repository.IProductRepo) ProductGetUsecase {
	return &productUsecase{
		productRepo: productRepo,
	}
}

func (uc *productUsecase) Do(ctx context.Context, id int64) (*entity.Product, error) {
	product, err := uc.productRepo.Get(ctx, query_builder.NewProductByIDQuery(id))
	if err != nil {
		return nil, err
	}

	return &product, nil
}
