package repository

type Repository struct{}

func (r Repository) GetA() (string, error) {
	return "A", nil
}
func (r Repository) GetB() (string, error) {
	return "B", nil
}

func NewRepository() Repository {
	return Repository{}
}
