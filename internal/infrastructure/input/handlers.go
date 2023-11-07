package input

import (
	"encoding/json"
	"hexagonal/internal/business"
	"net/http"
)

func NewUserHandler(creator business.UserCreator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := User{}

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			// TODO: error
			return
		}

		creator.CreateUser(r.Context(), (business.User)(user))
	}
}
