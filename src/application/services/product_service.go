package services

import (
	"github.com/mahenrique94/products-catalog-go/src/domain"
	"github.com/mahenrique94/products-catalog-go/src/domain/dtos"
	"github.com/mahenrique94/products-catalog-go/src/infrastructure/storage/repositories"
)

type ProductService struct {
	Repository *repositories.ProductRepository
}

func NewProductService(repository *repositories.ProductRepository) *ProductService {
	return &ProductService{Repository: repository}
}

func (p *ProductService) Create(data dtos.ProductDTOIn) (entity dtos.ProductDTOOut, err error) {
	product, err := p.Repository.Create(domain.Product{
		Name:        data.Name,
		Description: data.Description,
		Price:       data.Price,
		Quantity:    data.Quantity,
	})

	if err != nil {
		return dtos.ProductDTOOut{}, err
	}

	return dtos.ProductDTOOut(product), nil
}

func (p *ProductService) Get(ID int64) (entity dtos.ProductDTOOut, err error) {
	product, err := p.Repository.Get(ID)
	return dtos.ProductDTOOut(product), err
}

func (p *ProductService) List() (collection []dtos.ProductDTOOut, err error) {
	var result []dtos.ProductDTOOut
	products, err := p.Repository.List()

	for _, product := range products {
		result = append(result, dtos.ProductDTOOut(product))
	}

	return result, err
}

func (p *ProductService) Remove(ID int64) (success bool, err error) {
	return p.Repository.Remove(ID)
}

func (p *ProductService) Update(data dtos.ProductDTOIn, ID int64) (entity dtos.ProductDTOOut, err error) {
	product, err := p.Repository.Update(domain.Product{
		ID:          uint(ID),
		Name:        data.Name,
		Description: data.Description,
		Price:       data.Price,
		Quantity:    data.Quantity,
	})

	if err != nil {
		return dtos.ProductDTOOut{}, err
	}

	return dtos.ProductDTOOut(product), nil
}
