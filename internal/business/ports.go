package business

import "context"

// Ports for primary adapters
type (
	UserCreator interface {
		CreateUser(context.Context, User) error
	}
)

// Ports for secondary adapters
type (
	UserSaver interface {
		SaveUser(context.Context, User) error
	}
)
