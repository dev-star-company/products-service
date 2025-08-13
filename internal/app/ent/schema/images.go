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
			Unique().
			Required(),
		edge.To("product_has_image", ProductHasImage.Type),
	}
}
