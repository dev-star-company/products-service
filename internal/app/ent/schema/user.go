package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("Surname").NotEmpty(),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("products", Products.Type),
		edge.To("created_products", Products.Type),
		edge.To("updated_products", Products.Type),
		edge.To("deleted_products", Products.Type),
	}
}
