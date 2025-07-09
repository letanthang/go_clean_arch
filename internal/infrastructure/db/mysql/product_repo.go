package mysql

import (
	"app/internal/domain/entity"

	"gorm.io/gorm"
)

type ProductRepo struct {
	*BaseRepo[entity.Product]
}

func NewProductRepository(db *gorm.DB) ProductRepo {
	return ProductRepo{
		BaseRepo: &BaseRepo[entity.Product]{db: db},
	}
}
