package jwt

import (
	"bitsports/ent"
	contexthandler "bitsports/usecase/context_handler"
	"bitsports/utils"
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	ErrNoTokenInContext = errors.New("no token in context")
	ErrCantAssert       = errors.New("cant assert token to claims")
)

type UserClaims struct {
	ID int `json:"id"`
	jwt.StandardClaims
}

func GetSecret() string {
	return utils.GetEnvOrDefault("JWT_SECRET", "TOPSECRET")
}

func ValidateToken(ctx context.Context, client *ent.Client) (*ent.User, error) {
	tokenString, ok := ctx.Value(contexthandler.AuthorizationContexKey).(string)
	if !ok {
		return nil, ErrNoTokenInContext
	}
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrCantAssert
		}
		return []byte(GetSecret()), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok || !token.Valid {
		return nil, ErrCantAssert
	}
	user, err := client.User.Get(ctx, claims.ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func CreateToken(user UserClaims) (string, error) {

	claims := jwt.MapClaims{}
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(2 * time.Hour).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return at.SignedString([]byte(GetSecret()))
}

func ValidateOperation(ctx context.Context, client *ent.Client, operation string) (context.Context, error) {

	if !utils.Contains(contexthandler.NoValidationPaths, operation) {
		user, err := ValidateToken(ctx, client)
		if err != nil {
			return ctx, err
		}
		ctx = context.WithValue(ctx, contexthandler.UserContextKey, user)
	}
	return ctx, nil

}

func SetJwtTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jwt := r.Header.Get(fmt.Sprint(contexthandler.AuthorizationHeaderKey))
		ctx := contexthandler.SetJwToken(r.Context(), jwt)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
