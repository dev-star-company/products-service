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
		field.Int("images_id").Nillable().Optional(),
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
			Field("category_id").
			Unique(),

		edge.From("brand", Brand.Type).
			Ref("products").
			Field("brand_id").
			Unique(),

		edge.From("variant_type", VariantType.Type).
			Ref("products").
			Field("variant_type_id").
			Unique(),

		edge.From("product_references", ProductReferences.Type).
			Ref("products").
			Field("product_references_id").
			Unique(),

		edge.From("images", Images.Type).
			Ref("products").
			Field("images_id").
			Unique(),
		edge.To("product_has_image", ProductHasImage.Type),
		edge.To("promotion_has_product", PromotionHasProduct.Type),
		edge.To("tool_has_product", ToolHasProduct.Type),
		edge.To("product_has_feature", ProductHasFeature.Type),
		edge.To("product_has_info", ProductHasInfo.Type),
		edge.To("product_has_product_reference", ProductHasProductReference.Type),
		edge.To("product_prices", ProductPrices.Type),
	}
}
