package foo

type repository interface {
	GetA() (string, error)
	GetB() (string, error)
}

type FooService struct {
	repo repository
}

type FooResponse struct {
	FieldA string
	FieldB string
}

func NewFoo(repo repository) *FooService {
	return &FooService{repo}
}

func (s FooService) DoSometingSimple() (FooResponse, error) {

	a, err := s.repo.GetA()
	if err != nil {
		return FooResponse{}, err
	}

	b, err := s.repo.GetB()
	if err != nil {
		return FooResponse{}, err
	}
	return FooResponse{a, b}, nil
}
