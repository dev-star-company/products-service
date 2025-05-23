package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PromotionHasProduct holds the schema definition for the PromotionHasProduct entity.
type PromotionHasProduct struct {
	ent.Schema
}

// Fields of the PromotionHasProduct.
func (PromotionHasProduct) Fields() []ent.Field {
	return []ent.Field{
		field.Int("product_id").Nillable().Optional(),
		field.Int("promotion_id").Nillable(),
		field.Float("promocional_price").Nillable(),
	}
}

func (PromotionHasProduct) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Edges of the PromotionHasProduct.
func (PromotionHasProduct) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("product", Products.Type).
			Ref("promotion_has_product").
			Unique().
			Required(),

		edge.From("promotion", Promotions.Type).
			Ref("promotion_has_product").
			Unique().
			Required(),
	}
}
