package storage

import "time"

type Sprint struct {
	ID        int64     `bson:"id"`
	IsDone    bool      `bson:"is_done"`
	UpdatedAt time.Time `bson:"updated_at"`
}

type User struct {
	ID        int64     `bson:"id"`
	Login     string    `bson:"login"`
	Anchor    string    `bson:"anchor"`
	UpdatedAt time.Time `bson:"updated_at"`
	CreatedAt time.Time `bson:"created_at"`
	Sprints   []Sprint  `bson:"sprints"`
}
