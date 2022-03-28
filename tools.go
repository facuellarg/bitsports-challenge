//go:build tools
// +build tools

// Package tools records tool dependencies. It cannot actually be compiled.
package tools

import (
	_ "entgo.io/contrib/entgql"
	_ "entgo.io/ent/cmd/ent"
	_ "github.com/99designs/gqlgen"
)
