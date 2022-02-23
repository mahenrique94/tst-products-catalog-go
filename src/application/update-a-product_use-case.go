package application

import (
	"github.com/mahenrique94/products-catalog-go/src/domain"
	"github.com/mahenrique94/products-catalog-go/src/domain/dtos"
	"github.com/mahenrique94/products-catalog-go/src/infrastructure/storage/repositories"
)

type UpdateAProductUseCase struct {
	Repository *repositories.ProductRepository
}

func NewUpdateAProductUseCase(repository *repositories.ProductRepository) *UpdateAProductUseCase {
	return &UpdateAProductUseCase{Repository: repository}
}

func (u *UpdateAProductUseCase) Execute(data dtos.ProductDTOIn, ID int64) (entity dtos.ProductDTOOut, err error) {
	product, err := u.Repository.Update(domain.Product{
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
