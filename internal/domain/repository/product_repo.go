package repository

import (
	"app/internal/domain/entity"
)

type IProductRepo interface {
	IBaseRepo[entity.Product]
}
