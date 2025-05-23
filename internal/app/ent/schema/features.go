package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Features holds the schema definition for the Features entity.
type Features struct {
	ent.Schema
}

// Fields of the Features.
func (Features) Fields() []ent.Field {
	return []ent.Field{
		field.Int("feature_value_id").Nillable(),
		field.String("name").Nillable(),
	}
}

func (Features) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Edges of the Features.
func (Features) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("feature_values", FeaturesValues.Type),
		edge.To("product_has_feature", ProductHasFeature.Type),
	}
}
