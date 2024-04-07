package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// GenerationOutput holds the schema definition for the GenerationOutput entity.
type GenerationOutput struct {
	ent.Schema
}

// Fields of the GenerationOutput.
func (GenerationOutput) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Text("image_path"),
		field.Text("upscaled_image_path").Optional().Nillable(),
		field.Enum("gallery_status").Values("not_submitted", "submitted", "approved", "rejected", "waiting_for_approval").Default("not_submitted"),
		field.Bool("is_favorited").Default(false),
		field.Bool("has_embeddings").Default(false),
		field.Bool("has_embeddings_new").Default(false),
		field.Bool("is_public").Default(false),
		field.Float32("aesthetic_rating_score").Default(0),
		field.Float32("aesthetic_artifact_score").Default(0),
		// Populated by the triggers based on generation_output_likes.
		field.Int("like_count").Default(0),
		// ! Relationships / many-to-one
		field.UUID("generation_id", uuid.UUID{}),
		// ! End relationships
		field.Time("deleted_at").Optional().Nillable(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the GenerationOutput.
func (GenerationOutput) Edges() []ent.Edge {
	return []ent.Edge{
		// M2O with generations
		edge.From("generations", Generation.Type).
			Ref("generation_outputs").
			Field("generation_id").
			Required().
			Unique(),
		// O2O with upscale_outputs
		edge.To("upscale_outputs", UpscaleOutput.Type).Unique(),
		// O2M with generation_output_likes
		edge.To("generation_output_likes", GenerationOutputLike.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
	}
}

// Annotations of the GenerationOutput.
func (GenerationOutput) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "generation_outputs"},
	}
}

// Indexes of the GenerationOutput.
func (GenerationOutput) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id", "gallery_status"),
		index.Fields("gallery_status"),
		index.Fields("created_at"),
		index.Fields("updated_at"),
		index.Fields("generation_id"),
		index.Fields("deleted_at", "is_public"),
	}
}
