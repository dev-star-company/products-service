package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// InfoTypes holds the schema definition for the InfoTypes entity.
type InfoTypes struct {
	ent.Schema
}

// Fields of the InfoTypes.
func (InfoTypes) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Nillable(),
	}
}

func (InfoTypes) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Edges of the InfoTypes.
func (InfoTypes) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("product_info", ProductInfo.Type),
	}
}
