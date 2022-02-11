package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mahenrique94/products-catalog-go/src/modules"
)

func main() {
	c := modules.CreateContainer()

	err := c.Invoke(func(r *gin.Engine, group *gin.RouterGroup) {
		r.Run(":8080")
	})

	if err != nil {
		panic(err)
	}
}
