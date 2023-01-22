// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/stablecog/go-apps/database/ent/deviceinfo"
	"github.com/stablecog/go-apps/database/ent/predicate"
	"github.com/stablecog/go-apps/database/ent/upscale"
	"github.com/stablecog/go-apps/database/ent/upscalemodel"
	"github.com/stablecog/go-apps/database/ent/upscaleoutput"
	"github.com/stablecog/go-apps/database/ent/user"
)

// UpscaleUpdate is the builder for updating Upscale entities.
type UpscaleUpdate struct {
	config
	hooks    []Hook
	mutation *UpscaleMutation
}

// Where appends a list predicates to the UpscaleUpdate builder.
func (uu *UpscaleUpdate) Where(ps ...predicate.Upscale) *UpscaleUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetWidth sets the "width" field.
func (uu *UpscaleUpdate) SetWidth(i int) *UpscaleUpdate {
	uu.mutation.ResetWidth()
	uu.mutation.SetWidth(i)
	return uu
}

// AddWidth adds i to the "width" field.
func (uu *UpscaleUpdate) AddWidth(i int) *UpscaleUpdate {
	uu.mutation.AddWidth(i)
	return uu
}

// SetHeight sets the "height" field.
func (uu *UpscaleUpdate) SetHeight(i int) *UpscaleUpdate {
	uu.mutation.ResetHeight()
	uu.mutation.SetHeight(i)
	return uu
}

// AddHeight adds i to the "height" field.
func (uu *UpscaleUpdate) AddHeight(i int) *UpscaleUpdate {
	uu.mutation.AddHeight(i)
	return uu
}

// SetScale sets the "scale" field.
func (uu *UpscaleUpdate) SetScale(i int) *UpscaleUpdate {
	uu.mutation.ResetScale()
	uu.mutation.SetScale(i)
	return uu
}

// AddScale adds i to the "scale" field.
func (uu *UpscaleUpdate) AddScale(i int) *UpscaleUpdate {
	uu.mutation.AddScale(i)
	return uu
}

// SetDurationMs sets the "duration_ms" field.
func (uu *UpscaleUpdate) SetDurationMs(i int) *UpscaleUpdate {
	uu.mutation.ResetDurationMs()
	uu.mutation.SetDurationMs(i)
	return uu
}

// AddDurationMs adds i to the "duration_ms" field.
func (uu *UpscaleUpdate) AddDurationMs(i int) *UpscaleUpdate {
	uu.mutation.AddDurationMs(i)
	return uu
}

// SetCountryCode sets the "country_code" field.
func (uu *UpscaleUpdate) SetCountryCode(s string) *UpscaleUpdate {
	uu.mutation.SetCountryCode(s)
	return uu
}

// SetStatus sets the "status" field.
func (uu *UpscaleUpdate) SetStatus(u upscale.Status) *UpscaleUpdate {
	uu.mutation.SetStatus(u)
	return uu
}

// SetFailureReason sets the "failure_reason" field.
func (uu *UpscaleUpdate) SetFailureReason(s string) *UpscaleUpdate {
	uu.mutation.SetFailureReason(s)
	return uu
}

// SetModelID sets the "model_id" field.
func (uu *UpscaleUpdate) SetModelID(u uuid.UUID) *UpscaleUpdate {
	uu.mutation.SetModelID(u)
	return uu
}

// SetUserID sets the "user_id" field.
func (uu *UpscaleUpdate) SetUserID(u uuid.UUID) *UpscaleUpdate {
	uu.mutation.SetUserID(u)
	return uu
}

// SetDeviceInfoID sets the "device_info_id" field.
func (uu *UpscaleUpdate) SetDeviceInfoID(u uuid.UUID) *UpscaleUpdate {
	uu.mutation.SetDeviceInfoID(u)
	return uu
}

// SetUpdatedAt sets the "updated_at" field.
func (uu *UpscaleUpdate) SetUpdatedAt(t time.Time) *UpscaleUpdate {
	uu.mutation.SetUpdatedAt(t)
	return uu
}

// SetUser sets the "user" edge to the User entity.
func (uu *UpscaleUpdate) SetUser(u *User) *UpscaleUpdate {
	return uu.SetUserID(u.ID)
}

// SetDeviceInfo sets the "device_info" edge to the DeviceInfo entity.
func (uu *UpscaleUpdate) SetDeviceInfo(d *DeviceInfo) *UpscaleUpdate {
	return uu.SetDeviceInfoID(d.ID)
}

