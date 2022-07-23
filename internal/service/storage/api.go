package storage

type Item interface {
	any
}

type StorageService[T Item] interface {
	Add(key string, value T) (T, error)
	Delete(key string) (T, error)
}

type storageService[T Item] struct {
}

func (s *storageService[T]) Add(key string, value T) (T, error) {
	return value, nil
}

func (s *storageService[T]) Delete(key string) error {
	return nil
}
