package storage

import (
	"context"
	"log/slog"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	defaultSprintsCollectionName = "sprints"
	defaultUsersCollectionName   = "users"
)

type Collections struct {
	Sprints string
	Users   string
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
	slog.Debug("try connect to database", slog.String("path", d.URI))
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(d.URI))
	if err != nil {
		slog.Error("Failed to connect to MongoDB", slog.Any("error", err))
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

func (d *mongodbDriver) FetchSprint(id int) (jsonData Sprint, err error) {
	sprints := d.ConnSprints()
	var result bson.M
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = sprints.FindOne(
		ctx,
		bson.M{"id": id},
	).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return jsonData, err
	} else if err != nil {
		return jsonData, err
	}
	binData, err := bson.Marshal(result)
	if err != nil {
		return jsonData, err
	}
	err = bson.Unmarshal(binData, &jsonData)
	return jsonData, err
}

func (d *mongodbDriver) InsertOne() (row, error) {
	return "", nil
}
