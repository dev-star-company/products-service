package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PriceType holds the schema definition for the PriceType entity.
type PriceType struct {
	ent.Schema
}

// Fields of the PriceType.
func (PriceType) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Nillable(),
	}
}

func (PriceType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Edges of the PriceType.
func (PriceType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("product_prices", ProductPrices.Type),
	}
}
