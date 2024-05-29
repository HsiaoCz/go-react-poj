package storage

import (
	"context"
	"errors"

	"github.com/HsiaoCz/go-react-poj/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserStorer interface {
	CreateUser(context.Context, *types.User) (*types.User, error)
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

func (m *MongoUserStore) CreateUser(ctx context.Context, user *types.User) (*types.User, error) {
	var decodeu types.User
	filter := bson.D{{Key: "email", Value: user.Email}}
	if err := m.coll.FindOne(ctx, filter).Decode(&decodeu); err != mongo.ErrNoDocuments {
		return nil, errors.New("create user failed because the database has this record")
	}
	resp, err := m.coll.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	user.ID = resp.InsertedID.(primitive.ObjectID)
	return user, nil
}
