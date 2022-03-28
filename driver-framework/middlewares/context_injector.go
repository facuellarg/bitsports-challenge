package middlewares

import (
	"context"
	"net/http"
)

type Key int

const (
	ResquestContextKey Key = iota
	ResponseContextKey
)

func InjectToContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), ResponseContextKey, r)
		ctx = context.WithValue(ctx, ResponseContextKey, &w)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
