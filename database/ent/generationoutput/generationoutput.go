// Code generated by ent, DO NOT EDIT.

package generationoutput

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the generationoutput type in the database.
	Label = "generation_output"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldImagePath holds the string denoting the image_path field in the database.
	FieldImagePath = "image_path"
	// FieldUpscaledImagePath holds the string denoting the upscaled_image_path field in the database.
	FieldUpscaledImagePath = "upscaled_image_path"
	// FieldGalleryStatus holds the string denoting the gallery_status field in the database.
	FieldGalleryStatus = "gallery_status"
	// FieldIsFavorited holds the string denoting the is_favorited field in the database.
	FieldIsFavorited = "is_favorited"
	// FieldHasEmbeddings holds the string denoting the has_embeddings field in the database.
	FieldHasEmbeddings = "has_embeddings"
	// FieldIsPublic holds the string denoting the is_public field in the database.
	FieldIsPublic = "is_public"
	// FieldLikeCount holds the string denoting the like_count field in the database.
	FieldLikeCount = "like_count"
	// FieldGenerationID holds the string denoting the generation_id field in the database.
	FieldGenerationID = "generation_id"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeGenerations holds the string denoting the generations edge name in mutations.
	EdgeGenerations = "generations"
	// EdgeUpscaleOutputs holds the string denoting the upscale_outputs edge name in mutations.
	EdgeUpscaleOutputs = "upscale_outputs"
	// EdgeGenerationOutputLikes holds the string denoting the generation_output_likes edge name in mutations.
	EdgeGenerationOutputLikes = "generation_output_likes"
	// Table holds the table name of the generationoutput in the database.
	Table = "generation_outputs"
	// GenerationsTable is the table that holds the generations relation/edge.
	GenerationsTable = "generation_outputs"
	// GenerationsInverseTable is the table name for the Generation entity.
	// It exists in this package in order to avoid circular dependency with the "generation" package.
	GenerationsInverseTable = "generations"
	// GenerationsColumn is the table column denoting the generations relation/edge.
	GenerationsColumn = "generation_id"
	// UpscaleOutputsTable is the table that holds the upscale_outputs relation/edge.
	UpscaleOutputsTable = "upscale_outputs"
	// UpscaleOutputsInverseTable is the table name for the UpscaleOutput entity.
	// It exists in this package in order to avoid circular dependency with the "upscaleoutput" package.
	UpscaleOutputsInverseTable = "upscale_outputs"
	// UpscaleOutputsColumn is the table column denoting the upscale_outputs relation/edge.
	UpscaleOutputsColumn = "generation_output_id"
	// GenerationOutputLikesTable is the table that holds the generation_output_likes relation/edge.
	GenerationOutputLikesTable = "generation_output_likes"
	// GenerationOutputLikesInverseTable is the table name for the GenerationOutputLike entity.
	// It exists in this package in order to avoid circular dependency with the "generationoutputlike" package.
	GenerationOutputLikesInverseTable = "generation_output_likes"
	// GenerationOutputLikesColumn is the table column denoting the generation_output_likes relation/edge.
	GenerationOutputLikesColumn = "output_id"
)

// Columns holds all SQL columns for generationoutput fields.
var Columns = []string{
	FieldID,
	FieldImagePath,
	FieldUpscaledImagePath,
	FieldGalleryStatus,
	FieldIsFavorited,
	FieldHasEmbeddings,
	FieldIsPublic,
	FieldLikeCount,
	FieldGenerationID,
	FieldDeletedAt,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultIsFavorited holds the default value on creation for the "is_favorited" field.
	DefaultIsFavorited bool
	// DefaultHasEmbeddings holds the default value on creation for the "has_embeddings" field.
	DefaultHasEmbeddings bool
	// DefaultIsPublic holds the default value on creation for the "is_public" field.
	DefaultIsPublic bool
	// DefaultLikeCount holds the default value on creation for the "like_count" field.
	DefaultLikeCount int
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// GalleryStatus defines the type for the "gallery_status" enum field.
type GalleryStatus string

// GalleryStatusNotSubmitted is the default value of the GalleryStatus enum.
const DefaultGalleryStatus = GalleryStatusNotSubmitted

// GalleryStatus values.
const (
	GalleryStatusNotSubmitted GalleryStatus = "not_submitted"
	GalleryStatusSubmitted    GalleryStatus = "submitted"
	GalleryStatusApproved     GalleryStatus = "approved"
	GalleryStatusRejected     GalleryStatus = "rejected"
)

func (gs GalleryStatus) String() string {
	return string(gs)
}

// GalleryStatusValidator is a validator for the "gallery_status" field enum values. It is called by the builders before save.
func GalleryStatusValidator(gs GalleryStatus) error {
	switch gs {
	case GalleryStatusNotSubmitted, GalleryStatusSubmitted, GalleryStatusApproved, GalleryStatusRejected:
		return nil
	default:
		return fmt.Errorf("generationoutput: invalid enum value for gallery_status field: %q", gs)
	}
}
