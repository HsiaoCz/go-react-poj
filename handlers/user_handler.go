package handlers

import (
	"net/http"

	"github.com/HsiaoCz/go-react-poj/storage"
	"github.com/HsiaoCz/go-react-poj/types"
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
	var req types.CreateUserReq
	if err := c.BodyParser(&req); err != nil {
		return err
	}
	user, err := u.store.User.CreateUser(c.Context(), &types.User{Username: req.Username, Password: req.Password, Email: req.Email})
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "create user success!",
		"user":    user,
	})
}
