package products

type Service interface {
	GetAll() ([]Product, error)
	Store(stock int, name, color, code, createdAt string, price float64, isPublished bool) (Product, error)
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
