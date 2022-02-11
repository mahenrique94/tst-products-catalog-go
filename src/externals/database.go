package externals

import (
	"errors"

	"github.com/mahenrique94/products-catalog-go/src/infrastructure/storage/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func CreateAConnection() (database *gorm.DB, err error) {
	databaseUri := "root:root@tcp(127.0.0.1:3306)/products_catalog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(databaseUri), &gorm.Config{})

	if err != nil {
		return nil, errors.New("I can't connect to mysql database")
	}

	db.AutoMigrate(&models.ProductModel{})

	return db, nil
}
