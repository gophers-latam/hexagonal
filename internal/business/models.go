package business

import "errors"

type User struct {
	Name string
}

func (u User) Validate() error {
	switch {
	case u.Name == "":
		return errors.New("invalid name")
	}

	return nil
}
