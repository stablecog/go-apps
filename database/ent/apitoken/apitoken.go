// Code generated by ent, DO NOT EDIT.

package apitoken

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the apitoken type in the database.
	Label = "api_token"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldHashedToken holds the string denoting the hashed_token field in the database.
	FieldHashedToken = "hashed_token"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldShortString holds the string denoting the short_string field in the database.
	FieldShortString = "short_string"
	// FieldIsActive holds the string denoting the is_active field in the database.
	FieldIsActive = "is_active"
	// FieldUses holds the string denoting the uses field in the database.
	FieldUses = "uses"
	// FieldCreditsSpent holds the string denoting the credits_spent field in the database.
	FieldCreditsSpent = "credits_spent"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldLastUsedAt holds the string denoting the last_used_at field in the database.
	FieldLastUsedAt = "last_used_at"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeGenerations holds the string denoting the generations edge name in mutations.
	EdgeGenerations = "generations"
	// EdgeUpscales holds the string denoting the upscales edge name in mutations.
	EdgeUpscales = "upscales"
	// Table holds the table name of the apitoken in the database.
	Table = "api_tokens"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "api_tokens"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
	// GenerationsTable is the table that holds the generations relation/edge.
	GenerationsTable = "generations"
	// GenerationsInverseTable is the table name for the Generation entity.
	// It exists in this package in order to avoid circular dependency with the "generation" package.
	GenerationsInverseTable = "generations"
	// GenerationsColumn is the table column denoting the generations relation/edge.
	GenerationsColumn = "api_token_id"
	// UpscalesTable is the table that holds the upscales relation/edge.
	UpscalesTable = "upscales"
	// UpscalesInverseTable is the table name for the Upscale entity.
	// It exists in this package in order to avoid circular dependency with the "upscale" package.
	UpscalesInverseTable = "upscales"
	// UpscalesColumn is the table column denoting the upscales relation/edge.
	UpscalesColumn = "api_token_id"
)

// Columns holds all SQL columns for apitoken fields.
var Columns = []string{
	FieldID,
	FieldHashedToken,
	FieldName,
	FieldShortString,
	FieldIsActive,
	FieldUses,
	FieldCreditsSpent,
	FieldUserID,
	FieldLastUsedAt,
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
	// DefaultIsActive holds the default value on creation for the "is_active" field.
	DefaultIsActive bool
	// DefaultUses holds the default value on creation for the "uses" field.
	DefaultUses int
	// DefaultCreditsSpent holds the default value on creation for the "credits_spent" field.
	DefaultCreditsSpent int
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
