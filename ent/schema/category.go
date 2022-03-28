package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Category holds the schema definition for the Category entity.
type Category struct {
	ent.Schema
}

// Fields of the Category.
func (Category) Fields() []ent.Field {
	return []ent.Field{
		// field.Enum("categoryName").NamedValues(
		// 	"ElectronicDevice", "ELECTRONIC_DEVICE",
		// 	"ElectronicAccesory", "ELECTRONIC_ACCESORY",
		// 	"HealAndBeauty", "HEALT_AND_BEAUTY",
		// 	"Grocery", "GROCERY",
		// 	"Sports", "SPORTS",
		// 	"Automotive", "AUTOMOTIVE",
		// ),
		field.String("name").NotEmpty().Unique(),
	}
}

// Edges of the Category.
func (Category) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("products", Product.Type).Ref("categories"),
	}
}
