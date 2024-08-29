package store

import (
	"errors"

	"github.com/awalker125/go-api/models/cars"
)

var (
	NotFoundErr = errors.New("not found")
)

type MemStore struct {
	list map[string]cars.Cars
}

func NewMemStore() *MemStore {
	list := make(map[string]cars.Cars)
	return &MemStore{
		list,
	}
}

func (m MemStore) Add(name string, car cars.Cars) error {
	m.list[name] = car
	return nil
}

func (m MemStore) Get(name string) (cars.Cars, error) {

	if val, ok := m.list[name]; ok {
		return val, nil
	}

	return cars.Cars{}, NotFoundErr
}

func (m MemStore) List() (map[string]cars.Cars, error) {
	return m.list, nil
}

func (m MemStore) Update(name string, car cars.Cars) error {

	if _, ok := m.list[name]; ok {
		m.list[name] = car
		return nil
	}

	return NotFoundErr
}

func (m MemStore) Remove(name string) error {
	delete(m.list, name)
	return nil
}
