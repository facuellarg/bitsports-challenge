package contexthandler

import (
	"bitsports/ent"
	"context"
	"errors"
)

type Key int

const (
	AuthorizationHeaderKey     = "Authorization"
	AuthorizationContexKey Key = iota
	UserContextKey
)

var (
	NoValidationPaths = []string{
		"CreateUser",
		"Login",
	}

	ErrNoUserInContext = errors.New("cant get user from ctx")
)

func SetJwToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, AuthorizationContexKey, token)
}

func SetUserInContext(ctx context.Context, user *ent.User) context.Context {
	return context.WithValue(ctx, UserContextKey, user)
}

func GetUserFromContext(ctx context.Context) (*ent.User, error) {
	user, ok := ctx.Value(UserContextKey).(*ent.User)
	if !ok {
		return nil, ErrNoUserInContext
	}
	return user, nil
}
