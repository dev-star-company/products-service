package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// ProductHasProductReference holds the schema definition for the ProductHasProductReference entity.
type ProductHasProductReference struct {
	ent.Schema
}

// Fields of the ProductHasProductReference.
func (ProductHasProductReference) Fields() []ent.Field {
	return []ent.Field{
		field.Int("product_reference_id").Nillable(),
		field.Int("product_id").Nillable().Optional(),
	}
}

func (ProductHasProductReference) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Edges of the ProductHasProductReference.
func (ProductHasProductReference) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("product_reference", ProductReferences.Type).
			Ref("product_has_product_reference").
			Field("product_reference_id").
			Unique().
			Required(),

		edge.From("product", Products.Type).
			Ref("product_has_product_reference").
			Field("product_id").
			Unique(),
	}
}

func (ProductHasProductReference) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("product_id", "product_reference_id").Unique(),
	}
}
