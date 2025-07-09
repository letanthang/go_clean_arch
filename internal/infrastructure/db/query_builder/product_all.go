package query_builder

import (
	"app/internal/domain/entity"

	"gorm.io/gorm"
)

type productAll struct{}

func NewProductAllQuery() I {
	return productAll{}
}

func (builder productAll) Query(db *gorm.DB) *gorm.DB {
	return db.Model(&entity.Product{})
}
