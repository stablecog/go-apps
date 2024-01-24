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
	"github.com/stablecog/sc-go/database/ent/usernameblacklist"
)

// UsernameBlacklistCreate is the builder for creating a UsernameBlacklist entity.
type UsernameBlacklistCreate struct {
	config
	mutation *UsernameBlacklistMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetUsername sets the "username" field.
func (ubc *UsernameBlacklistCreate) SetUsername(s string) *UsernameBlacklistCreate {
	ubc.mutation.SetUsername(s)
	return ubc
}

// SetCreatedAt sets the "created_at" field.
func (ubc *UsernameBlacklistCreate) SetCreatedAt(t time.Time) *UsernameBlacklistCreate {
	ubc.mutation.SetCreatedAt(t)
	return ubc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ubc *UsernameBlacklistCreate) SetNillableCreatedAt(t *time.Time) *UsernameBlacklistCreate {
	if t != nil {
		ubc.SetCreatedAt(*t)
	}
	return ubc
}

// SetUpdatedAt sets the "updated_at" field.
func (ubc *UsernameBlacklistCreate) SetUpdatedAt(t time.Time) *UsernameBlacklistCreate {
	ubc.mutation.SetUpdatedAt(t)
	return ubc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ubc *UsernameBlacklistCreate) SetNillableUpdatedAt(t *time.Time) *UsernameBlacklistCreate {
	if t != nil {
		ubc.SetUpdatedAt(*t)
	}
	return ubc
}

// SetID sets the "id" field.
func (ubc *UsernameBlacklistCreate) SetID(u uuid.UUID) *UsernameBlacklistCreate {
	ubc.mutation.SetID(u)
	return ubc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ubc *UsernameBlacklistCreate) SetNillableID(u *uuid.UUID) *UsernameBlacklistCreate {
	if u != nil {
		ubc.SetID(*u)
	}
	return ubc
}

// Mutation returns the UsernameBlacklistMutation object of the builder.
func (ubc *UsernameBlacklistCreate) Mutation() *UsernameBlacklistMutation {
	return ubc.mutation
}

// Save creates the UsernameBlacklist in the database.
func (ubc *UsernameBlacklistCreate) Save(ctx context.Context) (*UsernameBlacklist, error) {
	ubc.defaults()
	return withHooks[*UsernameBlacklist, UsernameBlacklistMutation](ctx, ubc.sqlSave, ubc.mutation, ubc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ubc *UsernameBlacklistCreate) SaveX(ctx context.Context) *UsernameBlacklist {
	v, err := ubc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ubc *UsernameBlacklistCreate) Exec(ctx context.Context) error {
	_, err := ubc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ubc *UsernameBlacklistCreate) ExecX(ctx context.Context) {
	if err := ubc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ubc *UsernameBlacklistCreate) defaults() {
	if _, ok := ubc.mutation.CreatedAt(); !ok {
		v := usernameblacklist.DefaultCreatedAt()
		ubc.mutation.SetCreatedAt(v)
	}
	if _, ok := ubc.mutation.UpdatedAt(); !ok {
		v := usernameblacklist.DefaultUpdatedAt()
		ubc.mutation.SetUpdatedAt(v)
	}
	if _, ok := ubc.mutation.ID(); !ok {
		v := usernameblacklist.DefaultID()
		ubc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ubc *UsernameBlacklistCreate) check() error {
	if _, ok := ubc.mutation.Username(); !ok {
		return &ValidationError{Name: "username", err: errors.New(`ent: missing required field "UsernameBlacklist.username"`)}
	}
	if _, ok := ubc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "UsernameBlacklist.created_at"`)}
	}
	if _, ok := ubc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "UsernameBlacklist.updated_at"`)}
	}
	return nil
}

func (ubc *UsernameBlacklistCreate) sqlSave(ctx context.Context) (*UsernameBlacklist, error) {
	if err := ubc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ubc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ubc.driver, _spec); err != nil {
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
	ubc.mutation.id = &_node.ID
	ubc.mutation.done = true
	return _node, nil
}

