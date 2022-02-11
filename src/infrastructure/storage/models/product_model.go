package models

import (
	"gorm.io/gorm"
)

type ProductModel struct {
	gorm.Model
	Name        string  `gorm:"not null;size:120"`
	Description string  `gorm:"not null;size:255"`
	Price       float32 `gorm:"not null"`
	Quantity    uint16  `gorm:"not null"`
}

func (ProductModel) TableName() string {
	return "products"
}
