package handlers

import (
	"net/http"

	"github.com/HsiaoCz/go-react-poj/storage"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	store *storage.Store
}

func NewUserHandler(store *storage.Store) *UserHandler {
	return &UserHandler{
		store: store,
	}
}

func (u *UserHandler) HandleUserSignup(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "good luck",
	})
}
