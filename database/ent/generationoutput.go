// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/stablecog/sc-go/database/ent/generation"
	"github.com/stablecog/sc-go/database/ent/generationoutput"
	"github.com/stablecog/sc-go/database/ent/upscaleoutput"
)

// GenerationOutput is the model entity for the GenerationOutput schema.
type GenerationOutput struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// ImagePath holds the value of the "image_path" field.
	ImagePath string `json:"image_path,omitempty"`
	// UpscaledImagePath holds the value of the "upscaled_image_path" field.
	UpscaledImagePath *string `json:"upscaled_image_path,omitempty"`
	// GalleryStatus holds the value of the "gallery_status" field.
	GalleryStatus generationoutput.GalleryStatus `json:"gallery_status,omitempty"`
	// IsFavorited holds the value of the "is_favorited" field.
	IsFavorited bool `json:"is_favorited,omitempty"`
	// HasEmbeddings holds the value of the "has_embeddings" field.
	HasEmbeddings bool `json:"has_embeddings,omitempty"`
	// IsPublic holds the value of the "is_public" field.
	IsPublic bool `json:"is_public,omitempty"`
	// GenerationID holds the value of the "generation_id" field.
	GenerationID uuid.UUID `json:"generation_id,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GenerationOutputQuery when eager-loading is set.
	Edges GenerationOutputEdges `json:"edges"`
}

// GenerationOutputEdges holds the relations/edges for other nodes in the graph.
type GenerationOutputEdges struct {
	// Generations holds the value of the generations edge.
	Generations *Generation `json:"generations,omitempty"`
	// UpscaleOutputs holds the value of the upscale_outputs edge.
	UpscaleOutputs *UpscaleOutput `json:"upscale_outputs,omitempty"`
	// ZoomedFromGeneration holds the value of the zoomed_from_generation edge.
	ZoomedFromGeneration []*Generation `json:"zoomed_from_generation,omitempty"`
	// ZoomedFromOutput holds the value of the zoomed_from_output edge.
	ZoomedFromOutput []*GenerationOutput `json:"zoomed_from_output,omitempty"`
	// ZoomedOutputs holds the value of the zoomed_outputs edge.
	ZoomedOutputs []*GenerationOutput `json:"zoomed_outputs,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// GenerationsOrErr returns the Generations value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GenerationOutputEdges) GenerationsOrErr() (*Generation, error) {
	if e.loadedTypes[0] {
		if e.Generations == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: generation.Label}
		}
		return e.Generations, nil
	}
	return nil, &NotLoadedError{edge: "generations"}
}

// UpscaleOutputsOrErr returns the UpscaleOutputs value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GenerationOutputEdges) UpscaleOutputsOrErr() (*UpscaleOutput, error) {
	if e.loadedTypes[1] {
		if e.UpscaleOutputs == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: upscaleoutput.Label}
		}
		return e.UpscaleOutputs, nil
	}
	return nil, &NotLoadedError{edge: "upscale_outputs"}
}

// ZoomedFromGenerationOrErr returns the ZoomedFromGeneration value or an error if the edge
// was not loaded in eager-loading.
func (e GenerationOutputEdges) ZoomedFromGenerationOrErr() ([]*Generation, error) {
	if e.loadedTypes[2] {
		return e.ZoomedFromGeneration, nil
	}
	return nil, &NotLoadedError{edge: "zoomed_from_generation"}
}

// ZoomedFromOutputOrErr returns the ZoomedFromOutput value or an error if the edge
// was not loaded in eager-loading.
func (e GenerationOutputEdges) ZoomedFromOutputOrErr() ([]*GenerationOutput, error) {
	if e.loadedTypes[3] {
		return e.ZoomedFromOutput, nil
	}
	return nil, &NotLoadedError{edge: "zoomed_from_output"}
}

