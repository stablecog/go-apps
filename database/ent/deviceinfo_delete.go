// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/stablecog/go-apps/database/ent/deviceinfo"
	"github.com/stablecog/go-apps/database/ent/predicate"
)

// DeviceInfoDelete is the builder for deleting a DeviceInfo entity.
type DeviceInfoDelete struct {
	config
	hooks    []Hook
	mutation *DeviceInfoMutation
}

// Where appends a list predicates to the DeviceInfoDelete builder.
func (did *DeviceInfoDelete) Where(ps ...predicate.DeviceInfo) *DeviceInfoDelete {
	did.mutation.Where(ps...)
	return did
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (did *DeviceInfoDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, DeviceInfoMutation](ctx, did.sqlExec, did.mutation, did.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (did *DeviceInfoDelete) ExecX(ctx context.Context) int {
	n, err := did.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (did *DeviceInfoDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: deviceinfo.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: deviceinfo.FieldID,
			},
		},
	}
	if ps := did.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, did.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	did.mutation.done = true
	return affected, err
}

// DeviceInfoDeleteOne is the builder for deleting a single DeviceInfo entity.
type DeviceInfoDeleteOne struct {
	did *DeviceInfoDelete
}

// Exec executes the deletion query.
func (dido *DeviceInfoDeleteOne) Exec(ctx context.Context) error {
	n, err := dido.did.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{deviceinfo.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (dido *DeviceInfoDeleteOne) ExecX(ctx context.Context) {
	dido.did.ExecX(ctx)
}
