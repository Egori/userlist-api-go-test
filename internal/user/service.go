package user

type Service interface {
	GetAll() ([]User, error)
	Create(user *User) error
	Update(user *User) error
	Delete(id uint) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetAll() ([]User, error) {
	return s.repo.FindAll()
}

func (s *service) Create(user *User) error {
	return s.repo.Create(user)
}

func (s *service) Update(user *User) error {
	return s.repo.Update(user)
}

func (s *service) Delete(id uint) error {
	return s.repo.Delete(id)
}
