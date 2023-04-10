package main

type model struct {
	ID   int    "json:`id`"
	Name string "json:`name`"
}

type Service interface {
	add(name string) error
	remove(id int) error
	getall() ([]model, error)
}

type scv struct{}

func NewService() Service {
	return &scv{}
}

func (s *scv) add(name string) error {
	return nil
}

func (s *scv) remove(id int) error {
	return nil
}

func (s *scv) getall() ([]model, error) {
	return []model{}, nil
}
