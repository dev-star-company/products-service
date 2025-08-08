package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// ProductHasFeature holds the schema definition for the ProductHasFeature entity.
type ProductHasFeature struct {
	ent.Schema
}

// Fields of the ProductHasFeature.
func (ProductHasFeature) Fields() []ent.Field {
	return []ent.Field{
		field.Int("feature_id").Nillable(),
		field.Int("product_id").Nillable().Optional(),
	}
}

func (ProductHasFeature) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Edges of the ProductHasFeature.
func (ProductHasFeature) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("products", Products.Type).
			Ref("product_has_feature").
			Field("product_id").
			Unique(),
		edge.From("features", Features.Type).
			Ref("product_has_feature").
			Field("feature_id").
			Unique().
			Required(),
	}
}

func (ProductHasFeature) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("product_id", "feature_id").Unique(),
	}
}
