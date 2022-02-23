package application

import (
	"github.com/mahenrique94/products-catalog-go/src/domain"
	"github.com/mahenrique94/products-catalog-go/src/domain/dtos"
	"github.com/mahenrique94/products-catalog-go/src/infrastructure/storage/repositories"
)

type CreateAProductUseCase struct {
	Repository *repositories.ProductRepository
}

func NewCreateAProductUseCase(repository *repositories.ProductRepository) *CreateAProductUseCase {
	return &CreateAProductUseCase{Repository: repository}
}

func (c *CreateAProductUseCase) Execute(data dtos.ProductDTOIn) (entity dtos.ProductDTOOut, err error) {
	product, err := c.Repository.Create(domain.Product{
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