func (ubc *UsernameBlacklistCreate) createSpec() (*UsernameBlacklist, *sqlgraph.CreateSpec) {
	var (
		_node = &UsernameBlacklist{config: ubc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: usernameblacklist.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: usernameblacklist.FieldID,
			},
		}
	)
	_spec.OnConflict = ubc.conflict
	if id, ok := ubc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ubc.mutation.Username(); ok {
		_spec.SetField(usernameblacklist.FieldUsername, field.TypeString, value)
		_node.Username = value
	}
	if value, ok := ubc.mutation.CreatedAt(); ok {
		_spec.SetField(usernameblacklist.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ubc.mutation.UpdatedAt(); ok {
		_spec.SetField(usernameblacklist.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.UsernameBlacklist.Create().
//		SetUsername(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UsernameBlacklistUpsert) {
//			SetUsername(v+v).
//		}).
//		Exec(ctx)
func (ubc *UsernameBlacklistCreate) OnConflict(opts ...sql.ConflictOption) *UsernameBlacklistUpsertOne {
	ubc.conflict = opts
	return &UsernameBlacklistUpsertOne{
		create: ubc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.UsernameBlacklist.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ubc *UsernameBlacklistCreate) OnConflictColumns(columns ...string) *UsernameBlacklistUpsertOne {
	ubc.conflict = append(ubc.conflict, sql.ConflictColumns(columns...))
	return &UsernameBlacklistUpsertOne{
		create: ubc,
	}
}

type (
	// UsernameBlacklistUpsertOne is the builder for "upsert"-ing
	//  one UsernameBlacklist node.
	UsernameBlacklistUpsertOne struct {
		create *UsernameBlacklistCreate
	}

	// UsernameBlacklistUpsert is the "OnConflict" setter.
	UsernameBlacklistUpsert struct {
		*sql.UpdateSet
	}
)

// SetUsername sets the "username" field.
func (u *UsernameBlacklistUpsert) SetUsername(v string) *UsernameBlacklistUpsert {
	u.Set(usernameblacklist.FieldUsername, v)
	return u
}

// UpdateUsername sets the "username" field to the value that was provided on create.
func (u *UsernameBlacklistUpsert) UpdateUsername() *UsernameBlacklistUpsert {
	u.SetExcluded(usernameblacklist.FieldUsername)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *UsernameBlacklistUpsert) SetUpdatedAt(v time.Time) *UsernameBlacklistUpsert {
	u.Set(usernameblacklist.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *UsernameBlacklistUpsert) UpdateUpdatedAt() *UsernameBlacklistUpsert {
	u.SetExcluded(usernameblacklist.FieldUpdatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.UsernameBlacklist.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(usernameblacklist.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *UsernameBlacklistUpsertOne) UpdateNewValues() *UsernameBlacklistUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(usernameblacklist.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(usernameblacklist.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.UsernameBlacklist.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *UsernameBlacklistUpsertOne) Ignore() *UsernameBlacklistUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UsernameBlacklistUpsertOne) DoNothing() *UsernameBlacklistUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UsernameBlacklistCreate.OnConflict
// documentation for more info.
func (u *UsernameBlacklistUpsertOne) Update(set func(*UsernameBlacklistUpsert)) *UsernameBlacklistUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UsernameBlacklistUpsert{UpdateSet: update})
	}))
	return u
}

// SetUsername sets the "username" field.
func (u *UsernameBlacklistUpsertOne) SetUsername(v string) *UsernameBlacklistUpsertOne {
	return u.Update(func(s *UsernameBlacklistUpsert) {
		s.SetUsername(v)
	})
}

// UpdateUsername sets the "username" field to the value that was provided on create.
func (u *UsernameBlacklistUpsertOne) UpdateUsername() *UsernameBlacklistUpsertOne {
	return u.Update(func(s *UsernameBlacklistUpsert) {
		s.UpdateUsername()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *UsernameBlacklistUpsertOne) SetUpdatedAt(v time.Time) *UsernameBlacklistUpsertOne {
	return u.Update(func(s *UsernameBlacklistUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *UsernameBlacklistUpsertOne) UpdateUpdatedAt() *UsernameBlacklistUpsertOne {
	return u.Update(func(s *UsernameBlacklistUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *UsernameBlacklistUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UsernameBlacklistCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UsernameBlacklistUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *UsernameBlacklistUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: UsernameBlacklistUpsertOne.ID is not supported by MySQL driver. Use UsernameBlacklistUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *UsernameBlacklistUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// UsernameBlacklistCreateBulk is the builder for creating many UsernameBlacklist entities in bulk.
type UsernameBlacklistCreateBulk struct {
	config
	builders []*UsernameBlacklistCreate
	conflict []sql.ConflictOption
}

// Save creates the UsernameBlacklist entities in the database.
func (ubcb *UsernameBlacklistCreateBulk) Save(ctx context.Context) ([]*UsernameBlacklist, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ubcb.builders))
	nodes := make([]*UsernameBlacklist, len(ubcb.builders))
	mutators := make([]Mutator, len(ubcb.builders))
	for i := range ubcb.builders {
		func(i int, root context.Context) {
			builder := ubcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UsernameBlacklistMutation)
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
					_, err = mutators[i+1].Mutate(root, ubcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ubcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ubcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ubcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ubcb *UsernameBlacklistCreateBulk) SaveX(ctx context.Context) []*UsernameBlacklist {
	v, err := ubcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ubcb *UsernameBlacklistCreateBulk) Exec(ctx context.Context) error {
	_, err := ubcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ubcb *UsernameBlacklistCreateBulk) ExecX(ctx context.Context) {
	if err := ubcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.UsernameBlacklist.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UsernameBlacklistUpsert) {
//			SetUsername(v+v).
//		}).
//		Exec(ctx)
func (ubcb *UsernameBlacklistCreateBulk) OnConflict(opts ...sql.ConflictOption) *UsernameBlacklistUpsertBulk {
	ubcb.conflict = opts
	return &UsernameBlacklistUpsertBulk{
		create: ubcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.UsernameBlacklist.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ubcb *UsernameBlacklistCreateBulk) OnConflictColumns(columns ...string) *UsernameBlacklistUpsertBulk {
	ubcb.conflict = append(ubcb.conflict, sql.ConflictColumns(columns...))
	return &UsernameBlacklistUpsertBulk{
		create: ubcb,
	}
}

// UsernameBlacklistUpsertBulk is the builder for "upsert"-ing
// a bulk of UsernameBlacklist nodes.
type UsernameBlacklistUpsertBulk struct {
	create *UsernameBlacklistCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.UsernameBlacklist.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(usernameblacklist.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *UsernameBlacklistUpsertBulk) UpdateNewValues() *UsernameBlacklistUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(usernameblacklist.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(usernameblacklist.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.UsernameBlacklist.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *UsernameBlacklistUpsertBulk) Ignore() *UsernameBlacklistUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UsernameBlacklistUpsertBulk) DoNothing() *UsernameBlacklistUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UsernameBlacklistCreateBulk.OnConflict
// documentation for more info.
func (u *UsernameBlacklistUpsertBulk) Update(set func(*UsernameBlacklistUpsert)) *UsernameBlacklistUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UsernameBlacklistUpsert{UpdateSet: update})
	}))
	return u
}

// SetUsername sets the "username" field.
func (u *UsernameBlacklistUpsertBulk) SetUsername(v string) *UsernameBlacklistUpsertBulk {
	return u.Update(func(s *UsernameBlacklistUpsert) {
		s.SetUsername(v)
	})
}

// UpdateUsername sets the "username" field to the value that was provided on create.
func (u *UsernameBlacklistUpsertBulk) UpdateUsername() *UsernameBlacklistUpsertBulk {
	return u.Update(func(s *UsernameBlacklistUpsert) {
		s.UpdateUsername()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *UsernameBlacklistUpsertBulk) SetUpdatedAt(v time.Time) *UsernameBlacklistUpsertBulk {
	return u.Update(func(s *UsernameBlacklistUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *UsernameBlacklistUpsertBulk) UpdateUpdatedAt() *UsernameBlacklistUpsertBulk {
	return u.Update(func(s *UsernameBlacklistUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *UsernameBlacklistUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the UsernameBlacklistCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UsernameBlacklistCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UsernameBlacklistUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}