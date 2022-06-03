package products

import (
	"fmt"

	"github.com/eduaraujogf/web-server/pkg/store"
)

var ps []Product = []Product{}

var lastID int

type Repository interface {
	GetAll() ([]Product, error)
	Store(id, stock int, name, color, code, createdAt string, price float64, isPublished bool) (Product, error)
	LastID() (int, error)
	Update(id, stock int, name, color, code, createdAt string, price float64, isPublished bool) (Product, error)
	UpdateNamePrice(id int, name string, price float64) (Product, error)
	Delete(id int) error
}

type repository struct {
	db store.Store
}

func (r *repository) GetAll() ([]Product, error) {
	var ps []Product
	r.db.Read(&ps)
	return ps, nil
}

func (r *repository) LastID() (int, error) {
	var ps []Product
	if err := r.db.Read(&ps); err != nil {
		return 0, err
	}
	if len(ps) == 0 {
		return 0, nil
	}
	return ps[len(ps)-1].Id, nil
}

func (r *repository) Store(id, stock int, name, color, code, createdAt string, price float64, isPublished bool) (Product, error) {
	var ps []Product
	r.db.Read(&ps)
	p := Product{id, name, color, price, stock, code, isPublished, createdAt}
	ps = append(ps, p)

	if err := r.db.Write(ps); err != nil {
		return Product{}, err
	}

	return p, nil
}

func (r *repository) Update(id, stock int, name, color, code, createdAt string, price float64, isPublished bool) (Product, error) {
	var ps []Product
	r.db.Read(&ps)
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
	if err := r.db.Write(ps); err != nil {
		return Product{}, err
	}
	return p, nil
}

func (r *repository) Delete(id int) error {
	var ps []Product
	r.db.Read(&ps)

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
	if err := r.db.Write(ps); err != nil {
		return err
	}
	return nil
}

func (r *repository) UpdateNamePrice(id int, name string, price float64) (Product, error) {
	var ps []Product
	var p Product
	r.db.Read(&ps)
	updated := false
	for i := range ps {
		if ps[i].Id == id {
			ps[i].Name = name
			ps[i].Price = price
			updated = true
			p = ps[i]
		}
	}
	if !updated {
		return Product{}, fmt.Errorf("product with id %d not found", id)
	}
	if err := r.db.Write(ps); err != nil {
		return Product{}, err
	}
	return p, nil
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}
