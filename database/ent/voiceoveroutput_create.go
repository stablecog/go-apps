// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/stablecog/sc-go/database/ent/voiceover"
	"github.com/stablecog/sc-go/database/ent/voiceoveroutput"
)

// VoiceoverOutputCreate is the builder for creating a VoiceoverOutput entity.
type VoiceoverOutputCreate struct {
	config
	mutation *VoiceoverOutputMutation
	hooks    []Hook
}

// SetAudioPath sets the "audio_path" field.
func (voc *VoiceoverOutputCreate) SetAudioPath(s string) *VoiceoverOutputCreate {
	voc.mutation.SetAudioPath(s)
	return voc
}

// SetVideoPath sets the "video_path" field.
func (voc *VoiceoverOutputCreate) SetVideoPath(s string) *VoiceoverOutputCreate {
	voc.mutation.SetVideoPath(s)
	return voc
}

// SetNillableVideoPath sets the "video_path" field if the given value is not nil.
func (voc *VoiceoverOutputCreate) SetNillableVideoPath(s *string) *VoiceoverOutputCreate {
	if s != nil {
		voc.SetVideoPath(*s)
	}
	return voc
}

// SetAudioArray sets the "audio_array" field.
func (voc *VoiceoverOutputCreate) SetAudioArray(f []float64) *VoiceoverOutputCreate {
	voc.mutation.SetAudioArray(f)
	return voc
}

// SetIsFavorited sets the "is_favorited" field.
func (voc *VoiceoverOutputCreate) SetIsFavorited(b bool) *VoiceoverOutputCreate {
	voc.mutation.SetIsFavorited(b)
	return voc
}

// SetNillableIsFavorited sets the "is_favorited" field if the given value is not nil.
func (voc *VoiceoverOutputCreate) SetNillableIsFavorited(b *bool) *VoiceoverOutputCreate {
	if b != nil {
		voc.SetIsFavorited(*b)
	}
	return voc
}

// SetAudioDuration sets the "audio_duration" field.
func (voc *VoiceoverOutputCreate) SetAudioDuration(f float32) *VoiceoverOutputCreate {
	voc.mutation.SetAudioDuration(f)
	return voc
}

// SetGalleryStatus sets the "gallery_status" field.
func (voc *VoiceoverOutputCreate) SetGalleryStatus(vs voiceoveroutput.GalleryStatus) *VoiceoverOutputCreate {
	voc.mutation.SetGalleryStatus(vs)
	return voc
}

// SetNillableGalleryStatus sets the "gallery_status" field if the given value is not nil.
func (voc *VoiceoverOutputCreate) SetNillableGalleryStatus(vs *voiceoveroutput.GalleryStatus) *VoiceoverOutputCreate {
	if vs != nil {
		voc.SetGalleryStatus(*vs)
	}
	return voc
}

// SetVoiceoverID sets the "voiceover_id" field.
func (voc *VoiceoverOutputCreate) SetVoiceoverID(u uuid.UUID) *VoiceoverOutputCreate {
	voc.mutation.SetVoiceoverID(u)
	return voc
}

