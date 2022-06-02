package products

type Service interface {
	GetAll() ([]Product, error)
	Store(stock int, name, color, code, createdAt string, price float64, isPublished bool) (Product, error)
	Update(id, stock int, name, color, code, createdAt string, price float64, isPublished bool) (Product, error)
	UpdateNamePrice(id int, name string, price float64) (Product, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s service) GetAll() ([]Product, error) {
	ps, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return ps, nil
}

func (s service) Store(stock int, name, color, code, createdAt string, price float64, isPublished bool) (Product, error) {
	lastID, err := s.repository.LastID()

	if err != nil {
		return Product{}, nil
	}
	lastID++

	product, err := s.repository.Store(lastID, stock, name, color, code, createdAt, price, isPublished)

	if err != nil {
		return Product{}, err
	}

	return product, nil
}

func (s service) Update(id, stock int, name, color, code, createdAt string, price float64, isPublished bool) (Product, error) {
	product, err := s.repository.Update(id, stock, name, color, code, createdAt, price, isPublished)

	if err != nil {
		return Product{}, err
	}

	return product, err
}

func (s service) Delete(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return err
}

func (s service) UpdateNamePrice(id int, name string, price float64) (Product, error) {
	product, err := s.repository.UpdateNamePrice(id, name, price)
	if err != nil {
		return Product{}, err
	}
	return product, nil
}
