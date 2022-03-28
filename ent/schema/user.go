package schema

import (
	"regexp"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

const (
	EmailPattern = `^(?i)([A-Za-z0-9!#$%&'*+\/=?^_{|.}~-]+@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?)$`
)

var (
	// passwordRegex, _ = regexp.Compile("^(?=.*[A-Za-z])(?=.*\\d)[A-Za-z\\d]{8,}$")
	EmailRegex = regexp.MustCompile(EmailPattern)
)

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		// field.String("email").Unique().NotEmpty(),
		field.String("email").NotEmpty().Unique(),
		field.String("password").NotEmpty(),
		field.String("name").NotEmpty(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("products", Product.Type).Ref("users"),
	}
}
