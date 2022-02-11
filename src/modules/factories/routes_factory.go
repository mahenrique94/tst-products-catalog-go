package factories

import (
	"github.com/gin-gonic/gin"
	"github.com/mahenrique94/products-catalog-go/src/interface/http/controllers"
)

func RoutesFactory(router *gin.Engine, productController *controllers.ProductController) *gin.RouterGroup {
	api := router.Group("/api")
	v1 := api.Group("/v1")

	v1.DELETE("/products/:id", productController.Remove)
	v1.GET("/products", productController.List)
	v1.GET("/products/:id", productController.Get)
	v1.POST("/products", productController.Create)
	v1.PUT("/products/:id", productController.Update)

	return v1
}
