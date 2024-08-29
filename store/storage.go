package store

import "github.com/awalker125/go-api/models/cars"

type CarStore interface {
	Add(name string, car cars.Cars) error
	Get(name string) (cars.Cars, error)
	Update(name string, car cars.Cars) error
	List() (map[string]cars.Cars, error)
	Remove(name string) error
}
