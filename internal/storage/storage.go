package storage

const (
	MemDriverName     = "mem"
	MongodbDriverName = "mongodb"
)

func NewStorage(driver string, path string, database string) *MongodbDriver {
	return NewMongoDBDriver(path, database)
}
