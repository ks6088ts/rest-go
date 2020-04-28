package controller

import (
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ks6088ts/rest-go/pkg/entity"
	"github.com/ks6088ts/rest-go/pkg/service"
)

// Controller ...
type Controller struct {
	Service *service.Service
}

// NewController creates a service
func NewController(service *service.Service) (*Controller, error) {
	return &Controller{
		Service: service,
	}, nil
}

// ReadProduct ...
func (c *Controller) ReadProduct(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	p, err := c.Service.ReadProduct(id)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	}
	ctx.JSON(http.StatusOK, p)
}

// ReadProducts ...
func (c *Controller) ReadProducts(ctx *gin.Context) {
	p, err := c.Service.ReadProducts()
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	}
	ctx.JSON(http.StatusOK, p)
}

// CreateProduct ...
func (c *Controller) CreateProduct(ctx *gin.Context) {
	var p entity.Product
	if err := ctx.BindJSON(&p); err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	}

	created, err := c.Service.CreateProduct(&p)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	}
	ctx.JSON(http.StatusCreated, created)
}
