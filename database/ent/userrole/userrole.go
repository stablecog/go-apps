// Code generated by ent, DO NOT EDIT.

package userrole

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the userrole type in the database.
	Label = "user_role"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldRoleName holds the string denoting the role_name field in the database.
	FieldRoleName = "role_name"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// Table holds the table name of the userrole in the database.
	Table = "user_roles"
	// UsersTable is the table that holds the users relation/edge.
	UsersTable = "user_roles"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "users"
	// UsersColumn is the table column denoting the users relation/edge.
	UsersColumn = "user_id"
)

// Columns holds all SQL columns for userrole fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldRoleName,
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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// RoleName defines the type for the "role_name" enum field.
type RoleName string

// RoleName values.
const (
	RoleNameADMIN RoleName = "ADMIN"
	RoleNamePRO   RoleName = "PRO"
)

func (rn RoleName) String() string {
	return string(rn)
}

// RoleNameValidator is a validator for the "role_name" field enum values. It is called by the builders before save.
func RoleNameValidator(rn RoleName) error {
	switch rn {
	case RoleNameADMIN, RoleNamePRO:
		return nil
	default:
		return fmt.Errorf("userrole: invalid enum value for role_name field: %q", rn)
	}
}
