package main

import (
	"hexagonal/internal/business"
	"hexagonal/internal/infrastructure/input"
	"hexagonal/internal/infrastructure/output"
	"net/http"
)

func main() {
	saver := output.NewUserSaver()            // Secondary adapter
	creator := business.NewUserCreator(saver) // Business logic (Port for primary adapter)
	handler := input.NewUserHandler(creator)  // Primary adapter

	http.Handle("/v1/users/", handler)

	http.ListenAndServe(":8080", nil)
}
