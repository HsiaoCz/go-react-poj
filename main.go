package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/HsiaoCz/go-react-poj/handlers"
	"github.com/HsiaoCz/go-react-poj/settings"
	"github.com/HsiaoCz/go-react-poj/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	file, err := os.OpenFile("./log.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(settings.GetMongoUri()))
	if err != nil {
		log.Fatal(err)
	}

	logger := slog.New(slog.NewJSONHandler(file, &slog.HandlerOptions{}))
	slog.SetDefault(logger)

	var (
		mongoUserStore = storage.NewMongoUserStore(client, settings.GetDBname(), settings.GetUserColl())
		mongoTodoStore = storage.NewMongoTodoStore(client, settings.GetDBname(), settings.GetTodoColl())
		store          = &storage.Store{User: mongoUserStore, Todo: mongoTodoStore}
		userHandler    = handlers.NewUserHandler(store)
		todoHandler    = handlers.NewTodoHandler(store)
		port           = settings.GetPort()
	)

	app := fiber.New()

	app.Post("/user/signup", userHandler.HandleUserSignup)
	app.Get("/todo/{id}", todoHandler.HandleGetTodoByID)

	go func() {
		slog.Info("the server is running", "port", port)
		if err := app.Listen(port); err != nil {
			log.Fatal(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c

	app.Shutdown()

	slog.Info("the server is shuting down")

}
