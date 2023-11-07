package output

import (
	"context"
	"hexagonal/internal/business"
	"log"
)

func NewUserSaver() business.UserSaver {
	return userSaver{}
}

type userSaver struct{}

func (u userSaver) SaveUser(ctx context.Context, user business.User) error {
	// TODO: save user
	log.Println((User)(user))
	return nil
}
