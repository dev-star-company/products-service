package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// FeaturesUnitValues holds the schema definition for the FeaturesUnitValues entity.
type FeaturesUnitValues struct {
	ent.Schema
}

// Fields of the FeaturesUnitValues.
func (FeaturesUnitValues) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Nillable(),
		field.Int("decimals").Nillable().Optional(),
	}
}

func (FeaturesUnitValues) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Edges of the FeaturesUnitValues.
func (FeaturesUnitValues) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("feature_values", FeaturesValues.Type),
	}
}