// ZoomedOutputsOrErr returns the ZoomedOutputs value or an error if the edge
// was not loaded in eager-loading.
func (e GenerationOutputEdges) ZoomedOutputsOrErr() ([]*GenerationOutput, error) {
	if e.loadedTypes[4] {
		return e.ZoomedOutputs, nil
	}
	return nil, &NotLoadedError{edge: "zoomed_outputs"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*GenerationOutput) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case generationoutput.FieldIsFavorited, generationoutput.FieldHasEmbeddings, generationoutput.FieldIsPublic:
			values[i] = new(sql.NullBool)
		case generationoutput.FieldImagePath, generationoutput.FieldUpscaledImagePath, generationoutput.FieldGalleryStatus:
			values[i] = new(sql.NullString)
		case generationoutput.FieldDeletedAt, generationoutput.FieldCreatedAt, generationoutput.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case generationoutput.FieldID, generationoutput.FieldGenerationID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type GenerationOutput", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the GenerationOutput fields.
func (_go *GenerationOutput) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case generationoutput.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				_go.ID = *value
			}
		case generationoutput.FieldImagePath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field image_path", values[i])
			} else if value.Valid {
				_go.ImagePath = value.String
			}
		case generationoutput.FieldUpscaledImagePath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field upscaled_image_path", values[i])
			} else if value.Valid {
				_go.UpscaledImagePath = new(string)
				*_go.UpscaledImagePath = value.String
			}
		case generationoutput.FieldGalleryStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field gallery_status", values[i])
			} else if value.Valid {
				_go.GalleryStatus = generationoutput.GalleryStatus(value.String)
			}
		case generationoutput.FieldIsFavorited:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_favorited", values[i])
			} else if value.Valid {
				_go.IsFavorited = value.Bool
			}
		case generationoutput.FieldHasEmbeddings:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field has_embeddings", values[i])
			} else if value.Valid {
				_go.HasEmbeddings = value.Bool
			}
		case generationoutput.FieldIsPublic:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_public", values[i])
			} else if value.Valid {
				_go.IsPublic = value.Bool
			}
		case generationoutput.FieldGenerationID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field generation_id", values[i])
			} else if value != nil {
				_go.GenerationID = *value
			}
		case generationoutput.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				_go.DeletedAt = new(time.Time)
				*_go.DeletedAt = value.Time
			}
		case generationoutput.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				_go.CreatedAt = value.Time
			}
		case generationoutput.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				_go.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryGenerations queries the "generations" edge of the GenerationOutput entity.
func (_go *GenerationOutput) QueryGenerations() *GenerationQuery {
	return NewGenerationOutputClient(_go.config).QueryGenerations(_go)
}

// QueryUpscaleOutputs queries the "upscale_outputs" edge of the GenerationOutput entity.
func (_go *GenerationOutput) QueryUpscaleOutputs() *UpscaleOutputQuery {
	return NewGenerationOutputClient(_go.config).QueryUpscaleOutputs(_go)
}

// QueryZoomedFromGeneration queries the "zoomed_from_generation" edge of the GenerationOutput entity.
func (_go *GenerationOutput) QueryZoomedFromGeneration() *GenerationQuery {
	return NewGenerationOutputClient(_go.config).QueryZoomedFromGeneration(_go)
}

// QueryZoomedFromOutput queries the "zoomed_from_output" edge of the GenerationOutput entity.
func (_go *GenerationOutput) QueryZoomedFromOutput() *GenerationOutputQuery {
	return NewGenerationOutputClient(_go.config).QueryZoomedFromOutput(_go)
}

// QueryZoomedOutputs queries the "zoomed_outputs" edge of the GenerationOutput entity.
func (_go *GenerationOutput) QueryZoomedOutputs() *GenerationOutputQuery {
	return NewGenerationOutputClient(_go.config).QueryZoomedOutputs(_go)
}

// Update returns a builder for updating this GenerationOutput.
// Note that you need to call GenerationOutput.Unwrap() before calling this method if this GenerationOutput
// was returned from a transaction, and the transaction was committed or rolled back.
func (_go *GenerationOutput) Update() *GenerationOutputUpdateOne {
	return NewGenerationOutputClient(_go.config).UpdateOne(_go)
}

// Unwrap unwraps the GenerationOutput entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (_go *GenerationOutput) Unwrap() *GenerationOutput {
	_tx, ok := _go.config.driver.(*txDriver)
	if !ok {
		panic("ent: GenerationOutput is not a transactional entity")
	}
	_go.config.driver = _tx.drv
	return _go
}

// String implements the fmt.Stringer.
func (_go *GenerationOutput) String() string {
	var builder strings.Builder
	builder.WriteString("GenerationOutput(")
	builder.WriteString(fmt.Sprintf("id=%v, ", _go.ID))
	builder.WriteString("image_path=")
	builder.WriteString(_go.ImagePath)
	builder.WriteString(", ")
	if v := _go.UpscaledImagePath; v != nil {
		builder.WriteString("upscaled_image_path=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("gallery_status=")
	builder.WriteString(fmt.Sprintf("%v", _go.GalleryStatus))
	builder.WriteString(", ")
	builder.WriteString("is_favorited=")
	builder.WriteString(fmt.Sprintf("%v", _go.IsFavorited))
	builder.WriteString(", ")
	builder.WriteString("has_embeddings=")
	builder.WriteString(fmt.Sprintf("%v", _go.HasEmbeddings))
	builder.WriteString(", ")
	builder.WriteString("is_public=")
	builder.WriteString(fmt.Sprintf("%v", _go.IsPublic))
	builder.WriteString(", ")
	builder.WriteString("generation_id=")
	builder.WriteString(fmt.Sprintf("%v", _go.GenerationID))
	builder.WriteString(", ")
	if v := _go.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(_go.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(_go.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// GenerationOutputs is a parsable slice of GenerationOutput.
type GenerationOutputs []*GenerationOutput

func (_go GenerationOutputs) config(cfg config) {
	for _i := range _go {
		_go[_i].config = cfg
	}
}
