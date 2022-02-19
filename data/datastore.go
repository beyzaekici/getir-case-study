package data

import (
	"github.com/beyzaekici/getir-case-study/data/store"
)

type Store struct {
	store.Store
}

func New(holder store.Store) *Store {
	return &Store{
		holder,
	}
}
