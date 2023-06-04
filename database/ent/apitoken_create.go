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
	"github.com/stablecog/sc-go/database/ent/apitoken"
	"github.com/stablecog/sc-go/database/ent/generation"
	"github.com/stablecog/sc-go/database/ent/upscale"
	"github.com/stablecog/sc-go/database/ent/user"
	"github.com/stablecog/sc-go/database/ent/voiceover"
)

// ApiTokenCreate is the builder for creating a ApiToken entity.
type ApiTokenCreate struct {
	config
	mutation *ApiTokenMutation
	hooks    []Hook
}

// SetHashedToken sets the "hashed_token" field.
func (atc *ApiTokenCreate) SetHashedToken(s string) *ApiTokenCreate {
	atc.mutation.SetHashedToken(s)
	return atc
}

// SetName sets the "name" field.
func (atc *ApiTokenCreate) SetName(s string) *ApiTokenCreate {
	atc.mutation.SetName(s)
	return atc
}

// SetShortString sets the "short_string" field.
func (atc *ApiTokenCreate) SetShortString(s string) *ApiTokenCreate {
	atc.mutation.SetShortString(s)
	return atc
}

// SetIsActive sets the "is_active" field.
func (atc *ApiTokenCreate) SetIsActive(b bool) *ApiTokenCreate {
	atc.mutation.SetIsActive(b)
	return atc
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (atc *ApiTokenCreate) SetNillableIsActive(b *bool) *ApiTokenCreate {
	if b != nil {
		atc.SetIsActive(*b)
	}
	return atc
}

// SetUses sets the "uses" field.
func (atc *ApiTokenCreate) SetUses(i int) *ApiTokenCreate {
	atc.mutation.SetUses(i)
	return atc
}

// SetNillableUses sets the "uses" field if the given value is not nil.
func (atc *ApiTokenCreate) SetNillableUses(i *int) *ApiTokenCreate {
	if i != nil {
		atc.SetUses(*i)
	}
	return atc
}

// SetCreditsSpent sets the "credits_spent" field.
func (atc *ApiTokenCreate) SetCreditsSpent(i int) *ApiTokenCreate {
	atc.mutation.SetCreditsSpent(i)
	return atc
}

// SetNillableCreditsSpent sets the "credits_spent" field if the given value is not nil.
func (atc *ApiTokenCreate) SetNillableCreditsSpent(i *int) *ApiTokenCreate {
	if i != nil {
		atc.SetCreditsSpent(*i)
	}
	return atc
}

// SetUserID sets the "user_id" field.
func (atc *ApiTokenCreate) SetUserID(u uuid.UUID) *ApiTokenCreate {
	atc.mutation.SetUserID(u)
	return atc
}

// SetLastUsedAt sets the "last_used_at" field.
func (atc *ApiTokenCreate) SetLastUsedAt(t time.Time) *ApiTokenCreate {
	atc.mutation.SetLastUsedAt(t)
	return atc
}

// SetNillableLastUsedAt sets the "last_used_at" field if the given value is not nil.
func (atc *ApiTokenCreate) SetNillableLastUsedAt(t *time.Time) *ApiTokenCreate {
	if t != nil {
		atc.SetLastUsedAt(*t)
	}
	return atc
}

// SetCreatedAt sets the "created_at" field.
func (atc *ApiTokenCreate) SetCreatedAt(t time.Time) *ApiTokenCreate {
	atc.mutation.SetCreatedAt(t)
	return atc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (atc *ApiTokenCreate) SetNillableCreatedAt(t *time.Time) *ApiTokenCreate {
	if t != nil {
		atc.SetCreatedAt(*t)
	}
	return atc
}

// SetUpdatedAt sets the "updated_at" field.
func (atc *ApiTokenCreate) SetUpdatedAt(t time.Time) *ApiTokenCreate {
	atc.mutation.SetUpdatedAt(t)
	return atc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (atc *ApiTokenCreate) SetNillableUpdatedAt(t *time.Time) *ApiTokenCreate {
	if t != nil {
		atc.SetUpdatedAt(*t)
	}
	return atc
}

// SetID sets the "id" field.
func (atc *ApiTokenCreate) SetID(u uuid.UUID) *ApiTokenCreate {
	atc.mutation.SetID(u)
	return atc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (atc *ApiTokenCreate) SetNillableID(u *uuid.UUID) *ApiTokenCreate {
	if u != nil {
		atc.SetID(*u)
	}
	return atc
}

// SetUser sets the "user" edge to the User entity.
func (atc *ApiTokenCreate) SetUser(u *User) *ApiTokenCreate {
	return atc.SetUserID(u.ID)
}

// AddGenerationIDs adds the "generations" edge to the Generation entity by IDs.
func (atc *ApiTokenCreate) AddGenerationIDs(ids ...uuid.UUID) *ApiTokenCreate {
	atc.mutation.AddGenerationIDs(ids...)
	return atc
}

// AddGenerations adds the "generations" edges to the Generation entity.
func (atc *ApiTokenCreate) AddGenerations(g ...*Generation) *ApiTokenCreate {
	ids := make([]uuid.UUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return atc.AddGenerationIDs(ids...)
}

// AddUpscaleIDs adds the "upscales" edge to the Upscale entity by IDs.
func (atc *ApiTokenCreate) AddUpscaleIDs(ids ...uuid.UUID) *ApiTokenCreate {
	atc.mutation.AddUpscaleIDs(ids...)
	return atc
}

// AddUpscales adds the "upscales" edges to the Upscale entity.
func (atc *ApiTokenCreate) AddUpscales(u ...*Upscale) *ApiTokenCreate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return atc.AddUpscaleIDs(ids...)
}

// AddVoiceoverIDs adds the "voiceovers" edge to the Voiceover entity by IDs.
func (atc *ApiTokenCreate) AddVoiceoverIDs(ids ...uuid.UUID) *ApiTokenCreate {
	atc.mutation.AddVoiceoverIDs(ids...)
	return atc
}

// AddVoiceovers adds the "voiceovers" edges to the Voiceover entity.
func (atc *ApiTokenCreate) AddVoiceovers(v ...*Voiceover) *ApiTokenCreate {
	ids := make([]uuid.UUID, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return atc.AddVoiceoverIDs(ids...)
}

// Mutation returns the ApiTokenMutation object of the builder.
func (atc *ApiTokenCreate) Mutation() *ApiTokenMutation {
	return atc.mutation
}

// Save creates the ApiToken in the database.
func (atc *ApiTokenCreate) Save(ctx context.Context) (*ApiToken, error) {
	atc.defaults()
	return withHooks[*ApiToken, ApiTokenMutation](ctx, atc.sqlSave, atc.mutation, atc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (atc *ApiTokenCreate) SaveX(ctx context.Context) *ApiToken {
	v, err := atc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (atc *ApiTokenCreate) Exec(ctx context.Context) error {
	_, err := atc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atc *ApiTokenCreate) ExecX(ctx context.Context) {
	if err := atc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (atc *ApiTokenCreate) defaults() {
	if _, ok := atc.mutation.IsActive(); !ok {
		v := apitoken.DefaultIsActive
		atc.mutation.SetIsActive(v)
	}
	if _, ok := atc.mutation.Uses(); !ok {
		v := apitoken.DefaultUses
		atc.mutation.SetUses(v)
	}
	if _, ok := atc.mutation.CreditsSpent(); !ok {
		v := apitoken.DefaultCreditsSpent
		atc.mutation.SetCreditsSpent(v)
	}
	if _, ok := atc.mutation.CreatedAt(); !ok {
		v := apitoken.DefaultCreatedAt()
		atc.mutation.SetCreatedAt(v)
	}
	if _, ok := atc.mutation.UpdatedAt(); !ok {
		v := apitoken.DefaultUpdatedAt()
		atc.mutation.SetUpdatedAt(v)
	}
	if _, ok := atc.mutation.ID(); !ok {
		v := apitoken.DefaultID()
		atc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (atc *ApiTokenCreate) check() error {
	if _, ok := atc.mutation.HashedToken(); !ok {
		return &ValidationError{Name: "hashed_token", err: errors.New(`ent: missing required field "ApiToken.hashed_token"`)}
	}
	if _, ok := atc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "ApiToken.name"`)}
	}
	if _, ok := atc.mutation.ShortString(); !ok {
		return &ValidationError{Name: "short_string", err: errors.New(`ent: missing required field "ApiToken.short_string"`)}
	}
	if _, ok := atc.mutation.IsActive(); !ok {
		return &ValidationError{Name: "is_active", err: errors.New(`ent: missing required field "ApiToken.is_active"`)}
	}
	if _, ok := atc.mutation.Uses(); !ok {
		return &ValidationError{Name: "uses", err: errors.New(`ent: missing required field "ApiToken.uses"`)}
	}
	if _, ok := atc.mutation.CreditsSpent(); !ok {
		return &ValidationError{Name: "credits_spent", err: errors.New(`ent: missing required field "ApiToken.credits_spent"`)}
	}
	if _, ok := atc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "ApiToken.user_id"`)}
	}
	if _, ok := atc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "ApiToken.created_at"`)}
	}
	if _, ok := atc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "ApiToken.updated_at"`)}
	}
	if _, ok := atc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "ApiToken.user"`)}
	}
	return nil
}

func (atc *ApiTokenCreate) sqlSave(ctx context.Context) (*ApiToken, error) {
	if err := atc.check(); err != nil {
		return nil, err
	}
	_node, _spec := atc.createSpec()
	if err := sqlgraph.CreateNode(ctx, atc.driver, _spec); err != nil {
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
	atc.mutation.id = &_node.ID
	atc.mutation.done = true
	return _node, nil
}

func (atc *ApiTokenCreate) createSpec() (*ApiToken, *sqlgraph.CreateSpec) {
	var (
		_node = &ApiToken{config: atc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: apitoken.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: apitoken.FieldID,
			},
		}
	)
	if id, ok := atc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := atc.mutation.HashedToken(); ok {
		_spec.SetField(apitoken.FieldHashedToken, field.TypeString, value)
		_node.HashedToken = value
	}
	if value, ok := atc.mutation.Name(); ok {
		_spec.SetField(apitoken.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := atc.mutation.ShortString(); ok {
		_spec.SetField(apitoken.FieldShortString, field.TypeString, value)
		_node.ShortString = value
	}
	if value, ok := atc.mutation.IsActive(); ok {
		_spec.SetField(apitoken.FieldIsActive, field.TypeBool, value)
		_node.IsActive = value
	}
	if value, ok := atc.mutation.Uses(); ok {
		_spec.SetField(apitoken.FieldUses, field.TypeInt, value)
		_node.Uses = value
	}
	if value, ok := atc.mutation.CreditsSpent(); ok {
		_spec.SetField(apitoken.FieldCreditsSpent, field.TypeInt, value)
		_node.CreditsSpent = value
	}
	if value, ok := atc.mutation.LastUsedAt(); ok {
		_spec.SetField(apitoken.FieldLastUsedAt, field.TypeTime, value)
		_node.LastUsedAt = &value
	}
	if value, ok := atc.mutation.CreatedAt(); ok {
		_spec.SetField(apitoken.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := atc.mutation.UpdatedAt(); ok {
		_spec.SetField(apitoken.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := atc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   apitoken.UserTable,
			Columns: []string{apitoken.UserColumn},
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
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := atc.mutation.GenerationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   apitoken.GenerationsTable,
			Columns: []string{apitoken.GenerationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: generation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := atc.mutation.UpscalesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   apitoken.UpscalesTable,
			Columns: []string{apitoken.UpscalesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: upscale.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := atc.mutation.VoiceoversIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   apitoken.VoiceoversTable,
			Columns: []string{apitoken.VoiceoversColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ApiTokenCreateBulk is the builder for creating many ApiToken entities in bulk.
type ApiTokenCreateBulk struct {
	config
	builders []*ApiTokenCreate
}

// Save creates the ApiToken entities in the database.
func (atcb *ApiTokenCreateBulk) Save(ctx context.Context) ([]*ApiToken, error) {
	specs := make([]*sqlgraph.CreateSpec, len(atcb.builders))
	nodes := make([]*ApiToken, len(atcb.builders))
	mutators := make([]Mutator, len(atcb.builders))
	for i := range atcb.builders {
		func(i int, root context.Context) {
			builder := atcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ApiTokenMutation)
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
					_, err = mutators[i+1].Mutate(root, atcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, atcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, atcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (atcb *ApiTokenCreateBulk) SaveX(ctx context.Context) []*ApiToken {
	v, err := atcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (atcb *ApiTokenCreateBulk) Exec(ctx context.Context) error {
	_, err := atcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atcb *ApiTokenCreateBulk) ExecX(ctx context.Context) {
	if err := atcb.Exec(ctx); err != nil {
		panic(err)
	}
}
