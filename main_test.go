package main

import (
	"context"
	"testing"

	"github.com/HsiaoCz/go-react-poj/settings"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestMongoConnect(t *testing.T) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(settings.GetMongoUri()))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(client)
}
