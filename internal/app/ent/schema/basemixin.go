package schema

import (
	"context"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// BaseMixin holds the schema definition for the BaseMixin entity.
type BaseMixin struct {
	mixin.Schema
}

// Fields of the BaseMixin.
func (BaseMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Optional().Nillable(),
		field.Int("created_by").Positive().Immutable(),
		field.Int("updated_by").Positive(),
		field.Int("deleted_by").Optional().Nillable(),
	}
}

func (BaseMixin) Edges() []ent.Edge {
	return nil
}

func (m BaseMixin) EdgesMixin(schemaName string) []ent.Edge {
	switch schemaName {
	case "Products":
		return []ent.Edge{
			edge.From("created_by_user", User.Type).
				Ref("created_products").
				Field("created_by").
				Unique(),

			edge.From("updated_by_user", User.Type).
				Ref("updated_products").
				Field("updated_by").
				Unique(),

			edge.From("deleted_by_user", User.Type).
				Ref("deleted_products").
				Field("deleted_by").
				Unique(),
		}
	default:
		return nil
	}
}

type softDeleteKey struct{}

func SkipSoftDelete(parent context.Context) context.Context {
	return context.WithValue(parent, softDeleteKey{}, true)
}

func (d BaseMixin) P(w interface{ WhereP(...func(*sql.Selector)) }) {
	w.WhereP(
		sql.FieldIsNull(d.Fields()[2].Descriptor().Name),
	)
}
