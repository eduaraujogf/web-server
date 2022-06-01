package products

import "fmt"

var ps []Product = []Product{}

var lastID int

type Repository interface {
	GetAll() ([]Product, error)
	Store(id, stock int, name, color, code, createdAt string, price float64, isPublished bool) (Product, error)
	LastID() (int, error)
	Update(id, stock int, name, color, code, createdAt string, price float64, isPublished bool) (Product, error)
	// UpdateName(id int, name string) (Product, error)
	Delete(id int) error
}

type repository struct{}

func (repository) GetAll() ([]Product, error) {
	return ps, nil
}

func (repository) LastID() (int, error) {
	return lastID, nil
}

func (repository) Store(id, stock int, name, color, code, createdAt string, price float64, isPublished bool) (Product, error) {
	p := Product{id, name, color, price, stock, code, isPublished, createdAt}
	ps = append(ps, p)
	lastID = p.Id
	return p, nil
}

func (repository) Update(id, stock int, name, color, code, createdAt string, price float64, isPublished bool) (Product, error) {
	p := Product{Name: name, Color: color, Price: price, Stock: stock, Code: code, IsPublished: isPublished, CreatedAt: createdAt}
	updated := false
	for i := range ps {
		if ps[i].Id == id {
			p.Id = id
			ps[i] = p
			updated = true
		}
	}
	if !updated {
		return Product{}, fmt.Errorf("product with id %d not found", id)
	}
	fmt.Println(p)
	return p, nil
}

func (repository) Delete(id int) error {
	deleted := false
	var index int
	for i := range ps {
		if ps[i].Id == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("product with id %d not found", id)
	}
	ps = append(ps[:index], ps[index+1:]...)
	return nil
}

func NewRepository() Repository {
	return &repository{}
}
