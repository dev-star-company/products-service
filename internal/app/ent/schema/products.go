package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Products holds the schema definition for the Products entity.
type Products struct {
	ent.Schema
}

// Fields of the Products.
func (Products) Fields() []ent.Field {
	return []ent.Field{
		field.Int("category_id").Nillable().Optional(),
		field.Int("brand_id").Nillable().Optional(),
		field.Int("variant_type_id").Nillable().Optional(),
		field.Int("product_references_id").Nillable().Optional(),
		field.Int("image_id").Nillable(),
		field.String("name").Nillable(),
		field.Int("stock"),
	}
}

func (Products) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Edges of the Products.
func (Products) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("category", Category.Type).
			Ref("products").
			Unique(),

		edge.From("brand", Brand.Type).
			Ref("products").
			Unique(),

		edge.From("variant_type", VariantType.Type).
			Ref("products").
			Unique(),

		edge.To("product_references", ProductReferences.Type),
		edge.To("images", Images.Type),
		edge.To("product_has_image", ProductHasImage.Type),
		edge.To("promotion_has_product", PromotionHasProduct.Type),
		edge.To("tool_has_product", ToolHasProduct.Type),
		edge.To("product_has_feature", ProductHasFeature.Type),
		edge.To("product_has_info", ProductHasInfo.Type),
		edge.To("product_has_product_reference", ProductHasProductReference.Type),
		edge.To("product_prices", ProductPrices.Type),
	}
}
