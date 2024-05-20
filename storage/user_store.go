package storage

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserStorer interface {
	CreateUser(context.Context) error
}

type MongoStorage struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoStorage(client *mongo.Client, dbname string, collname string) *MongoStorage {
	return &MongoStorage{
		client: client,
		coll:   client.Database(dbname).Collection(collname),
	}
}

func (m *MongoStorage) CreateUser(context.Context) error {
	return nil
}
