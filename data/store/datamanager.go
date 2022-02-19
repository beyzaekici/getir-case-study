package store

// Store Dynamic code structure implemented to add new features in the future
type Store interface {
	SetKey(key string, value string) error
	GetKey(key string) (string, error)
}

// DataManager Dynamic code structure implemented to add new features in the future
type DataManager interface {
	Retrieve(input interface{}) (out interface{}, err error)
}
