package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// FeaturesValues holds the schema definition for the FeaturesValues entity.
type FeaturesValues struct {
	ent.Schema
}

// Fields of the FeaturesValues.
func (FeaturesValues) Fields() []ent.Field {
	return []ent.Field{
		field.Int("feature_id").Nillable(),
		field.Int("feature_unit_values_id").Nillable().Optional(),
		field.Int("feature_values_id").Nillable().Optional(),
		field.String("value").Nillable(),
	}
}

func (FeaturesValues) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Edges of the FeaturesValues.
func (FeaturesValues) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("feature", Features.Type).
			Ref("feature_values").
			Field("feature_id").
			Unique().
			Required(),

		edge.From("feature_unit_values", FeaturesUnitValues.Type).
			Ref("feature_values").
			Field("feature_unit_values_id").
			Unique(),
	}
}
