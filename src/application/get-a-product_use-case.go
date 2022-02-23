package application

import (
	"github.com/mahenrique94/products-catalog-go/src/domain/dtos"
	"github.com/mahenrique94/products-catalog-go/src/infrastructure/storage/repositories"
)

type GetAProductUseCase struct {
	Repository *repositories.ProductRepository
}

func NewGetAProductUseCase(repository *repositories.ProductRepository) *GetAProductUseCase {
	return &GetAProductUseCase{Repository: repository}
}

func (g *GetAProductUseCase) Execute(ID int64) (entity dtos.ProductDTOOut, err error) {
	product, err := g.Repository.Get(ID)
	return dtos.ProductDTOOut(product), err
}
