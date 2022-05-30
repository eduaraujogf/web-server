package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Products struct {
	Products []Product `json:"products"`
}
type Product struct {
	Id        uint    `json:"id"`
	Name      string  `json:"name"`
	Color     string  `json:"color"`
	Price     float64 `json:"price"`
	Stock     uint    `json:"stock"`
	Code      string  `json:"code"`
	Published bool    `json:"published"`
	CreatedAt string  `json:"created-at"`
}

func readJsonFile() Products {
	jsonFile, err := os.Open("products.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValues, _ := ioutil.ReadAll(jsonFile)

	var p Products

	p = ParseJson(byteValues, &p)

	return p
}

func ParseJson(productsJson []byte, p *Products) Products {

	if err := json.Unmarshal(productsJson, &p); err != nil {
		fmt.Printf("error:%v", err)
	}
	return *p
}

func helloHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, Eduardo!",
	})
}

func GetAllHandler(c *gin.Context) {
	p := readJsonFile()
	c.JSON(http.StatusOK, p)
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
	router.GET("/", GetAllHandler)
	router.Run()
}
