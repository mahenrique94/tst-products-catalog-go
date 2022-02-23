package application

import (
	"github.com/mahenrique94/products-catalog-go/src/infrastructure/storage/repositories"
)

type RemoveAProductUseCase struct {
	Repository *repositories.ProductRepository
}

func NewRemoveAProductUseCase(repository *repositories.ProductRepository) *RemoveAProductUseCase {
	return &RemoveAProductUseCase{Repository: repository}
}

func (r *RemoveAProductUseCase) Execute(ID int64) (success bool, err error) {
	return r.Repository.Remove(ID)
}
