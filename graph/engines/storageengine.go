package engines

type StorageEngine interface {
	CreateIndexes()

	Close()

	Open(path string) error
}

type StorageEngineProvider struct {
}

func (sep StorageEngineProvider) CreateEngine() StorageEngine {
	db := boltdb{}
	return &db
}
