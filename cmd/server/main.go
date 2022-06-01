package main

import (
	"github.com/eduaraujogf/web-server/cmd/server/controllers"
	"github.com/eduaraujogf/web-server/internal/products"
	"github.com/gin-gonic/gin"
)

func main() {

	repo := products.NewRepository()

	service := products.NewService(repo)

	p := controllers.NewProduct(service)

	r := gin.Default()

	pr := r.Group("products")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	r.Run()

}
