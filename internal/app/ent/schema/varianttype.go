package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// VariantType holds the schema definition for the VariantType entity.
type VariantType struct {
	ent.Schema
}

// Fields of the VariantType.
func (VariantType) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Nillable(),
	}
}

func (VariantType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Edges of the VariantType.
func (VariantType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("products", Products.Type),
	}
}
