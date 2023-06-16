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
	"github.com/stablecog/sc-go/database/ent/predicate"
	"github.com/stablecog/sc-go/database/ent/voiceover"
	"github.com/stablecog/sc-go/database/ent/voiceovermodel"
	"github.com/stablecog/sc-go/database/ent/voiceoverspeaker"
)

// VoiceoverSpeakerUpdate is the builder for updating VoiceoverSpeaker entities.
type VoiceoverSpeakerUpdate struct {
	config
	hooks     []Hook
	mutation  *VoiceoverSpeakerMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the VoiceoverSpeakerUpdate builder.
func (vsu *VoiceoverSpeakerUpdate) Where(ps ...predicate.VoiceoverSpeaker) *VoiceoverSpeakerUpdate {
	vsu.mutation.Where(ps...)
	return vsu
}

// SetNameInWorker sets the "name_in_worker" field.
func (vsu *VoiceoverSpeakerUpdate) SetNameInWorker(s string) *VoiceoverSpeakerUpdate {
	vsu.mutation.SetNameInWorker(s)
	return vsu
}

// SetIsActive sets the "is_active" field.
func (vsu *VoiceoverSpeakerUpdate) SetIsActive(b bool) *VoiceoverSpeakerUpdate {
	vsu.mutation.SetIsActive(b)
	return vsu
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (vsu *VoiceoverSpeakerUpdate) SetNillableIsActive(b *bool) *VoiceoverSpeakerUpdate {
	if b != nil {
		vsu.SetIsActive(*b)
	}
	return vsu
}

// SetIsDefault sets the "is_default" field.
func (vsu *VoiceoverSpeakerUpdate) SetIsDefault(b bool) *VoiceoverSpeakerUpdate {
	vsu.mutation.SetIsDefault(b)
	return vsu
}

// SetNillableIsDefault sets the "is_default" field if the given value is not nil.
func (vsu *VoiceoverSpeakerUpdate) SetNillableIsDefault(b *bool) *VoiceoverSpeakerUpdate {
	if b != nil {
		vsu.SetIsDefault(*b)
	}
	return vsu
}

// SetIsHidden sets the "is_hidden" field.
func (vsu *VoiceoverSpeakerUpdate) SetIsHidden(b bool) *VoiceoverSpeakerUpdate {
	vsu.mutation.SetIsHidden(b)
	return vsu
}

// SetNillableIsHidden sets the "is_hidden" field if the given value is not nil.
func (vsu *VoiceoverSpeakerUpdate) SetNillableIsHidden(b *bool) *VoiceoverSpeakerUpdate {
	if b != nil {
		vsu.SetIsHidden(*b)
	}
	return vsu
}

// SetLocale sets the "locale" field.
func (vsu *VoiceoverSpeakerUpdate) SetLocale(s string) *VoiceoverSpeakerUpdate {
	vsu.mutation.SetLocale(s)
	return vsu
}

// SetNillableLocale sets the "locale" field if the given value is not nil.
func (vsu *VoiceoverSpeakerUpdate) SetNillableLocale(s *string) *VoiceoverSpeakerUpdate {
	if s != nil {
		vsu.SetLocale(*s)
	}
	return vsu
}

// SetModelID sets the "model_id" field.
func (vsu *VoiceoverSpeakerUpdate) SetModelID(u uuid.UUID) *VoiceoverSpeakerUpdate {
	vsu.mutation.SetModelID(u)
	return vsu
}

// SetUpdatedAt sets the "updated_at" field.
func (vsu *VoiceoverSpeakerUpdate) SetUpdatedAt(t time.Time) *VoiceoverSpeakerUpdate {
	vsu.mutation.SetUpdatedAt(t)
	return vsu
}

// AddVoiceoverIDs adds the "voiceovers" edge to the Voiceover entity by IDs.
func (vsu *VoiceoverSpeakerUpdate) AddVoiceoverIDs(ids ...uuid.UUID) *VoiceoverSpeakerUpdate {
	vsu.mutation.AddVoiceoverIDs(ids...)
	return vsu
}

// AddVoiceovers adds the "voiceovers" edges to the Voiceover entity.
func (vsu *VoiceoverSpeakerUpdate) AddVoiceovers(v ...*Voiceover) *VoiceoverSpeakerUpdate {
	ids := make([]uuid.UUID, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return vsu.AddVoiceoverIDs(ids...)
}

// SetVoiceoverModelsID sets the "voiceover_models" edge to the VoiceoverModel entity by ID.
func (vsu *VoiceoverSpeakerUpdate) SetVoiceoverModelsID(id uuid.UUID) *VoiceoverSpeakerUpdate {
	vsu.mutation.SetVoiceoverModelsID(id)
	return vsu
}

// SetVoiceoverModels sets the "voiceover_models" edge to the VoiceoverModel entity.
func (vsu *VoiceoverSpeakerUpdate) SetVoiceoverModels(v *VoiceoverModel) *VoiceoverSpeakerUpdate {
	return vsu.SetVoiceoverModelsID(v.ID)
}

// Mutation returns the VoiceoverSpeakerMutation object of the builder.
func (vsu *VoiceoverSpeakerUpdate) Mutation() *VoiceoverSpeakerMutation {
	return vsu.mutation
}

// ClearVoiceovers clears all "voiceovers" edges to the Voiceover entity.
func (vsu *VoiceoverSpeakerUpdate) ClearVoiceovers() *VoiceoverSpeakerUpdate {
	vsu.mutation.ClearVoiceovers()
	return vsu
}

// RemoveVoiceoverIDs removes the "voiceovers" edge to Voiceover entities by IDs.
func (vsu *VoiceoverSpeakerUpdate) RemoveVoiceoverIDs(ids ...uuid.UUID) *VoiceoverSpeakerUpdate {
	vsu.mutation.RemoveVoiceoverIDs(ids...)
	return vsu
}

// RemoveVoiceovers removes "voiceovers" edges to Voiceover entities.
func (vsu *VoiceoverSpeakerUpdate) RemoveVoiceovers(v ...*Voiceover) *VoiceoverSpeakerUpdate {
	ids := make([]uuid.UUID, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return vsu.RemoveVoiceoverIDs(ids...)
}

// ClearVoiceoverModels clears the "voiceover_models" edge to the VoiceoverModel entity.
func (vsu *VoiceoverSpeakerUpdate) ClearVoiceoverModels() *VoiceoverSpeakerUpdate {
	vsu.mutation.ClearVoiceoverModels()
	return vsu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (vsu *VoiceoverSpeakerUpdate) Save(ctx context.Context) (int, error) {
	vsu.defaults()
	return withHooks[int, VoiceoverSpeakerMutation](ctx, vsu.sqlSave, vsu.mutation, vsu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (vsu *VoiceoverSpeakerUpdate) SaveX(ctx context.Context) int {
	affected, err := vsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (vsu *VoiceoverSpeakerUpdate) Exec(ctx context.Context) error {
	_, err := vsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vsu *VoiceoverSpeakerUpdate) ExecX(ctx context.Context) {
	if err := vsu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vsu *VoiceoverSpeakerUpdate) defaults() {
	if _, ok := vsu.mutation.UpdatedAt(); !ok {
		v := voiceoverspeaker.UpdateDefaultUpdatedAt()
		vsu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vsu *VoiceoverSpeakerUpdate) check() error {
	if _, ok := vsu.mutation.VoiceoverModelsID(); vsu.mutation.VoiceoverModelsCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "VoiceoverSpeaker.voiceover_models"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (vsu *VoiceoverSpeakerUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *VoiceoverSpeakerUpdate {
	vsu.modifiers = append(vsu.modifiers, modifiers...)
	return vsu
}

func (vsu *VoiceoverSpeakerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := vsu.check(); err != nil {
		return n, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   voiceoverspeaker.Table,
			Columns: voiceoverspeaker.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: voiceoverspeaker.FieldID,
			},
		},
	}
	if ps := vsu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vsu.mutation.NameInWorker(); ok {
		_spec.SetField(voiceoverspeaker.FieldNameInWorker, field.TypeString, value)
	}
	if value, ok := vsu.mutation.IsActive(); ok {
		_spec.SetField(voiceoverspeaker.FieldIsActive, field.TypeBool, value)
	}
	if value, ok := vsu.mutation.IsDefault(); ok {
		_spec.SetField(voiceoverspeaker.FieldIsDefault, field.TypeBool, value)
	}
	if value, ok := vsu.mutation.IsHidden(); ok {
		_spec.SetField(voiceoverspeaker.FieldIsHidden, field.TypeBool, value)
	}
	if value, ok := vsu.mutation.Locale(); ok {
		_spec.SetField(voiceoverspeaker.FieldLocale, field.TypeString, value)
	}
	if value, ok := vsu.mutation.UpdatedAt(); ok {
		_spec.SetField(voiceoverspeaker.FieldUpdatedAt, field.TypeTime, value)
	}
	if vsu.mutation.VoiceoversCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   voiceoverspeaker.VoiceoversTable,
			Columns: []string{voiceoverspeaker.VoiceoversColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: voiceover.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vsu.mutation.RemovedVoiceoversIDs(); len(nodes) > 0 && !vsu.mutation.VoiceoversCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   voiceoverspeaker.VoiceoversTable,
			Columns: []string{voiceoverspeaker.VoiceoversColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: voiceover.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vsu.mutation.VoiceoversIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   voiceoverspeaker.VoiceoversTable,
			Columns: []string{voiceoverspeaker.VoiceoversColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: voiceover.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if vsu.mutation.VoiceoverModelsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   voiceoverspeaker.VoiceoverModelsTable,
			Columns: []string{voiceoverspeaker.VoiceoverModelsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: voiceovermodel.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vsu.mutation.VoiceoverModelsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   voiceoverspeaker.VoiceoverModelsTable,
			Columns: []string{voiceoverspeaker.VoiceoverModelsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: voiceovermodel.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(vsu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, vsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{voiceoverspeaker.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	vsu.mutation.done = true
	return n, nil
}

// VoiceoverSpeakerUpdateOne is the builder for updating a single VoiceoverSpeaker entity.
type VoiceoverSpeakerUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *VoiceoverSpeakerMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetNameInWorker sets the "name_in_worker" field.
func (vsuo *VoiceoverSpeakerUpdateOne) SetNameInWorker(s string) *VoiceoverSpeakerUpdateOne {
	vsuo.mutation.SetNameInWorker(s)
	return vsuo
}

// SetIsActive sets the "is_active" field.
func (vsuo *VoiceoverSpeakerUpdateOne) SetIsActive(b bool) *VoiceoverSpeakerUpdateOne {
	vsuo.mutation.SetIsActive(b)
	return vsuo
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (vsuo *VoiceoverSpeakerUpdateOne) SetNillableIsActive(b *bool) *VoiceoverSpeakerUpdateOne {
	if b != nil {
		vsuo.SetIsActive(*b)
	}
	return vsuo
}

// SetIsDefault sets the "is_default" field.
func (vsuo *VoiceoverSpeakerUpdateOne) SetIsDefault(b bool) *VoiceoverSpeakerUpdateOne {
	vsuo.mutation.SetIsDefault(b)
	return vsuo
}

// SetNillableIsDefault sets the "is_default" field if the given value is not nil.
func (vsuo *VoiceoverSpeakerUpdateOne) SetNillableIsDefault(b *bool) *VoiceoverSpeakerUpdateOne {
	if b != nil {
		vsuo.SetIsDefault(*b)
	}
	return vsuo
}

// SetIsHidden sets the "is_hidden" field.
func (vsuo *VoiceoverSpeakerUpdateOne) SetIsHidden(b bool) *VoiceoverSpeakerUpdateOne {
	vsuo.mutation.SetIsHidden(b)
	return vsuo
}

// SetNillableIsHidden sets the "is_hidden" field if the given value is not nil.
func (vsuo *VoiceoverSpeakerUpdateOne) SetNillableIsHidden(b *bool) *VoiceoverSpeakerUpdateOne {
	if b != nil {
		vsuo.SetIsHidden(*b)
	}
	return vsuo
}

// SetLocale sets the "locale" field.
func (vsuo *VoiceoverSpeakerUpdateOne) SetLocale(s string) *VoiceoverSpeakerUpdateOne {
	vsuo.mutation.SetLocale(s)
	return vsuo
}

// SetNillableLocale sets the "locale" field if the given value is not nil.
func (vsuo *VoiceoverSpeakerUpdateOne) SetNillableLocale(s *string) *VoiceoverSpeakerUpdateOne {
	if s != nil {
		vsuo.SetLocale(*s)
	}
	return vsuo
}

// SetModelID sets the "model_id" field.
func (vsuo *VoiceoverSpeakerUpdateOne) SetModelID(u uuid.UUID) *VoiceoverSpeakerUpdateOne {
	vsuo.mutation.SetModelID(u)
	return vsuo
}

// SetUpdatedAt sets the "updated_at" field.
func (vsuo *VoiceoverSpeakerUpdateOne) SetUpdatedAt(t time.Time) *VoiceoverSpeakerUpdateOne {
	vsuo.mutation.SetUpdatedAt(t)
	return vsuo
}

// AddVoiceoverIDs adds the "voiceovers" edge to the Voiceover entity by IDs.
func (vsuo *VoiceoverSpeakerUpdateOne) AddVoiceoverIDs(ids ...uuid.UUID) *VoiceoverSpeakerUpdateOne {
	vsuo.mutation.AddVoiceoverIDs(ids...)
	return vsuo
}

// AddVoiceovers adds the "voiceovers" edges to the Voiceover entity.
func (vsuo *VoiceoverSpeakerUpdateOne) AddVoiceovers(v ...*Voiceover) *VoiceoverSpeakerUpdateOne {
	ids := make([]uuid.UUID, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return vsuo.AddVoiceoverIDs(ids...)
}

// SetVoiceoverModelsID sets the "voiceover_models" edge to the VoiceoverModel entity by ID.
func (vsuo *VoiceoverSpeakerUpdateOne) SetVoiceoverModelsID(id uuid.UUID) *VoiceoverSpeakerUpdateOne {
	vsuo.mutation.SetVoiceoverModelsID(id)
	return vsuo
}

// SetVoiceoverModels sets the "voiceover_models" edge to the VoiceoverModel entity.
func (vsuo *VoiceoverSpeakerUpdateOne) SetVoiceoverModels(v *VoiceoverModel) *VoiceoverSpeakerUpdateOne {
	return vsuo.SetVoiceoverModelsID(v.ID)
}

// Mutation returns the VoiceoverSpeakerMutation object of the builder.
func (vsuo *VoiceoverSpeakerUpdateOne) Mutation() *VoiceoverSpeakerMutation {
	return vsuo.mutation
}

// ClearVoiceovers clears all "voiceovers" edges to the Voiceover entity.
func (vsuo *VoiceoverSpeakerUpdateOne) ClearVoiceovers() *VoiceoverSpeakerUpdateOne {
	vsuo.mutation.ClearVoiceovers()
	return vsuo
}

// RemoveVoiceoverIDs removes the "voiceovers" edge to Voiceover entities by IDs.
func (vsuo *VoiceoverSpeakerUpdateOne) RemoveVoiceoverIDs(ids ...uuid.UUID) *VoiceoverSpeakerUpdateOne {
	vsuo.mutation.RemoveVoiceoverIDs(ids...)
	return vsuo
}

// RemoveVoiceovers removes "voiceovers" edges to Voiceover entities.
func (vsuo *VoiceoverSpeakerUpdateOne) RemoveVoiceovers(v ...*Voiceover) *VoiceoverSpeakerUpdateOne {
	ids := make([]uuid.UUID, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return vsuo.RemoveVoiceoverIDs(ids...)
}

// ClearVoiceoverModels clears the "voiceover_models" edge to the VoiceoverModel entity.
func (vsuo *VoiceoverSpeakerUpdateOne) ClearVoiceoverModels() *VoiceoverSpeakerUpdateOne {
	vsuo.mutation.ClearVoiceoverModels()
	return vsuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (vsuo *VoiceoverSpeakerUpdateOne) Select(field string, fields ...string) *VoiceoverSpeakerUpdateOne {
	vsuo.fields = append([]string{field}, fields...)
	return vsuo
}

// Save executes the query and returns the updated VoiceoverSpeaker entity.
func (vsuo *VoiceoverSpeakerUpdateOne) Save(ctx context.Context) (*VoiceoverSpeaker, error) {
	vsuo.defaults()
	return withHooks[*VoiceoverSpeaker, VoiceoverSpeakerMutation](ctx, vsuo.sqlSave, vsuo.mutation, vsuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (vsuo *VoiceoverSpeakerUpdateOne) SaveX(ctx context.Context) *VoiceoverSpeaker {
	node, err := vsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (vsuo *VoiceoverSpeakerUpdateOne) Exec(ctx context.Context) error {
	_, err := vsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vsuo *VoiceoverSpeakerUpdateOne) ExecX(ctx context.Context) {
	if err := vsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vsuo *VoiceoverSpeakerUpdateOne) defaults() {
	if _, ok := vsuo.mutation.UpdatedAt(); !ok {
		v := voiceoverspeaker.UpdateDefaultUpdatedAt()
		vsuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vsuo *VoiceoverSpeakerUpdateOne) check() error {
	if _, ok := vsuo.mutation.VoiceoverModelsID(); vsuo.mutation.VoiceoverModelsCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "VoiceoverSpeaker.voiceover_models"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (vsuo *VoiceoverSpeakerUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *VoiceoverSpeakerUpdateOne {
	vsuo.modifiers = append(vsuo.modifiers, modifiers...)
	return vsuo
}

func (vsuo *VoiceoverSpeakerUpdateOne) sqlSave(ctx context.Context) (_node *VoiceoverSpeaker, err error) {
	if err := vsuo.check(); err != nil {
		return _node, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   voiceoverspeaker.Table,
			Columns: voiceoverspeaker.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: voiceoverspeaker.FieldID,
			},
		},
	}
	id, ok := vsuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "VoiceoverSpeaker.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := vsuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, voiceoverspeaker.FieldID)
		for _, f := range fields {
			if !voiceoverspeaker.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != voiceoverspeaker.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := vsuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vsuo.mutation.NameInWorker(); ok {
		_spec.SetField(voiceoverspeaker.FieldNameInWorker, field.TypeString, value)
	}
	if value, ok := vsuo.mutation.IsActive(); ok {
		_spec.SetField(voiceoverspeaker.FieldIsActive, field.TypeBool, value)
	}
	if value, ok := vsuo.mutation.IsDefault(); ok {
		_spec.SetField(voiceoverspeaker.FieldIsDefault, field.TypeBool, value)
	}
	if value, ok := vsuo.mutation.IsHidden(); ok {
		_spec.SetField(voiceoverspeaker.FieldIsHidden, field.TypeBool, value)
	}
	if value, ok := vsuo.mutation.Locale(); ok {
		_spec.SetField(voiceoverspeaker.FieldLocale, field.TypeString, value)
	}
	if value, ok := vsuo.mutation.UpdatedAt(); ok {
		_spec.SetField(voiceoverspeaker.FieldUpdatedAt, field.TypeTime, value)
	}
	if vsuo.mutation.VoiceoversCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   voiceoverspeaker.VoiceoversTable,
			Columns: []string{voiceoverspeaker.VoiceoversColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: voiceover.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vsuo.mutation.RemovedVoiceoversIDs(); len(nodes) > 0 && !vsuo.mutation.VoiceoversCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   voiceoverspeaker.VoiceoversTable,
			Columns: []string{voiceoverspeaker.VoiceoversColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: voiceover.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vsuo.mutation.VoiceoversIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   voiceoverspeaker.VoiceoversTable,
			Columns: []string{voiceoverspeaker.VoiceoversColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: voiceover.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if vsuo.mutation.VoiceoverModelsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   voiceoverspeaker.VoiceoverModelsTable,
			Columns: []string{voiceoverspeaker.VoiceoverModelsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: voiceovermodel.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vsuo.mutation.VoiceoverModelsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   voiceoverspeaker.VoiceoverModelsTable,
			Columns: []string{voiceoverspeaker.VoiceoverModelsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: voiceovermodel.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(vsuo.modifiers...)
	_node = &VoiceoverSpeaker{config: vsuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, vsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{voiceoverspeaker.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	vsuo.mutation.done = true
	return _node, nil
}