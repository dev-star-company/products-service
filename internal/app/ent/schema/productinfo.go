package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ProductInfo holds the schema definition for the ProductInfo entity.
type ProductInfo struct {
	ent.Schema
}

// Fields of the ProductInfo.
func (ProductInfo) Fields() []ent.Field {
	return []ent.Field{
		field.Int("info_types_id").Nillable().Optional(),
		field.String("value").Nillable(),
	}
}

func (ProductInfo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Edges of the ProductInfo.
func (ProductInfo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("info_type", InfoTypes.Type).
			Ref("product_info").
			Field("info_types_id").
			Unique(),
		edge.To("products", ProductHasInfo.Type),
		edge.To("product_has_info", ProductHasInfo.Type),
	}
}
