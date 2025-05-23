package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Promotions holds the schema definition for the Promotions entity.
type Promotions struct {
	ent.Schema
}

// Fields of the Promotions.
func (Promotions) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Nillable(),
		field.Time("starting_datetime").Nillable(),
		field.Time("ending_datetime").Nillable(),
	}
}

func (Promotions) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Edges of the Promotions.
func (Promotions) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("promotion_has_product", PromotionHasProduct.Type),
	}
}
