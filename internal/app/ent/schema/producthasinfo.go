package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// ProductHasInfo holds the schema definition for the ProductHasInfo entity.
type ProductHasInfo struct {
	ent.Schema
}

// Fields of the ProductHasInfo.
func (ProductHasInfo) Fields() []ent.Field {
	return []ent.Field{
		field.Int("product_id").Nillable().Optional(),
		field.Int("product_info_id").Nillable().Optional(),
	}
}

func (ProductHasInfo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Edges of the ProductHasInfo.
func (ProductHasInfo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("products", Products.Type).
			Ref("product_has_info").
			Field("product_id").
			Unique(),

		edge.From("product_info", ProductInfo.Type).
			Ref("product_has_info").
			Field("product_info_id").
			Unique(),
	}
}

func (ProductHasInfo) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("product_id", "product_info_id").Unique(),
	}
}
