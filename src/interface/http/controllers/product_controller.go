package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mahenrique94/products-catalog-go/src/application/services"
	"github.com/mahenrique94/products-catalog-go/src/domain/dtos"
)

type ProductController struct {
	Service *services.ProductService
}

func NewProductController(service *services.ProductService) *ProductController {
	return &ProductController{Service: service}
}

func (p *ProductController) Create(c *gin.Context) {
	var productIn dtos.ProductDTOIn

	if err := c.ShouldBindJSON(&productIn); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	product, err := p.Service.Create(productIn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"entitiy": product,
	})
}

func (p *ProductController) Get(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 0, 64)
	product, err := p.Service.Get(id)

	if product.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Product with id " + strconv.FormatInt(id, 10) + " not found",
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"entitiy": product,
	})
}

func (p *ProductController) List(c *gin.Context) {
	products, err := p.Service.List()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"collection": products,
	})
}

func (p *ProductController) Remove(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 0, 64)
	product, err := p.Service.Get(id)

	if product.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Product with id " + strconv.FormatInt(id, 10) + " not found",
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	success, _ := p.Service.Remove(id)
	if success {
		c.JSON(http.StatusOK, gin.H{
			"message": "Product remove with success",
		})
		return
	}
}

func (p *ProductController) Update(c *gin.Context) {
	var productIn dtos.ProductDTOIn
	id, _ := strconv.ParseInt(c.Param("id"), 0, 64)
	product, _ := p.Service.Get(id)

	if product.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Product with id " + strconv.FormatInt(id, 10) + " not found",
		})
		return
	}

	if err := c.ShouldBindJSON(&productIn); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	product, err := p.Service.Update(productIn, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"entitiy": product,
	})
}
