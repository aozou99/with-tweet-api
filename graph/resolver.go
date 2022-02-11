package graph

//go:generate go run github.com/99designs/gqlgen generate
import (
	"aozou99/with-tweet-api/pkg/repository"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Repository repository.Repository
}
