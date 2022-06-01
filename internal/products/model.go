package products

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name" binding:"required"`
	Color       string  `json:"color" binding:"required"`
	Price       float64 `json:"price" `
	Stock       int     `json:"stock" `
	Code        string  `json:"code" binding:"required"`
	IsPublished bool    `json:"isPublished"`
	CreatedAt   string  `json:"createdAt" binding:"required"`
}
