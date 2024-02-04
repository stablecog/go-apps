// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	pgvector "github.com/pgvector/pgvector-go"
	"github.com/stablecog/sc-go/database/ent/generationoutput"
	"github.com/stablecog/sc-go/database/ent/generationoutputembed"
)

// GenerationOutputEmbedCreate is the builder for creating a GenerationOutputEmbed entity.
type GenerationOutputEmbedCreate struct {
	config
	mutation *GenerationOutputEmbedMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetPromptEmbedding sets the "prompt_embedding" field.
func (goec *GenerationOutputEmbedCreate) SetPromptEmbedding(pg pgvector.Vector) *GenerationOutputEmbedCreate {
	goec.mutation.SetPromptEmbedding(pg)
	return goec
}

// SetImageEmbedding sets the "image_embedding" field.
func (goec *GenerationOutputEmbedCreate) SetImageEmbedding(pg pgvector.Vector) *GenerationOutputEmbedCreate {
	goec.mutation.SetImageEmbedding(pg)
	return goec
}

// SetOutputID sets the "output_id" field.
func (goec *GenerationOutputEmbedCreate) SetOutputID(u uuid.UUID) *GenerationOutputEmbedCreate {
	goec.mutation.SetOutputID(u)
	return goec
}

// SetCreatedAt sets the "created_at" field.
func (goec *GenerationOutputEmbedCreate) SetCreatedAt(t time.Time) *GenerationOutputEmbedCreate {
	goec.mutation.SetCreatedAt(t)
	return goec
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (goec *GenerationOutputEmbedCreate) SetNillableCreatedAt(t *time.Time) *GenerationOutputEmbedCreate {
	if t != nil {
		goec.SetCreatedAt(*t)
	}
	return goec
}

// SetUpdatedAt sets the "updated_at" field.
func (goec *GenerationOutputEmbedCreate) SetUpdatedAt(t time.Time) *GenerationOutputEmbedCreate {
	goec.mutation.SetUpdatedAt(t)
	return goec
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (goec *GenerationOutputEmbedCreate) SetNillableUpdatedAt(t *time.Time) *GenerationOutputEmbedCreate {
	if t != nil {
		goec.SetUpdatedAt(*t)
	}
	return goec
}

// SetID sets the "id" field.
func (goec *GenerationOutputEmbedCreate) SetID(u uuid.UUID) *GenerationOutputEmbedCreate {
	goec.mutation.SetID(u)
	return goec
}

// SetNillableID sets the "id" field if the given value is not nil.
func (goec *GenerationOutputEmbedCreate) SetNillableID(u *uuid.UUID) *GenerationOutputEmbedCreate {
	if u != nil {
		goec.SetID(*u)
	}
	return goec
}

// SetGenerationOutputsID sets the "generation_outputs" edge to the GenerationOutput entity by ID.
func (goec *GenerationOutputEmbedCreate) SetGenerationOutputsID(id uuid.UUID) *GenerationOutputEmbedCreate {
	goec.mutation.SetGenerationOutputsID(id)
	return goec
}

// SetGenerationOutputs sets the "generation_outputs" edge to the GenerationOutput entity.
func (goec *GenerationOutputEmbedCreate) SetGenerationOutputs(g *GenerationOutput) *GenerationOutputEmbedCreate {
	return goec.SetGenerationOutputsID(g.ID)
}

// Mutation returns the GenerationOutputEmbedMutation object of the builder.
func (goec *GenerationOutputEmbedCreate) Mutation() *GenerationOutputEmbedMutation {
	return goec.mutation
}

// Save creates the GenerationOutputEmbed in the database.
func (goec *GenerationOutputEmbedCreate) Save(ctx context.Context) (*GenerationOutputEmbed, error) {
	goec.defaults()
	return withHooks[*GenerationOutputEmbed, GenerationOutputEmbedMutation](ctx, goec.sqlSave, goec.mutation, goec.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (goec *GenerationOutputEmbedCreate) SaveX(ctx context.Context) *GenerationOutputEmbed {
	v, err := goec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (goec *GenerationOutputEmbedCreate) Exec(ctx context.Context) error {
	_, err := goec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (goec *GenerationOutputEmbedCreate) ExecX(ctx context.Context) {
	if err := goec.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (goec *GenerationOutputEmbedCreate) defaults() {
	if _, ok := goec.mutation.CreatedAt(); !ok {
		v := generationoutputembed.DefaultCreatedAt()
		goec.mutation.SetCreatedAt(v)
	}
	if _, ok := goec.mutation.UpdatedAt(); !ok {
		v := generationoutputembed.DefaultUpdatedAt()
		goec.mutation.SetUpdatedAt(v)
	}
	if _, ok := goec.mutation.ID(); !ok {
		v := generationoutputembed.DefaultID()
		goec.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (goec *GenerationOutputEmbedCreate) check() error {
	if _, ok := goec.mutation.PromptEmbedding(); !ok {
		return &ValidationError{Name: "prompt_embedding", err: errors.New(`ent: missing required field "GenerationOutputEmbed.prompt_embedding"`)}
	}
	if _, ok := goec.mutation.ImageEmbedding(); !ok {
		return &ValidationError{Name: "image_embedding", err: errors.New(`ent: missing required field "GenerationOutputEmbed.image_embedding"`)}
	}
	if _, ok := goec.mutation.OutputID(); !ok {
		return &ValidationError{Name: "output_id", err: errors.New(`ent: missing required field "GenerationOutputEmbed.output_id"`)}
	}
	if _, ok := goec.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "GenerationOutputEmbed.created_at"`)}
	}
	if _, ok := goec.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "GenerationOutputEmbed.updated_at"`)}
	}
	if _, ok := goec.mutation.GenerationOutputsID(); !ok {
		return &ValidationError{Name: "generation_outputs", err: errors.New(`ent: missing required edge "GenerationOutputEmbed.generation_outputs"`)}
	}
	return nil
}

func (goec *GenerationOutputEmbedCreate) sqlSave(ctx context.Context) (*GenerationOutputEmbed, error) {
	if err := goec.check(); err != nil {
		return nil, err
	}
	_node, _spec := goec.createSpec()
	if err := sqlgraph.CreateNode(ctx, goec.driver, _spec); err != nil {
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
	goec.mutation.id = &_node.ID
	goec.mutation.done = true
	return _node, nil
}

func (goec *GenerationOutputEmbedCreate) createSpec() (*GenerationOutputEmbed, *sqlgraph.CreateSpec) {
	var (
		_node = &GenerationOutputEmbed{config: goec.config}
		_spec = &sqlgraph.CreateSpec{
			Table: generationoutputembed.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: generationoutputembed.FieldID,
			},
		}
	)
	_spec.OnConflict = goec.conflict
	if id, ok := goec.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := goec.mutation.PromptEmbedding(); ok {
		_spec.SetField(generationoutputembed.FieldPromptEmbedding, field.TypeOther, value)
		_node.PromptEmbedding = value
	}
	if value, ok := goec.mutation.ImageEmbedding(); ok {
		_spec.SetField(generationoutputembed.FieldImageEmbedding, field.TypeOther, value)
		_node.ImageEmbedding = value
	}
	if value, ok := goec.mutation.CreatedAt(); ok {
		_spec.SetField(generationoutputembed.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := goec.mutation.UpdatedAt(); ok {
		_spec.SetField(generationoutputembed.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := goec.mutation.GenerationOutputsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   generationoutputembed.GenerationOutputsTable,
			Columns: []string{generationoutputembed.GenerationOutputsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: generationoutput.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.OutputID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.GenerationOutputEmbed.Create().
//		SetPromptEmbedding(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.GenerationOutputEmbedUpsert) {
//			SetPromptEmbedding(v+v).
//		}).
//		Exec(ctx)
func (goec *GenerationOutputEmbedCreate) OnConflict(opts ...sql.ConflictOption) *GenerationOutputEmbedUpsertOne {
	goec.conflict = opts
	return &GenerationOutputEmbedUpsertOne{
		create: goec,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.GenerationOutputEmbed.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (goec *GenerationOutputEmbedCreate) OnConflictColumns(columns ...string) *GenerationOutputEmbedUpsertOne {
	goec.conflict = append(goec.conflict, sql.ConflictColumns(columns...))
	return &GenerationOutputEmbedUpsertOne{
		create: goec,
	}
}

type (
	// GenerationOutputEmbedUpsertOne is the builder for "upsert"-ing
	//  one GenerationOutputEmbed node.
	GenerationOutputEmbedUpsertOne struct {
		create *GenerationOutputEmbedCreate
	}

	// GenerationOutputEmbedUpsert is the "OnConflict" setter.
	GenerationOutputEmbedUpsert struct {
		*sql.UpdateSet
	}
)

// SetPromptEmbedding sets the "prompt_embedding" field.
func (u *GenerationOutputEmbedUpsert) SetPromptEmbedding(v pgvector.Vector) *GenerationOutputEmbedUpsert {
	u.Set(generationoutputembed.FieldPromptEmbedding, v)
	return u
}

// UpdatePromptEmbedding sets the "prompt_embedding" field to the value that was provided on create.
func (u *GenerationOutputEmbedUpsert) UpdatePromptEmbedding() *GenerationOutputEmbedUpsert {
	u.SetExcluded(generationoutputembed.FieldPromptEmbedding)
	return u
}

// SetImageEmbedding sets the "image_embedding" field.
func (u *GenerationOutputEmbedUpsert) SetImageEmbedding(v pgvector.Vector) *GenerationOutputEmbedUpsert {
	u.Set(generationoutputembed.FieldImageEmbedding, v)
	return u
}

// UpdateImageEmbedding sets the "image_embedding" field to the value that was provided on create.
func (u *GenerationOutputEmbedUpsert) UpdateImageEmbedding() *GenerationOutputEmbedUpsert {
	u.SetExcluded(generationoutputembed.FieldImageEmbedding)
	return u
}

// SetOutputID sets the "output_id" field.
func (u *GenerationOutputEmbedUpsert) SetOutputID(v uuid.UUID) *GenerationOutputEmbedUpsert {
	u.Set(generationoutputembed.FieldOutputID, v)
	return u
}

// UpdateOutputID sets the "output_id" field to the value that was provided on create.
func (u *GenerationOutputEmbedUpsert) UpdateOutputID() *GenerationOutputEmbedUpsert {
	u.SetExcluded(generationoutputembed.FieldOutputID)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *GenerationOutputEmbedUpsert) SetUpdatedAt(v time.Time) *GenerationOutputEmbedUpsert {
	u.Set(generationoutputembed.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *GenerationOutputEmbedUpsert) UpdateUpdatedAt() *GenerationOutputEmbedUpsert {
	u.SetExcluded(generationoutputembed.FieldUpdatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.GenerationOutputEmbed.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(generationoutputembed.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *GenerationOutputEmbedUpsertOne) UpdateNewValues() *GenerationOutputEmbedUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(generationoutputembed.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(generationoutputembed.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.GenerationOutputEmbed.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *GenerationOutputEmbedUpsertOne) Ignore() *GenerationOutputEmbedUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *GenerationOutputEmbedUpsertOne) DoNothing() *GenerationOutputEmbedUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the GenerationOutputEmbedCreate.OnConflict
// documentation for more info.
func (u *GenerationOutputEmbedUpsertOne) Update(set func(*GenerationOutputEmbedUpsert)) *GenerationOutputEmbedUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&GenerationOutputEmbedUpsert{UpdateSet: update})
	}))
	return u
}

// SetPromptEmbedding sets the "prompt_embedding" field.
func (u *GenerationOutputEmbedUpsertOne) SetPromptEmbedding(v pgvector.Vector) *GenerationOutputEmbedUpsertOne {
	return u.Update(func(s *GenerationOutputEmbedUpsert) {
		s.SetPromptEmbedding(v)
	})
}

// UpdatePromptEmbedding sets the "prompt_embedding" field to the value that was provided on create.
func (u *GenerationOutputEmbedUpsertOne) UpdatePromptEmbedding() *GenerationOutputEmbedUpsertOne {
	return u.Update(func(s *GenerationOutputEmbedUpsert) {
		s.UpdatePromptEmbedding()
	})
}

// SetImageEmbedding sets the "image_embedding" field.
func (u *GenerationOutputEmbedUpsertOne) SetImageEmbedding(v pgvector.Vector) *GenerationOutputEmbedUpsertOne {
	return u.Update(func(s *GenerationOutputEmbedUpsert) {
		s.SetImageEmbedding(v)
	})
}

// UpdateImageEmbedding sets the "image_embedding" field to the value that was provided on create.
func (u *GenerationOutputEmbedUpsertOne) UpdateImageEmbedding() *GenerationOutputEmbedUpsertOne {
	return u.Update(func(s *GenerationOutputEmbedUpsert) {
		s.UpdateImageEmbedding()
	})
}

// SetOutputID sets the "output_id" field.
func (u *GenerationOutputEmbedUpsertOne) SetOutputID(v uuid.UUID) *GenerationOutputEmbedUpsertOne {
	return u.Update(func(s *GenerationOutputEmbedUpsert) {
		s.SetOutputID(v)
	})
}

// UpdateOutputID sets the "output_id" field to the value that was provided on create.
func (u *GenerationOutputEmbedUpsertOne) UpdateOutputID() *GenerationOutputEmbedUpsertOne {
	return u.Update(func(s *GenerationOutputEmbedUpsert) {
		s.UpdateOutputID()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *GenerationOutputEmbedUpsertOne) SetUpdatedAt(v time.Time) *GenerationOutputEmbedUpsertOne {
	return u.Update(func(s *GenerationOutputEmbedUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *GenerationOutputEmbedUpsertOne) UpdateUpdatedAt() *GenerationOutputEmbedUpsertOne {
	return u.Update(func(s *GenerationOutputEmbedUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *GenerationOutputEmbedUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for GenerationOutputEmbedCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *GenerationOutputEmbedUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *GenerationOutputEmbedUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: GenerationOutputEmbedUpsertOne.ID is not supported by MySQL driver. Use GenerationOutputEmbedUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *GenerationOutputEmbedUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// GenerationOutputEmbedCreateBulk is the builder for creating many GenerationOutputEmbed entities in bulk.
type GenerationOutputEmbedCreateBulk struct {
	config
	builders []*GenerationOutputEmbedCreate
	conflict []sql.ConflictOption
}

// Save creates the GenerationOutputEmbed entities in the database.
func (goecb *GenerationOutputEmbedCreateBulk) Save(ctx context.Context) ([]*GenerationOutputEmbed, error) {
	specs := make([]*sqlgraph.CreateSpec, len(goecb.builders))
	nodes := make([]*GenerationOutputEmbed, len(goecb.builders))
	mutators := make([]Mutator, len(goecb.builders))
	for i := range goecb.builders {
		func(i int, root context.Context) {
			builder := goecb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GenerationOutputEmbedMutation)
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
					_, err = mutators[i+1].Mutate(root, goecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = goecb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, goecb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, goecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (goecb *GenerationOutputEmbedCreateBulk) SaveX(ctx context.Context) []*GenerationOutputEmbed {
	v, err := goecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (goecb *GenerationOutputEmbedCreateBulk) Exec(ctx context.Context) error {
	_, err := goecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (goecb *GenerationOutputEmbedCreateBulk) ExecX(ctx context.Context) {
	if err := goecb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.GenerationOutputEmbed.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.GenerationOutputEmbedUpsert) {
//			SetPromptEmbedding(v+v).
//		}).
//		Exec(ctx)
func (goecb *GenerationOutputEmbedCreateBulk) OnConflict(opts ...sql.ConflictOption) *GenerationOutputEmbedUpsertBulk {
	goecb.conflict = opts
	return &GenerationOutputEmbedUpsertBulk{
		create: goecb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.GenerationOutputEmbed.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (goecb *GenerationOutputEmbedCreateBulk) OnConflictColumns(columns ...string) *GenerationOutputEmbedUpsertBulk {
	goecb.conflict = append(goecb.conflict, sql.ConflictColumns(columns...))
	return &GenerationOutputEmbedUpsertBulk{
		create: goecb,
	}
}

// GenerationOutputEmbedUpsertBulk is the builder for "upsert"-ing
// a bulk of GenerationOutputEmbed nodes.
type GenerationOutputEmbedUpsertBulk struct {
	create *GenerationOutputEmbedCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.GenerationOutputEmbed.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(generationoutputembed.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *GenerationOutputEmbedUpsertBulk) UpdateNewValues() *GenerationOutputEmbedUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(generationoutputembed.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(generationoutputembed.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.GenerationOutputEmbed.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *GenerationOutputEmbedUpsertBulk) Ignore() *GenerationOutputEmbedUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *GenerationOutputEmbedUpsertBulk) DoNothing() *GenerationOutputEmbedUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the GenerationOutputEmbedCreateBulk.OnConflict
// documentation for more info.
func (u *GenerationOutputEmbedUpsertBulk) Update(set func(*GenerationOutputEmbedUpsert)) *GenerationOutputEmbedUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&GenerationOutputEmbedUpsert{UpdateSet: update})
	}))
	return u
}

// SetPromptEmbedding sets the "prompt_embedding" field.
func (u *GenerationOutputEmbedUpsertBulk) SetPromptEmbedding(v pgvector.Vector) *GenerationOutputEmbedUpsertBulk {
	return u.Update(func(s *GenerationOutputEmbedUpsert) {
		s.SetPromptEmbedding(v)
	})
}

// UpdatePromptEmbedding sets the "prompt_embedding" field to the value that was provided on create.
func (u *GenerationOutputEmbedUpsertBulk) UpdatePromptEmbedding() *GenerationOutputEmbedUpsertBulk {
	return u.Update(func(s *GenerationOutputEmbedUpsert) {
		s.UpdatePromptEmbedding()
	})
}

// SetImageEmbedding sets the "image_embedding" field.
func (u *GenerationOutputEmbedUpsertBulk) SetImageEmbedding(v pgvector.Vector) *GenerationOutputEmbedUpsertBulk {
	return u.Update(func(s *GenerationOutputEmbedUpsert) {
		s.SetImageEmbedding(v)
	})
}

// UpdateImageEmbedding sets the "image_embedding" field to the value that was provided on create.
func (u *GenerationOutputEmbedUpsertBulk) UpdateImageEmbedding() *GenerationOutputEmbedUpsertBulk {
	return u.Update(func(s *GenerationOutputEmbedUpsert) {
		s.UpdateImageEmbedding()
	})
}

// SetOutputID sets the "output_id" field.
func (u *GenerationOutputEmbedUpsertBulk) SetOutputID(v uuid.UUID) *GenerationOutputEmbedUpsertBulk {
	return u.Update(func(s *GenerationOutputEmbedUpsert) {
		s.SetOutputID(v)
	})
}

// UpdateOutputID sets the "output_id" field to the value that was provided on create.
func (u *GenerationOutputEmbedUpsertBulk) UpdateOutputID() *GenerationOutputEmbedUpsertBulk {
	return u.Update(func(s *GenerationOutputEmbedUpsert) {
		s.UpdateOutputID()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *GenerationOutputEmbedUpsertBulk) SetUpdatedAt(v time.Time) *GenerationOutputEmbedUpsertBulk {
	return u.Update(func(s *GenerationOutputEmbedUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *GenerationOutputEmbedUpsertBulk) UpdateUpdatedAt() *GenerationOutputEmbedUpsertBulk {
	return u.Update(func(s *GenerationOutputEmbedUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *GenerationOutputEmbedUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the GenerationOutputEmbedCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for GenerationOutputEmbedCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *GenerationOutputEmbedUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}