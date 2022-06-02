package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/eduaraujogf/web-server/internal/products"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ProductController struct {
	service products.Service
}

type ErrorMsg struct {
	Message string `json:"message"`
}

func errorHandler(fe validator.FieldError) string {
	return fmt.Sprintf("field %s is required", fe.Field())
}

func errorRequest(ctx *gin.Context, err error) bool {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]ErrorMsg, len(ve))
		for i, fe := range ve {
			out[i] = ErrorMsg{errorHandler(fe)}
		}
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": out,
		})
		return true
	}
	return true
}

func errorBindJson(ctx *gin.Context, req any) bool {
	if err := ctx.ShouldBindJSON(&req); err != nil {
		if errorRequest(ctx, err) {
			return true
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return true
	}
	return false
}

func isValidToken(ctx *gin.Context) bool {
	token := ctx.Request.Header.Get("token")
	if token != "123456" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		return true
	}
	return false
}

func NewProduct(p products.Service) *ProductController {
	return &ProductController{
		service: p,
	}
}

func (c *ProductController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if isValidToken(ctx) {
			return
		}
		p, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, p)
	}
}

func (c *ProductController) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if isValidToken(ctx) {
			return
		}
		var req request
		if errorBindJson(ctx, &req) {
			return
		}

		p, err := c.service.Store(req.Stock, req.Name, req.Color, req.Code, req.CreatedAt, req.Price, req.IsPublished)

		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, p)
	}
}

func (c *ProductController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if isValidToken(ctx) {
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid ID",
			})
			return
		}

		var req request
		if errorBindJson(ctx, &req) {
			return
		}
		p, err := c.service.Update(id, req.Stock, req.Name, req.Color, req.Code, req.CreatedAt, req.Price, req.IsPublished)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, p)
	}
}

func (c *ProductController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if isValidToken(ctx) {
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid ID",
			})
			return
		}
		err = c.service.Delete(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"data": fmt.Sprintf("The product with id %d was deleted", id),
		})
	}
}

func (c *ProductController) UpdateNamePrice() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if isValidToken(ctx) {
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid ID",
			})
			return
		}

		var req requestPatch
		if errorBindJson(ctx, &req) {
			return
		}
		product, err := c.service.UpdateNamePrice(id, req.Name, req.Price)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"message": product,
		})
	}
}

type requestPatch struct {
	Name  string  `json:"name" binding:"required"`
	Price float64 `json:"price" binding:"required"`
}
type request struct {
	Name        string  `json:"name" binding:"required"`
	Color       string  `json:"color" binding:"required"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Code        string  `json:"code" binding:"required"`
	IsPublished bool    `json:"isPublished"`
	CreatedAt   string  `json:"createdAt" binding:"required"`
}
