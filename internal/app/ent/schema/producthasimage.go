package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ProductHasImage holds the schema definition for the ProductHasImage entity.
type ProductHasImage struct {
	ent.Schema
}

// Fields of the ProductHasImage.
func (ProductHasImage) Fields() []ent.Field {
	return []ent.Field{
		field.Int("image_id").Nillable(),
		field.Int("product_id").Nillable().Optional(),
		field.Int("priority"),
	}
}

func (ProductHasImage) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Edges of the ProductHasImage.
func (ProductHasImage) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("product", Products.Type).
			Ref("product_has_image").
			Field("product_id").
			Unique(),

		edge.From("image", Images.Type).
			Ref("product_has_image").
			Field("image_id").
			Unique().
			Required(),
	}
}
