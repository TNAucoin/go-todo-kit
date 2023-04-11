package main

import "math/rand"

var store []model

type model struct {
	ID   int    "json:`id`"
	Name string "json:`name`"
}

type Service interface {
	add(name string) (model, error)
	remove(id int) error
	getall() ([]model, error)
}

type scv struct{}

func NewService() Service {
	return &scv{}
}

func (s *scv) add(name string) (model, error) {
	id := rand.Intn(1000)
	item := model{
		ID:   id,
		Name: name,
	}
	store = append(store, item)
	return item, nil
}

func (s *scv) remove(id int) error {
	if len(store) == 1 {
		if store[0].ID == id {
			store = []model{}
			return nil
		}
	}
	for i := len(store) - 1; i >= 0; i-- {
		if store[i].ID == id {
			store = append(store[:i], store[i+1:]...)
			return nil
		}
	}
	return nil
}

func (s *scv) getall() ([]model, error) {
	return store, nil
}
