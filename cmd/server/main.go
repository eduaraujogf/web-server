package main

import (
	"log"

	"github.com/eduaraujogf/web-server/cmd/server/controllers"
	"github.com/eduaraujogf/web-server/internal/products"
	"github.com/eduaraujogf/web-server/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error to load .env")
	}
	db := store.New(store.FileType, "./products.json")

	repo := products.NewRepository(db)

	service := products.NewService(repo)

	p := controllers.NewProduct(service)

	r := gin.Default()

	pr := r.Group("products")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	pr.PUT("/:id", p.Update())
	pr.PATCH("/:id", p.UpdateNamePrice())
	pr.DELETE("/:id", p.Delete())
	r.Run()

}
