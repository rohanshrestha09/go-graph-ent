package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/rohanshrestha09/go-graph-ent/common/enums"
)

// Blog holds the schema definition for the Blog entity.
type Blog struct {
	ent.Schema
}

// Fields of the Blog.
func (Blog) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("title"),
		field.String("slug").
			Unique(),
		field.String("content").
			SchemaType(map[string]string{
				dialect.MySQL: "longtext",
			}),
		field.String("image").
			Optional(),
		field.Enum("status").
			GoType(enums.BlogStatus("")).
			Default(string(enums.Published)),
		field.UUID("user_id", uuid.UUID{}).
			Immutable(),
	}
}

// Edges of the Blog.
func (Blog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Field("user_id").
			Unique().
			Required().
			Immutable(),
	}
}

func (Blog) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
