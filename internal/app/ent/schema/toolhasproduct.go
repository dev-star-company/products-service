package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// ToolHasProduct holds the schema definition for the ToolHasProduct entity.
type ToolHasProduct struct {
	ent.Schema
}

// Fields of the ToolHasProduct.
func (ToolHasProduct) Fields() []ent.Field {
	return []ent.Field{
		field.Int("products_id").Nillable().Optional(),
		field.Int("tools_id").Nillable().Optional(),
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
		edge.From("products", Products.Type).
			Ref("tool_has_product").
			Field("products_id").
			Unique(),

		edge.From("tools", Tools.Type).
			Ref("tool_has_product").
			Field("tools_id").
			Unique(),
	}
}

func (ToolHasProduct) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("products_id", "tools_id").Unique(),
	}
}
