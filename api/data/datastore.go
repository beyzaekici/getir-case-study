package data

import (
	"getir-case/api/data/store"
)

type Store struct {
	store.Store
}

func New(holder store.Store) *Store {
	return &Store{
		holder,
	}
}
