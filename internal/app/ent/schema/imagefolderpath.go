package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ImageFolderPath holds the schema definition for the ImageFolderPath entity.
type ImageFolderPath struct {
	ent.Schema
}

// Fields of the ImageFolderPath.
func (ImageFolderPath) Fields() []ent.Field {
	return []ent.Field{
		field.Int("image_folder_source_id").Nillable(),
	}
}

func (ImageFolderPath) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Edges of the ImageFolderPath.
func (ImageFolderPath) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("image_folder_source", ImageFolderSource.Type).
			Ref("image_folder_path").
			Field("image_folder_source_id").
			Unique().
			Required(),

		edge.To("images", Images.Type),
	}
}
