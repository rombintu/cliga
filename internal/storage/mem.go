package storage

type data map[string]interface{}

type memDriver struct {
	Collections map[string]data
}

func NewMemDriver() *memDriver {
	return &memDriver{}
}

func (d *memDriver) Open() error {
	d.Collections = make(map[string]data)
	return nil
}

func (d *memDriver) Close() error {
	d.Collections = nil
	return nil
}

func (d *memDriver) FetchOne() error {
	return nil
}

func (d *memDriver) InsertOne() error {
	return nil
}
