package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// NegativePrompt holds the schema definition for the NegativePrompt entity.
type NegativePrompt struct {
	ent.Schema
}

// Fields of the NegativePrompt.
func (NegativePrompt) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Text("text"),
		field.Text("translated_text").Optional().Nillable(),
		field.Bool("ran_translation").Default(false),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the NegativePrompt.
func (NegativePrompt) Edges() []ent.Edge {
	return []ent.Edge{
		// O2M with generations
		edge.To("generations", Generation.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
	}
}

// Annotations of the NegativePrompt.
func (NegativePrompt) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "negative_prompts"},
	}
}
