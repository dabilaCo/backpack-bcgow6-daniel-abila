package products

type Service interface {
	GetAll() ([]request, error)
	Store(name, typeRequest string, amount int, price float64) (request, error)
	Update(id int, name, typeRequest string, amount int, price float64) (request, error)
	Delete(id int) error
	UpdateNameAndPrice(id int, name string, price float64)(request, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]request, error) {
	req, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (s *service) Store(name, typeRequest string, amount int, price float64) (request, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return request{}, err
	}
	lastID++

	req, err := s.repository.Store(lastID, name, typeRequest, amount, price)
	if err != nil {
		return request{}, err
	}
	return req, nil
}

func (s *service) Update(id int, name, typeRequest string, amount int, price float64) (request, error) {
	return s.repository.Update(id, name, typeRequest, amount, price)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}

func (s *service) UpdateNameAndPrice(id int, name string, price float64)(request, error) {
	return s.UpdateNameAndPrice(id, name, price)
}
