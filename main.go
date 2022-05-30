package main

import (
	"github.com/gin-gonic/gin"
)

const FileName = "products.json"

type Products struct {
	Products []Product `json:"products"`
}
type Product struct {
	Id          uint    `json:"id"`
	Name        string  `json:"name"`
	Color       string  `json:"color"`
	Price       float64 `json:"price"`
	Stock       uint    `json:"stock"`
	Code        string  `json:"code"`
	IsPublished bool    `json:"isPublished"`
	CreatedAt   string  `json:"createdAt"`
}

// func readJsonFile(fileName string) Products {

// 	jsonFile, err := os.Open(fileName)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	defer jsonFile.Close()

// 	fileByteValues, _ := ioutil.ReadAll(jsonFile)

// 	p := ParseObject(fileByteValues)

// 	return p
// }

// func ParseObject(fileByteValues []byte) (p Products) {

// 	if err := json.Unmarshal(fileByteValues, &p); err != nil {
// 		fmt.Printf("error:%v", err)
// 	}
// 	return p
// }

func helloHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, Eduardo!",
	})
}

func GetAllHandler(c *gin.Context) {
	// c.JSON(http.StatusOK, readJsonFile(FileName))
	c.File("./products.json")
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
	router.GET("/products", GetAllHandler)
	router.Run()
}