// SetUpscaleModelsID sets the "upscale_models" edge to the UpscaleModel entity by ID.
func (uu *UpscaleUpdate) SetUpscaleModelsID(id uuid.UUID) *UpscaleUpdate {
	uu.mutation.SetUpscaleModelsID(id)
	return uu
}

// SetUpscaleModels sets the "upscale_models" edge to the UpscaleModel entity.
func (uu *UpscaleUpdate) SetUpscaleModels(u *UpscaleModel) *UpscaleUpdate {
	return uu.SetUpscaleModelsID(u.ID)
}

// AddUpscaleOutputIDs adds the "upscale_outputs" edge to the UpscaleOutput entity by IDs.
func (uu *UpscaleUpdate) AddUpscaleOutputIDs(ids ...uuid.UUID) *UpscaleUpdate {
	uu.mutation.AddUpscaleOutputIDs(ids...)
	return uu
}

// AddUpscaleOutputs adds the "upscale_outputs" edges to the UpscaleOutput entity.
func (uu *UpscaleUpdate) AddUpscaleOutputs(u ...*UpscaleOutput) *UpscaleUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.AddUpscaleOutputIDs(ids...)
}

// Mutation returns the UpscaleMutation object of the builder.
func (uu *UpscaleUpdate) Mutation() *UpscaleMutation {
	return uu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (uu *UpscaleUpdate) ClearUser() *UpscaleUpdate {
	uu.mutation.ClearUser()
	return uu
}

// ClearDeviceInfo clears the "device_info" edge to the DeviceInfo entity.
func (uu *UpscaleUpdate) ClearDeviceInfo() *UpscaleUpdate {
	uu.mutation.ClearDeviceInfo()
	return uu
}

// ClearUpscaleModels clears the "upscale_models" edge to the UpscaleModel entity.
func (uu *UpscaleUpdate) ClearUpscaleModels() *UpscaleUpdate {
	uu.mutation.ClearUpscaleModels()
	return uu
}

// ClearUpscaleOutputs clears all "upscale_outputs" edges to the UpscaleOutput entity.
func (uu *UpscaleUpdate) ClearUpscaleOutputs() *UpscaleUpdate {
	uu.mutation.ClearUpscaleOutputs()
	return uu
}

// RemoveUpscaleOutputIDs removes the "upscale_outputs" edge to UpscaleOutput entities by IDs.
func (uu *UpscaleUpdate) RemoveUpscaleOutputIDs(ids ...uuid.UUID) *UpscaleUpdate {
	uu.mutation.RemoveUpscaleOutputIDs(ids...)
	return uu
}

// RemoveUpscaleOutputs removes "upscale_outputs" edges to UpscaleOutput entities.
func (uu *UpscaleUpdate) RemoveUpscaleOutputs(u ...*UpscaleOutput) *UpscaleUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.RemoveUpscaleOutputIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UpscaleUpdate) Save(ctx context.Context) (int, error) {
	uu.defaults()
	return withHooks[int, UpscaleMutation](ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UpscaleUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UpscaleUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UpscaleUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uu *UpscaleUpdate) defaults() {
	if _, ok := uu.mutation.UpdatedAt(); !ok {
		v := upscale.UpdateDefaultUpdatedAt()
		uu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UpscaleUpdate) check() error {
	if v, ok := uu.mutation.Status(); ok {
		if err := upscale.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Upscale.status": %w`, err)}
		}
	}
	if _, ok := uu.mutation.UserID(); uu.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Upscale.user"`)
	}
	if _, ok := uu.mutation.DeviceInfoID(); uu.mutation.DeviceInfoCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Upscale.device_info"`)
	}
	if _, ok := uu.mutation.UpscaleModelsID(); uu.mutation.UpscaleModelsCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Upscale.upscale_models"`)
	}
	return nil
}

