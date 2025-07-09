package query_builder

import (
	"app/internal/domain/entity"

	"gorm.io/gorm"
)

type productById struct {
	id         int64
	preloaders []string
}

func NewProductByIDQuery(id int64, preloaders ...string) I {
	return productById{
		id:         id,
		preloaders: preloaders,
	}
}

func (builder productById) Query(db *gorm.DB) *gorm.DB {
	db = db.Model(&entity.Product{}).Where("id = ?", builder.id)

	for _, preloader := range builder.preloaders {
		db = db.Preload(preloader)
	}

	return db
}
