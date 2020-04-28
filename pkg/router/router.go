package router

import (
	"fmt"
	"net/http"

	"github.com/ks6088ts/rest-go/pkg/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// NewRouter returns a router
func NewRouter(port int, c *controller.Controller) (*Router, error) {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()
	r.Use(cors.Default())

	// Ping test: curl localhost:port/ping
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	p := r.Group("/products")
	{
		p.GET("/:id", c.ReadProduct)
		p.GET("", c.ReadProducts)
		p.POST("", c.CreateProduct)
	}

	return &Router{
		router: r,
		port:   port,
	}, nil
}

// Run attaches a router to http server
func (r *Router) Run() {
	// Listen and Server in 0.0.0.0:port
	r.router.Run(fmt.Sprintf(":%d", r.port))
}

// Router is a type definition for router
type Router struct {
	router *gin.Engine
	port   int
}
