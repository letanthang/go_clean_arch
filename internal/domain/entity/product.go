package entity

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID    int64
	Name  string
	Price float64
	Stock int64
	SKU   string
}
