package products

var ps []Product = []Product{}

var lastID int

type Repository interface {
	GetAll() ([]Product, error)
	Store(id, stock int, name, color, code, createdAt string, price float64, isPublished bool) (Product, error)
	LastID() (int, error)
	// Update(id, stock int, name, color, code, createdAt string, price float64, isPublished bool) (Product, error)
	// UpdateName(id int, name string) (Product, error)
	// Delete(id int) error
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

func NewRepository() Repository {
	return &repository{}
}
