package storage

import "context"

type UserStorer interface {
	CreateUser(context.Context) error
}
