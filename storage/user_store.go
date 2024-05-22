package storage

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserStorer interface {
	CreateUser(context.Context) error
}

type MongoUserStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoUserStore(client *mongo.Client, dbname string, collname string) *MongoUserStore {
	return &MongoUserStore{
		client: client,
		coll:   client.Database(dbname).Collection(collname),
	}
}

func (m *MongoUserStore) CreateUser(context.Context) error {
	return nil
}
