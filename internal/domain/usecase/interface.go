package usecase

import (
	"app/internal/domain/entity"
	"context"
)

type ProductGetUsecase interface {
	Do(ctx context.Context, id int64) (*entity.Product, error)
}

type ProductListUsecase interface {
	Do(ctx context.Context) ([]entity.Product, error)
}
