package storage

const (
	MemDriverName     = "mem"
	MongodbDriverName = "mongodb"
)

type Storage interface {
	Open() error
	Close() error

	FetchSprint(id int) (Sprint, error)
	InsertOne() (row, error)
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
