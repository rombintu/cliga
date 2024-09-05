package storage

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	defaultSprintsCollectionName = "sprints"
)

type Collections struct {
	Sprints string
}

type Databases struct {
	Main string
}

type Meta struct {
	Databases   Databases
	Collections Collections
}

type mongodbDriver struct {
	Client *mongo.Client
	URI    string
	Meta   Meta
}

func NewMongoDBDriver(path string, database string) *mongodbDriver {
	return &mongodbDriver{
		URI: path,
		Meta: Meta{ // TODO
			Databases:   Databases{Main: database},
			Collections: Collections{Sprints: defaultSprintsCollectionName},
		},
	}
}

func (d *mongodbDriver) Open() error {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(d.URI))
	if err != nil {
		// logger.Log.Error("Failed to connect to MongoDB", slog.Any("error", err))
	}
	d.Client = client
	return nil
}

func (d *mongodbDriver) ConnSprints() *mongo.Collection {
	return d.Client.Database(d.Meta.Databases.Main).Collection(d.Meta.Collections.Sprints)
}

func (d *mongodbDriver) Close() error {
	return d.Client.Disconnect(context.Background())
}

type row interface{}

func (d *mongodbDriver) FetchOne(id int) (row, error) {
	sprints := d.ConnSprints()
	var result bson.M
	var err error
	err = sprints.FindOne(context.Background(), bson.D{{"id", id}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return result, err
	} else if err != nil {
		return result, err
	}
	return result, err
}

func (d *mongodbDriver) InsertOne() (row, error) {
	return "", nil
}
