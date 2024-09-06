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

// GenerationModel holds the schema definition for the GenerationModel entity.
type GenerationModel struct {
	ent.Schema
}

// Fields of the GenerationModel.
func (GenerationModel) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Text("name_in_worker"),
		field.Text("short_name"),
		field.Bool("is_active").Default(true),
		field.Bool("is_default").Default(false),
		field.Bool("is_hidden").Default(false),
		field.String("runpod_endpoint").Optional().Nillable(),
		field.Bool("runpod_active").Default(false),
		field.Int32("display_weight").Default(0),
		field.UUID("default_scheduler_id", uuid.UUID{}).Optional().Nillable(),
		field.Int32("default_width").Default(512),
		field.Int32("default_height").Default(512),
		field.Int32("default_inference_steps").Default(25),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Model.
func (GenerationModel) Edges() []ent.Edge {
	return []ent.Edge{
		// O2M with generation
		edge.To("generations", Generation.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
		// M2M schedulers
		edge.To("schedulers", Scheduler.Type).StorageKey(edge.Table("generation_model_compatible_schedulers")),
	}
}

// Annotations of the GenerationModel.
func (GenerationModel) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "generation_models"},
	}
}
