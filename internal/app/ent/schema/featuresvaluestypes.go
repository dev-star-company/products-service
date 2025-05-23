package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// FeaturesValuesTypes holds the schema definition for the FeaturesValuesTypes entity.
type FeaturesValuesTypes struct {
	ent.Schema
}

// Fields of the FeaturesValuesTypes.
func (FeaturesValuesTypes) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Nillable(),
	}
}

func (FeaturesValuesTypes) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Edges of the FeaturesValuesTypes.
func (FeaturesValuesTypes) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("product_info", ProductInfo.Type),
		edge.To("feature_values", FeaturesValues.Type),
	}
}
