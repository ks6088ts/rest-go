package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
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

// GetProduct ...
func (c *Controller) GetProduct(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	p, err := c.Service.GetProduct(id)
	if err != nil {
		ctx.AbortWithStatus(404)
		fmt.Println(err)
	}
	ctx.JSON(200, p)
}

// CreateProduct ...
func (c *Controller) CreateProduct(ctx *gin.Context) {
	p, err := c.Service.CreateProduct(ctx)
	if err != nil {
		ctx.AbortWithStatus(404)
		fmt.Println(err)
	}
	ctx.JSON(201, p)
}
