package handlers

import "github.com/gofiber/fiber/v2"

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (u *UserHandler) HandleUserSignup(c *fiber.Ctx) error{
	return nil
}
