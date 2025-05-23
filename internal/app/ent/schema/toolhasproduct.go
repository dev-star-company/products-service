package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ToolHasProduct holds the schema definition for the ToolHasProduct entity.
type ToolHasProduct struct {
	ent.Schema
}

// Fields of the ToolHasProduct.
func (ToolHasProduct) Fields() []ent.Field {
	return []ent.Field{
		field.Int("tool_id").Nillable(),
		field.Int("product_id").Nillable().Optional(),
	}
}

func (ToolHasProduct) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Edges of the ToolHasProduct.
func (ToolHasProduct) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("product", Products.Type).
			Ref("tool_has_product").
			Unique().
			Required(),

		edge.From("tool", Tools.Type).
			Ref("tool_has_product").
			Unique().
			Required(),
	}
}
