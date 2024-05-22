package storage

import (
	"context"

	"github.com/HsiaoCz/go-react-poj/types"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoStorer interface {
	GetTodoByID(context.Context, string) (*types.Todo, error)
	GetTodoList(context.Context, string) ([]*types.Todo, error)
	CreateTodo(context.Context, *types.Todo) (*types.Todo, error)
	UpdateTodoByID(context.Context, string) (*types.Todo, error)
	DeleteTodoByID(context.Context, string) error
	DeleteTodosByID(context.Context, []string) error
}

type MongoTodoStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoTodoStore(client *mongo.Client, dbname string, collname string) *MongoTodoStore {
	return &MongoTodoStore{
		client: client,
		coll:   client.Database(dbname).Collection(collname),
	}
}

func (mt *MongoTodoStore) GetTodoByID(ctx context.Context, id string) (*types.Todo, error) {
	return nil, nil
}
func (mt *MongoTodoStore) GetTodoList(context.Context, string) ([]*types.Todo, error) {
	return nil, nil
}
func (mt *MongoTodoStore) CreateTodo(context.Context, *types.Todo) (*types.Todo, error) {
	return nil, nil
}
func (mt *MongoTodoStore) UpdateTodoByID(context.Context, string) (*types.Todo, error) {
	return nil, nil
}
func (mt *MongoTodoStore) DeleteTodoByID(context.Context, string) error {
	return nil
}
func (my *MongoTodoStore) DeleteTodosByID(context.Context, []string) error {
	return nil
}