func (uu *UpscaleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   upscale.Table,
			Columns: upscale.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: upscale.FieldID,
			},
		},
	}
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Width(); ok {
		_spec.SetField(upscale.FieldWidth, field.TypeInt, value)
	}
	if value, ok := uu.mutation.AddedWidth(); ok {
		_spec.AddField(upscale.FieldWidth, field.TypeInt, value)
	}
	if value, ok := uu.mutation.Height(); ok {
		_spec.SetField(upscale.FieldHeight, field.TypeInt, value)
	}
	if value, ok := uu.mutation.AddedHeight(); ok {
		_spec.AddField(upscale.FieldHeight, field.TypeInt, value)
	}
	if value, ok := uu.mutation.Scale(); ok {
		_spec.SetField(upscale.FieldScale, field.TypeInt, value)
	}
	if value, ok := uu.mutation.AddedScale(); ok {
		_spec.AddField(upscale.FieldScale, field.TypeInt, value)
	}
	if value, ok := uu.mutation.DurationMs(); ok {
		_spec.SetField(upscale.FieldDurationMs, field.TypeInt, value)
	}
	if value, ok := uu.mutation.AddedDurationMs(); ok {
		_spec.AddField(upscale.FieldDurationMs, field.TypeInt, value)
	}
	if value, ok := uu.mutation.CountryCode(); ok {
		_spec.SetField(upscale.FieldCountryCode, field.TypeString, value)
	}
	if value, ok := uu.mutation.Status(); ok {
		_spec.SetField(upscale.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := uu.mutation.FailureReason(); ok {
		_spec.SetField(upscale.FieldFailureReason, field.TypeString, value)
	}
	if value, ok := uu.mutation.UpdatedAt(); ok {
		_spec.SetField(upscale.FieldUpdatedAt, field.TypeTime, value)
	}
	if uu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   upscale.UserTable,
			Columns: []string{upscale.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   upscale.UserTable,
			Columns: []string{upscale.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.DeviceInfoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   upscale.DeviceInfoTable,
			Columns: []string{upscale.DeviceInfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: deviceinfo.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.DeviceInfoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   upscale.DeviceInfoTable,
			Columns: []string{upscale.DeviceInfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: deviceinfo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.UpscaleModelsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   upscale.UpscaleModelsTable,
			Columns: []string{upscale.UpscaleModelsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: upscalemodel.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.UpscaleModelsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   upscale.UpscaleModelsTable,
			Columns: []string{upscale.UpscaleModelsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: upscalemodel.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.UpscaleOutputsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   upscale.UpscaleOutputsTable,
			Columns: []string{upscale.UpscaleOutputsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: upscaleoutput.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedUpscaleOutputsIDs(); len(nodes) > 0 && !uu.mutation.UpscaleOutputsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   upscale.UpscaleOutputsTable,
			Columns: []string{upscale.UpscaleOutputsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: upscaleoutput.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.UpscaleOutputsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   upscale.UpscaleOutputsTable,
			Columns: []string{upscale.UpscaleOutputsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: upscaleoutput.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{upscale.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UpscaleUpdateOne is the builder for updating a single Upscale entity.
type UpscaleUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UpscaleMutation
}

// SetWidth sets the "width" field.
func (uuo *UpscaleUpdateOne) SetWidth(i int) *UpscaleUpdateOne {
	uuo.mutation.ResetWidth()
	uuo.mutation.SetWidth(i)
	return uuo
}

// AddWidth adds i to the "width" field.
func (uuo *UpscaleUpdateOne) AddWidth(i int) *UpscaleUpdateOne {
	uuo.mutation.AddWidth(i)
	return uuo
}

// SetHeight sets the "height" field.
func (uuo *UpscaleUpdateOne) SetHeight(i int) *UpscaleUpdateOne {
	uuo.mutation.ResetHeight()
	uuo.mutation.SetHeight(i)
	return uuo
}

// AddHeight adds i to the "height" field.
func (uuo *UpscaleUpdateOne) AddHeight(i int) *UpscaleUpdateOne {
	uuo.mutation.AddHeight(i)
	return uuo
}

// SetScale sets the "scale" field.
func (uuo *UpscaleUpdateOne) SetScale(i int) *UpscaleUpdateOne {
	uuo.mutation.ResetScale()
	uuo.mutation.SetScale(i)
	return uuo
}

// AddScale adds i to the "scale" field.
func (uuo *UpscaleUpdateOne) AddScale(i int) *UpscaleUpdateOne {
	uuo.mutation.AddScale(i)
	return uuo
}

// SetDurationMs sets the "duration_ms" field.
func (uuo *UpscaleUpdateOne) SetDurationMs(i int) *UpscaleUpdateOne {
	uuo.mutation.ResetDurationMs()
	uuo.mutation.SetDurationMs(i)
	return uuo
}

// AddDurationMs adds i to the "duration_ms" field.
func (uuo *UpscaleUpdateOne) AddDurationMs(i int) *UpscaleUpdateOne {
	uuo.mutation.AddDurationMs(i)
	return uuo
}

// SetCountryCode sets the "country_code" field.
func (uuo *UpscaleUpdateOne) SetCountryCode(s string) *UpscaleUpdateOne {
	uuo.mutation.SetCountryCode(s)
	return uuo
}

// SetStatus sets the "status" field.
func (uuo *UpscaleUpdateOne) SetStatus(u upscale.Status) *UpscaleUpdateOne {
	uuo.mutation.SetStatus(u)
	return uuo
}

// SetFailureReason sets the "failure_reason" field.
func (uuo *UpscaleUpdateOne) SetFailureReason(s string) *UpscaleUpdateOne {
	uuo.mutation.SetFailureReason(s)
	return uuo
}

// SetModelID sets the "model_id" field.
func (uuo *UpscaleUpdateOne) SetModelID(u uuid.UUID) *UpscaleUpdateOne {
	uuo.mutation.SetModelID(u)
	return uuo
}

// SetUserID sets the "user_id" field.
func (uuo *UpscaleUpdateOne) SetUserID(u uuid.UUID) *UpscaleUpdateOne {
	uuo.mutation.SetUserID(u)
	return uuo
}

// SetDeviceInfoID sets the "device_info_id" field.
func (uuo *UpscaleUpdateOne) SetDeviceInfoID(u uuid.UUID) *UpscaleUpdateOne {
	uuo.mutation.SetDeviceInfoID(u)
	return uuo
}

// SetUpdatedAt sets the "updated_at" field.
func (uuo *UpscaleUpdateOne) SetUpdatedAt(t time.Time) *UpscaleUpdateOne {
	uuo.mutation.SetUpdatedAt(t)
	return uuo
}

// SetUser sets the "user" edge to the User entity.
func (uuo *UpscaleUpdateOne) SetUser(u *User) *UpscaleUpdateOne {
	return uuo.SetUserID(u.ID)
}

// SetDeviceInfo sets the "device_info" edge to the DeviceInfo entity.
func (uuo *UpscaleUpdateOne) SetDeviceInfo(d *DeviceInfo) *UpscaleUpdateOne {
	return uuo.SetDeviceInfoID(d.ID)
}

// SetUpscaleModelsID sets the "upscale_models" edge to the UpscaleModel entity by ID.
func (uuo *UpscaleUpdateOne) SetUpscaleModelsID(id uuid.UUID) *UpscaleUpdateOne {
	uuo.mutation.SetUpscaleModelsID(id)
	return uuo
}

// SetUpscaleModels sets the "upscale_models" edge to the UpscaleModel entity.
func (uuo *UpscaleUpdateOne) SetUpscaleModels(u *UpscaleModel) *UpscaleUpdateOne {
	return uuo.SetUpscaleModelsID(u.ID)
}

// AddUpscaleOutputIDs adds the "upscale_outputs" edge to the UpscaleOutput entity by IDs.
func (uuo *UpscaleUpdateOne) AddUpscaleOutputIDs(ids ...uuid.UUID) *UpscaleUpdateOne {
	uuo.mutation.AddUpscaleOutputIDs(ids...)
	return uuo
}

// AddUpscaleOutputs adds the "upscale_outputs" edges to the UpscaleOutput entity.
func (uuo *UpscaleUpdateOne) AddUpscaleOutputs(u ...*UpscaleOutput) *UpscaleUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.AddUpscaleOutputIDs(ids...)
}

// Mutation returns the UpscaleMutation object of the builder.
func (uuo *UpscaleUpdateOne) Mutation() *UpscaleMutation {
	return uuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (uuo *UpscaleUpdateOne) ClearUser() *UpscaleUpdateOne {
	uuo.mutation.ClearUser()
	return uuo
}

// ClearDeviceInfo clears the "device_info" edge to the DeviceInfo entity.
func (uuo *UpscaleUpdateOne) ClearDeviceInfo() *UpscaleUpdateOne {
	uuo.mutation.ClearDeviceInfo()
	return uuo
}

// ClearUpscaleModels clears the "upscale_models" edge to the UpscaleModel entity.
func (uuo *UpscaleUpdateOne) ClearUpscaleModels() *UpscaleUpdateOne {
	uuo.mutation.ClearUpscaleModels()
	return uuo
}

// ClearUpscaleOutputs clears all "upscale_outputs" edges to the UpscaleOutput entity.
func (uuo *UpscaleUpdateOne) ClearUpscaleOutputs() *UpscaleUpdateOne {
	uuo.mutation.ClearUpscaleOutputs()
	return uuo
}

// RemoveUpscaleOutputIDs removes the "upscale_outputs" edge to UpscaleOutput entities by IDs.
func (uuo *UpscaleUpdateOne) RemoveUpscaleOutputIDs(ids ...uuid.UUID) *UpscaleUpdateOne {
	uuo.mutation.RemoveUpscaleOutputIDs(ids...)
	return uuo
}

// RemoveUpscaleOutputs removes "upscale_outputs" edges to UpscaleOutput entities.
func (uuo *UpscaleUpdateOne) RemoveUpscaleOutputs(u ...*UpscaleOutput) *UpscaleUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.RemoveUpscaleOutputIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UpscaleUpdateOne) Select(field string, fields ...string) *UpscaleUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated Upscale entity.
func (uuo *UpscaleUpdateOne) Save(ctx context.Context) (*Upscale, error) {
	uuo.defaults()
	return withHooks[*Upscale, UpscaleMutation](ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UpscaleUpdateOne) SaveX(ctx context.Context) *Upscale {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UpscaleUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UpscaleUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uuo *UpscaleUpdateOne) defaults() {
	if _, ok := uuo.mutation.UpdatedAt(); !ok {
		v := upscale.UpdateDefaultUpdatedAt()
		uuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UpscaleUpdateOne) check() error {
	if v, ok := uuo.mutation.Status(); ok {
		if err := upscale.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Upscale.status": %w`, err)}
		}
	}
	if _, ok := uuo.mutation.UserID(); uuo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Upscale.user"`)
	}
	if _, ok := uuo.mutation.DeviceInfoID(); uuo.mutation.DeviceInfoCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Upscale.device_info"`)
	}
	if _, ok := uuo.mutation.UpscaleModelsID(); uuo.mutation.UpscaleModelsCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Upscale.upscale_models"`)
	}
	return nil
}

func (uuo *UpscaleUpdateOne) sqlSave(ctx context.Context) (_node *Upscale, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   upscale.Table,
			Columns: upscale.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: upscale.FieldID,
			},
		},
	}
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Upscale.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, upscale.FieldID)
		for _, f := range fields {
			if !upscale.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != upscale.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Width(); ok {
		_spec.SetField(upscale.FieldWidth, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.AddedWidth(); ok {
		_spec.AddField(upscale.FieldWidth, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.Height(); ok {
		_spec.SetField(upscale.FieldHeight, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.AddedHeight(); ok {
		_spec.AddField(upscale.FieldHeight, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.Scale(); ok {
		_spec.SetField(upscale.FieldScale, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.AddedScale(); ok {
		_spec.AddField(upscale.FieldScale, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.DurationMs(); ok {
		_spec.SetField(upscale.FieldDurationMs, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.AddedDurationMs(); ok {
		_spec.AddField(upscale.FieldDurationMs, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.CountryCode(); ok {
		_spec.SetField(upscale.FieldCountryCode, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Status(); ok {
		_spec.SetField(upscale.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := uuo.mutation.FailureReason(); ok {
		_spec.SetField(upscale.FieldFailureReason, field.TypeString, value)
	}
	if value, ok := uuo.mutation.UpdatedAt(); ok {
		_spec.SetField(upscale.FieldUpdatedAt, field.TypeTime, value)
	}
	if uuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   upscale.UserTable,
			Columns: []string{upscale.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   upscale.UserTable,
			Columns: []string{upscale.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.DeviceInfoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   upscale.DeviceInfoTable,
			Columns: []string{upscale.DeviceInfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: deviceinfo.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.DeviceInfoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   upscale.DeviceInfoTable,
			Columns: []string{upscale.DeviceInfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: deviceinfo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.UpscaleModelsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   upscale.UpscaleModelsTable,
			Columns: []string{upscale.UpscaleModelsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: upscalemodel.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.UpscaleModelsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   upscale.UpscaleModelsTable,
			Columns: []string{upscale.UpscaleModelsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: upscalemodel.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.UpscaleOutputsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   upscale.UpscaleOutputsTable,
			Columns: []string{upscale.UpscaleOutputsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: upscaleoutput.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedUpscaleOutputsIDs(); len(nodes) > 0 && !uuo.mutation.UpscaleOutputsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   upscale.UpscaleOutputsTable,
			Columns: []string{upscale.UpscaleOutputsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: upscaleoutput.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.UpscaleOutputsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   upscale.UpscaleOutputsTable,
			Columns: []string{upscale.UpscaleOutputsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: upscaleoutput.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Upscale{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{upscale.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
