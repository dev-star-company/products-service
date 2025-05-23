package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ImageFolderSource holds the schema definition for the ImageFolderSource entity.
type ImageFolderSource struct {
	ent.Schema
}

// Fields of the ImageFolderSource.
func (ImageFolderSource) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Nillable(),
		field.String("base_url").Nillable(),
		field.String("access_key").Nillable().Optional(),
		field.String("secret_key").Nillable().Optional(),
	}
}

func (ImageFolderSource) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Edges of the ImageFolderSource.
func (ImageFolderSource) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("image_folder_path", ImageFolderPath.Type),
	}
}
