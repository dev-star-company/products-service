package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Tools holds the schema definition for the Tools entity.
type Tools struct {
	ent.Schema
}

// Fields of the Tools.
func (Tools) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Nillable(),
	}
}

func (Tools) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Edges of the Tools.
func (Tools) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tool_has_product", ToolHasProduct.Type),
	}
}
