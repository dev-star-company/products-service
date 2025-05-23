package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Images holds the schema definition for the Images entity.
type Images struct {
	ent.Schema
}

// Fields of the Images.
func (Images) Fields() []ent.Field {
	return []ent.Field{
		field.Int("image_folder_path_id").Nillable(),
		field.String("content").Nillable(),
		field.String("path").Nillable(),
	}
}

func (Images) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Edges of the Images.
func (Images) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("image_folder_path", ImageFolderPath.Type).
			Ref("images").
			Field("image_folder_path_id").
			Unique().
			Required(),
		edge.To("product_has_image", ProductHasImage.Type),
	}
}
