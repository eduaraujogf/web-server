package products

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
