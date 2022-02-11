package repositories

import (
	"github.com/mahenrique94/products-catalog-go/src/domain"
	"github.com/mahenrique94/products-catalog-go/src/infrastructure/storage/models"
	"gorm.io/gorm"
)

type ProductRepository struct {
	Database *gorm.DB
}

func NewProductRepository(database *gorm.DB) *ProductRepository {
	return &ProductRepository{Database: database}
}

func (p *ProductRepository) Create(product domain.Product) (entitiy domain.Product, err error) {
	model := models.ProductModel{
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Quantity:    product.Quantity,
	}
	result := p.Database.Create(&model)

	return domain.Product{
		ID:          model.ID,
		Name:        model.Name,
		Description: model.Description,
		Price:       model.Price,
		Quantity:    model.Quantity,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
		DeletedAt:   model.DeletedAt.Time,
	}, result.Error
}

func (p *ProductRepository) Get(ID int64) (entity domain.Product, err error) {
	var product domain.Product
	result := p.Database.First(&product, ID)
	return product, result.Error
}

func (p *ProductRepository) List() (collection []domain.Product, err error) {
	var products []domain.Product
	result := p.Database.Find(&products)
	return products, result.Error
}

func (p *ProductRepository) Remove(ID int64) (success bool, err error) {
	p.Database.Delete(&domain.Product{}, ID)
	return true, nil
}

func (p *ProductRepository) Update(product domain.Product) (entitiy domain.Product, err error) {
	model := models.ProductModel{
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Quantity:    product.Quantity,
	}
	model.ID = product.ID
	result := p.Database.Model(&model).Updates(model)

	return domain.Product{
		ID:          model.ID,
		Name:        model.Name,
		Description: model.Description,
		Price:       model.Price,
		Quantity:    model.Quantity,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
		DeletedAt:   model.DeletedAt.Time,
	}, result.Error
}
