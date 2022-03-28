package bitsports

import (
	"bitsports/ent"
	passwordvalidator "bitsports/usecase/password_validator"

	"github.com/99designs/gqlgen/graphql"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	client            *ent.Client
	passwordValidator passwordvalidator.PasswordValidatorI
}

func NewSchema(client *ent.Client, pv passwordvalidator.PasswordValidatorI) graphql.ExecutableSchema {
	return NewExecutableSchema(Config{
		Resolvers: &Resolver{client, pv},
	})
}
