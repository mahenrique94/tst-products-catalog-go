package modules

import (
	"github.com/mahenrique94/products-catalog-go/src/application"
	"github.com/mahenrique94/products-catalog-go/src/externals"
	"github.com/mahenrique94/products-catalog-go/src/infrastructure/storage/repositories"
	"github.com/mahenrique94/products-catalog-go/src/interface/http/controllers"
	"github.com/mahenrique94/products-catalog-go/src/modules/factories"
	"go.uber.org/dig"
)

func CreateContainer() *dig.Container {
	c := dig.New()

	c.Provide(factories.RoutesFactory)
	c.Provide(externals.CreateAConnection)
	c.Provide(factories.ServiceFactory)

	c.Provide(repositories.NewProductRepository)

	c.Provide(application.NewCreateAProductUseCase)
	c.Provide(application.NewGetAProductUseCase)
	c.Provide(application.NewListProductsUseCase)
	c.Provide(application.NewRemoveAProductUseCase)
	c.Provide(application.NewUpdateAProductUseCase)

	c.Provide(controllers.NewProductController)

	return c
}
