package storage

import "context"

const (
	MemDriverName     = "mem"
	MongodbDriverName = "mongodb"
)

type Storage interface {
	Open(ctx context.Context) error
	Close(ctx context.Context) error

	UserFetch(login string) (User, error)
	UserUpsert(user User) error
}

func NewStorage(driver string, path string, database string) Storage {
	switch driver {
	case MongodbDriverName:
		return NewMongoDBDriver(path, database)
	default:
		return nil
		// return NewMemDriver()
	}
}
