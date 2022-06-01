package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

const FileName = "products.json"

var lastID int = 3
var products Products = readJsonFile("./products.json")

type Products []Product

// type Products struct {
// 	Products []Product `json:"products"`
// }

// type Products []Product
type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name" binding:"required"`
	Color       string  `json:"color" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	Stock       int     `json:"stock" binding:"required"`
	Code        string  `json:"code" binding:"required"`
	IsPublished bool    `json:"isPublished" binding:"required"`
	CreatedAt   string  `json:"createdAt" binding:"required"`
}

type ErrorMsg struct {
	Message string `json:"message"`
}

func isValidToken(c *gin.Context) bool {
	token := c.Request.Header.Get("token")
	if token != "123456" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		return true
	}
	return false
}

func errorHandler(fe validator.FieldError) string {
	return fmt.Sprintf("field %s is required", fe.Field())
}

func readJsonFile(fileName string) (p Products) {

	jsonFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	fileByteValues, _ := ioutil.ReadAll(jsonFile)

	p = ParseObject(fileByteValues)

	return p
}

func ParseObject(fileByteValues []byte) (p Products) {

	if err := json.Unmarshal(fileByteValues, &p); err != nil {
		fmt.Printf("error:%v", err)
	}
	return p
}

func helloHandler(c *gin.Context) {
	if isValidToken(c) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, Eduardo!",
	})
}

func GetName(c *gin.Context) {
	if isValidToken(c) {
		return
	}
	name := c.Query("name")
	if name == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "name could not be empty",
		})
		return
	}

	// products := readJsonFile("./products.json")
	for _, product := range products {

		if product.Name == name {
			c.JSON(http.StatusOK, gin.H{
				"data": product,
			})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"message": "product not found",
	})
}

func GetId(c *gin.Context) {
	if isValidToken(c) {
		return
	}
	id := c.Param("id")
	convertedId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "id is not a number",
		})
		return
	}
	for _, product := range products {

		if product.Id == convertedId {
			c.JSON(http.StatusOK, gin.H{
				"data": product,
			})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"message": "product not found",
	})

}

func GetAllProducts(c *gin.Context) {
	// c.JSON(http.StatusOK, readJsonFile(FileName))
	// c.File("./products.json")

	if isValidToken(c) {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": products,
	})
}

func CreateProduct(c *gin.Context) {
	if isValidToken(c) {
		return
	}
	var product Product
	if err := c.ShouldBindJSON(&product); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = ErrorMsg{errorHandler(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": out,
			})
		}
		// c.JSON(http.StatusBadRequest, gin.H{
		// 	"error": err.Error(),
		// })
		return
	}

	lastID++
	product.Id = lastID
	products = append(products, product)
	c.JSON(http.StatusOK, gin.H{
		"data": product,
	})
}

func main() {
	gin.SetMode("debug")

	// f, err := os.Create("gin.log")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// gin.DefaultWriter = io.MultiWriter(f)

	// router := gin.Default()

	// router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

	// 	// your custom format
	// 	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
	// 		param.ClientIP,
	// 		param.TimeStamp.Format(time.RFC1123),
	// 		param.Method,
	// 		param.Path,
	// 		param.Request.Proto,
	// 		param.StatusCode,
	// 		param.Latency,
	// 		param.Request.UserAgent(),
	// 		param.ErrorMessage,
	// 	)
	// }))

	router := gin.Default()

	router.GET("/hello", helloHandler)
	group := router.Group("/products")
	{
		group.GET("/", GetAllProducts)
		group.GET("/:id", GetId)
		group.GET("/filter", GetName)
		group.POST("/", CreateProduct)
	}

	router.Run()
}
