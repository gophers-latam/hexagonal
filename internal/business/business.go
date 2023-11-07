package business

import "context"

func NewUserCreator(saver UserSaver) UserCreator {
	return userCreator{
		saver: saver,
	}
}

type userCreator struct {
	saver UserSaver
}

func (u userCreator) CreateUser(ctx context.Context, user User) error {
	if err := user.Validate(); err != nil {
		return err
	}

	return u.saver.SaveUser(ctx, user)
}
