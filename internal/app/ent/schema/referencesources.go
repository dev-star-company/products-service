package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ReferenceSources holds the schema definition for the ReferenceSources entity.
type ReferenceSources struct {
	ent.Schema
}

// Fields of the ReferenceSources.
func (ReferenceSources) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Nillable(),
	}
}

func (ReferenceSources) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Edges of the ReferenceSources.
func (ReferenceSources) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("product_references", ProductReferences.Type),
	}
}
