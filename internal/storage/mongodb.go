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

type MongodbDriver struct {
	client   *mongo.Client
	db       *mongo.Database
	URI      string
	Meta     Meta
	location *time.Location
}

func NewMongoDBDriver(path string, database string) *MongodbDriver {
	location, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		slog.Error(err.Error())
	}
	return &MongodbDriver{
		URI: path,
		Meta: Meta{ // TODO
			Databases: Databases{Main: database},
			Collections: Collections{
				Sprints: defaultSprintsCollectionName,
				Users:   defaultUsersCollectionName,
			},
		},
		location: location,
	}
}

func (d *MongodbDriver) Open(ctx context.Context) error {
	slog.Debug("try connect to database", slog.String("path", d.URI))
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(d.URI).SetMaxPoolSize(20))
	if err != nil {
		slog.Error("Failed to connect to MongoDB", slog.Any("error", err))
	}
	d.client = client

	d.SwitchDatabase(d.Meta.Databases.Main)
	return nil
}

func (d *MongodbDriver) Close(ctx context.Context) error {
	err := d.client.Disconnect(ctx)
	if err != nil {
		return err
	}
	d.client = nil
	d.db = nil
	return err
}

func (d *MongodbDriver) SwitchDatabase(name string) {
	d.db = d.client.Database(name)
}

func (d *MongodbDriver) ConnSprints() *mongo.Collection {
	return d.db.Collection(d.Meta.Collections.Sprints)
}

func (d *MongodbDriver) ConnUsers() *mongo.Collection {
	return d.db.Collection(d.Meta.Collections.Users)
}

func (d *MongodbDriver) UserFetch(login string) (User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{"login": login}
	users := d.ConnUsers()
	var result User
	if err := users.FindOne(
		ctx, filter,
	).Decode(&result); err != nil {
		return User{}, err
	}
	return result, nil
}

func (d *MongodbDriver) UserLogin(user User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	users := d.ConnUsers()
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now
	user.Sprints = []Sprint{}

	users.InsertOne(
		ctx,
		user,
	)
	return nil
}

func (d *MongodbDriver) UserPushSprint(login string, sprint Sprint) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{"login": login}

	opts := options.Update()
	users := d.ConnUsers()
	update := bson.M{
		"$set": bson.M{
			"updated_at": time.Now(),
		},
		"$push": bson.M{
			"sprints": sprint,
		},
	}
	if _, err := users.UpdateOne(
		ctx,
		filter,
		update,
		opts,
	); err != nil {
		return err
	}
	return nil
}
