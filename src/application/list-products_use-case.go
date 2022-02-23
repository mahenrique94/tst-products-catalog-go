package application

import (
	"github.com/mahenrique94/products-catalog-go/src/domain/dtos"
	"github.com/mahenrique94/products-catalog-go/src/infrastructure/storage/repositories"
)

type ListProductsUseCase struct {
	Repository *repositories.ProductRepository
}

func NewListProductsUseCase(repository *repositories.ProductRepository) *ListProductsUseCase {
	return &ListProductsUseCase{Repository: repository}
}

func (l *ListProductsUseCase) Execute() (collection []dtos.ProductDTOOut, err error) {
	var result []dtos.ProductDTOOut
	products, err := l.Repository.List()

	for _, product := range products {
		result = append(result, dtos.ProductDTOOut(product))
	}

	if result == nil {
		return []dtos.ProductDTOOut{}, nil
	}

	return result, err
}
