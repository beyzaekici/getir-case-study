package data

import (
	"getir-case/api/store"
)

type DataStore struct {
	store.Store
}

func New(holder store.Store) *DataStore {
	return &DataStore{
		holder,
	}
}
