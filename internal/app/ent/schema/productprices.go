package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ProductPrices holds the schema definition for the ProductPrices entity.
type ProductPrices struct {
	ent.Schema
}

// Fields of the ProductPrices.
func (ProductPrices) Fields() []ent.Field {
	return []ent.Field{
		field.Int("price_type_id").Nillable(),
		field.Int("product_id").Nillable().Optional(),
		field.Float("default_value").Nillable(),
		field.Float("min_value").Nillable().Optional(),
	}
}

func (ProductPrices) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Edges of the ProductPrices.
func (ProductPrices) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("product", Products.Type).
			Ref("product_prices").
			Field("product_id").
			Unique(),

		edge.From("price_type", PriceType.Type).
			Ref("product_prices").
			Field("price_type_id").
			Unique().
			Required(),
	}
}