// SetDeletedAt sets the "deleted_at" field.
func (voc *VoiceoverOutputCreate) SetDeletedAt(t time.Time) *VoiceoverOutputCreate {
	voc.mutation.SetDeletedAt(t)
	return voc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (voc *VoiceoverOutputCreate) SetNillableDeletedAt(t *time.Time) *VoiceoverOutputCreate {
	if t != nil {
		voc.SetDeletedAt(*t)
	}
	return voc
}

// SetCreatedAt sets the "created_at" field.
func (voc *VoiceoverOutputCreate) SetCreatedAt(t time.Time) *VoiceoverOutputCreate {
	voc.mutation.SetCreatedAt(t)
	return voc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (voc *VoiceoverOutputCreate) SetNillableCreatedAt(t *time.Time) *VoiceoverOutputCreate {
	if t != nil {
		voc.SetCreatedAt(*t)
	}
	return voc
}

// SetUpdatedAt sets the "updated_at" field.
func (voc *VoiceoverOutputCreate) SetUpdatedAt(t time.Time) *VoiceoverOutputCreate {
	voc.mutation.SetUpdatedAt(t)
	return voc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (voc *VoiceoverOutputCreate) SetNillableUpdatedAt(t *time.Time) *VoiceoverOutputCreate {
	if t != nil {
		voc.SetUpdatedAt(*t)
	}
	return voc
}

// SetID sets the "id" field.
func (voc *VoiceoverOutputCreate) SetID(u uuid.UUID) *VoiceoverOutputCreate {
	voc.mutation.SetID(u)
	return voc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (voc *VoiceoverOutputCreate) SetNillableID(u *uuid.UUID) *VoiceoverOutputCreate {
	if u != nil {
		voc.SetID(*u)
	}
	return voc
}

// SetVoiceoversID sets the "voiceovers" edge to the Voiceover entity by ID.
func (voc *VoiceoverOutputCreate) SetVoiceoversID(id uuid.UUID) *VoiceoverOutputCreate {
	voc.mutation.SetVoiceoversID(id)
	return voc
}

// SetVoiceovers sets the "voiceovers" edge to the Voiceover entity.
func (voc *VoiceoverOutputCreate) SetVoiceovers(v *Voiceover) *VoiceoverOutputCreate {
	return voc.SetVoiceoversID(v.ID)
}

// Mutation returns the VoiceoverOutputMutation object of the builder.
func (voc *VoiceoverOutputCreate) Mutation() *VoiceoverOutputMutation {
	return voc.mutation
}

// Save creates the VoiceoverOutput in the database.
func (voc *VoiceoverOutputCreate) Save(ctx context.Context) (*VoiceoverOutput, error) {
	voc.defaults()
	return withHooks[*VoiceoverOutput, VoiceoverOutputMutation](ctx, voc.sqlSave, voc.mutation, voc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (voc *VoiceoverOutputCreate) SaveX(ctx context.Context) *VoiceoverOutput {
	v, err := voc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (voc *VoiceoverOutputCreate) Exec(ctx context.Context) error {
	_, err := voc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (voc *VoiceoverOutputCreate) ExecX(ctx context.Context) {
	if err := voc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (voc *VoiceoverOutputCreate) defaults() {
	if _, ok := voc.mutation.IsFavorited(); !ok {
		v := voiceoveroutput.DefaultIsFavorited
		voc.mutation.SetIsFavorited(v)
	}
	if _, ok := voc.mutation.GalleryStatus(); !ok {
		v := voiceoveroutput.DefaultGalleryStatus
		voc.mutation.SetGalleryStatus(v)
	}
	if _, ok := voc.mutation.CreatedAt(); !ok {
		v := voiceoveroutput.DefaultCreatedAt()
		voc.mutation.SetCreatedAt(v)
	}
	if _, ok := voc.mutation.UpdatedAt(); !ok {
		v := voiceoveroutput.DefaultUpdatedAt()
		voc.mutation.SetUpdatedAt(v)
	}
	if _, ok := voc.mutation.ID(); !ok {
		v := voiceoveroutput.DefaultID()
		voc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (voc *VoiceoverOutputCreate) check() error {
	if _, ok := voc.mutation.AudioPath(); !ok {
		return &ValidationError{Name: "audio_path", err: errors.New(`ent: missing required field "VoiceoverOutput.audio_path"`)}
	}
	if _, ok := voc.mutation.IsFavorited(); !ok {
		return &ValidationError{Name: "is_favorited", err: errors.New(`ent: missing required field "VoiceoverOutput.is_favorited"`)}
	}
	if _, ok := voc.mutation.AudioDuration(); !ok {
		return &ValidationError{Name: "audio_duration", err: errors.New(`ent: missing required field "VoiceoverOutput.audio_duration"`)}
	}
	if _, ok := voc.mutation.GalleryStatus(); !ok {
		return &ValidationError{Name: "gallery_status", err: errors.New(`ent: missing required field "VoiceoverOutput.gallery_status"`)}
	}
	if v, ok := voc.mutation.GalleryStatus(); ok {
		if err := voiceoveroutput.GalleryStatusValidator(v); err != nil {
			return &ValidationError{Name: "gallery_status", err: fmt.Errorf(`ent: validator failed for field "VoiceoverOutput.gallery_status": %w`, err)}
		}
	}
	if _, ok := voc.mutation.VoiceoverID(); !ok {
		return &ValidationError{Name: "voiceover_id", err: errors.New(`ent: missing required field "VoiceoverOutput.voiceover_id"`)}
	}
	if _, ok := voc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "VoiceoverOutput.created_at"`)}
	}
	if _, ok := voc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "VoiceoverOutput.updated_at"`)}
	}
	if _, ok := voc.mutation.VoiceoversID(); !ok {
		return &ValidationError{Name: "voiceovers", err: errors.New(`ent: missing required edge "VoiceoverOutput.voiceovers"`)}
	}
	return nil
}

func (voc *VoiceoverOutputCreate) sqlSave(ctx context.Context) (*VoiceoverOutput, error) {
	if err := voc.check(); err != nil {
		return nil, err
	}
	_node, _spec := voc.createSpec()
	if err := sqlgraph.CreateNode(ctx, voc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	voc.mutation.id = &_node.ID
	voc.mutation.done = true
	return _node, nil
}

func (voc *VoiceoverOutputCreate) createSpec() (*VoiceoverOutput, *sqlgraph.CreateSpec) {
	var (
		_node = &VoiceoverOutput{config: voc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: voiceoveroutput.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: voiceoveroutput.FieldID,
			},
		}
	)
	if id, ok := voc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := voc.mutation.AudioPath(); ok {
		_spec.SetField(voiceoveroutput.FieldAudioPath, field.TypeString, value)
		_node.AudioPath = value
	}
	if value, ok := voc.mutation.VideoPath(); ok {
		_spec.SetField(voiceoveroutput.FieldVideoPath, field.TypeString, value)
		_node.VideoPath = &value
	}
	if value, ok := voc.mutation.AudioArray(); ok {
		_spec.SetField(voiceoveroutput.FieldAudioArray, field.TypeJSON, value)
		_node.AudioArray = value
	}
	if value, ok := voc.mutation.IsFavorited(); ok {
		_spec.SetField(voiceoveroutput.FieldIsFavorited, field.TypeBool, value)
		_node.IsFavorited = value
	}
	if value, ok := voc.mutation.AudioDuration(); ok {
		_spec.SetField(voiceoveroutput.FieldAudioDuration, field.TypeFloat32, value)
		_node.AudioDuration = value
	}
	if value, ok := voc.mutation.GalleryStatus(); ok {
		_spec.SetField(voiceoveroutput.FieldGalleryStatus, field.TypeEnum, value)
		_node.GalleryStatus = value
	}
	if value, ok := voc.mutation.DeletedAt(); ok {
		_spec.SetField(voiceoveroutput.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if value, ok := voc.mutation.CreatedAt(); ok {
		_spec.SetField(voiceoveroutput.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := voc.mutation.UpdatedAt(); ok {
		_spec.SetField(voiceoveroutput.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := voc.mutation.VoiceoversIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   voiceoveroutput.VoiceoversTable,
			Columns: []string{voiceoveroutput.VoiceoversColumn},
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
		_node.VoiceoverID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// VoiceoverOutputCreateBulk is the builder for creating many VoiceoverOutput entities in bulk.
type VoiceoverOutputCreateBulk struct {
	config
	builders []*VoiceoverOutputCreate
}

// Save creates the VoiceoverOutput entities in the database.
func (vocb *VoiceoverOutputCreateBulk) Save(ctx context.Context) ([]*VoiceoverOutput, error) {
	specs := make([]*sqlgraph.CreateSpec, len(vocb.builders))
	nodes := make([]*VoiceoverOutput, len(vocb.builders))
	mutators := make([]Mutator, len(vocb.builders))
	for i := range vocb.builders {
		func(i int, root context.Context) {
			builder := vocb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*VoiceoverOutputMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, vocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vocb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, vocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (vocb *VoiceoverOutputCreateBulk) SaveX(ctx context.Context) []*VoiceoverOutput {
	v, err := vocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vocb *VoiceoverOutputCreateBulk) Exec(ctx context.Context) error {
	_, err := vocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vocb *VoiceoverOutputCreateBulk) ExecX(ctx context.Context) {
	if err := vocb.Exec(ctx); err != nil {
		panic(err)
	}
}
