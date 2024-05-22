package handlers

import (
	"github.com/HsiaoCz/go-react-poj/storage"
	"github.com/gofiber/fiber/v2"
)

type TodoHandler struct {
	store *storage.Store
}

func NewTodoHandler(store *storage.Store) *TodoHandler {
	return &TodoHandler{
		store: store,
	}
}

func (t *TodoHandler) HandleGetTodos(c *fiber.Ctx) error {
	return nil
}

func (t *TodoHandler) HandleGetTodoByID(c *fiber.Ctx) error {
	return nil
}

func (t *TodoHandler) HandleCreateTodo(c *fiber.Ctx) error {
	return nil
}

func (t *TodoHandler) HandleUpdateTodoByID(c *fiber.Ctx) error {
	return nil
}

func (t *TodoHandler) HandleDeleteTodoByID(c *fiber.Ctx) error {
	return nil
}

func (t *TodoHandler) HandleDeleteTodos(c *fiber.Ctx) error {
	return nil
}

func (t *TodoHandler) HandleClearTodoList(c *fiber.Ctx) error {
	return nil
}

func (t *TodoHandler) HandleDiscoveryTodoList(c *fiber.Ctx) error {
	return nil
}
