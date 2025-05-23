package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ProductReferences holds the schema definition for the ProductReferences entity.
type ProductReferences struct {
	ent.Schema
}

// Fields of the ProductReferences.
func (ProductReferences) Fields() []ent.Field {
	return []ent.Field{
		field.Int("reference_source_id").Nillable(),
		field.String("value").Nillable(),
	}
}

func (ProductReferences) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Edges of the ProductReferences.
func (ProductReferences) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("product", Products.Type).
			Ref("product_references").
			Unique().
			Required(),

		edge.From("reference_sources", ReferenceSources.Type).
			Ref("product_references").
			Unique().
			Required(),

		edge.To("product_has_product_reference", ProductHasProductReference.Type),
	}
}
