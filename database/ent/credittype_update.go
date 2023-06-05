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
	"github.com/stablecog/sc-go/database/ent/credit"
	"github.com/stablecog/sc-go/database/ent/credittype"
	"github.com/stablecog/sc-go/database/ent/predicate"
)

// CreditTypeUpdate is the builder for updating CreditType entities.
type CreditTypeUpdate struct {
	config
	hooks     []Hook
	mutation  *CreditTypeMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the CreditTypeUpdate builder.
func (ctu *CreditTypeUpdate) Where(ps ...predicate.CreditType) *CreditTypeUpdate {
	ctu.mutation.Where(ps...)
	return ctu
}

// SetName sets the "name" field.
func (ctu *CreditTypeUpdate) SetName(s string) *CreditTypeUpdate {
	ctu.mutation.SetName(s)
	return ctu
}

// SetDescription sets the "description" field.
func (ctu *CreditTypeUpdate) SetDescription(s string) *CreditTypeUpdate {
	ctu.mutation.SetDescription(s)
	return ctu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ctu *CreditTypeUpdate) SetNillableDescription(s *string) *CreditTypeUpdate {
	if s != nil {
		ctu.SetDescription(*s)
	}
	return ctu
}

// ClearDescription clears the value of the "description" field.
func (ctu *CreditTypeUpdate) ClearDescription() *CreditTypeUpdate {
	ctu.mutation.ClearDescription()
	return ctu
}

// SetAmount sets the "amount" field.
func (ctu *CreditTypeUpdate) SetAmount(i int32) *CreditTypeUpdate {
	ctu.mutation.ResetAmount()
	ctu.mutation.SetAmount(i)
	return ctu
}

// AddAmount adds i to the "amount" field.
func (ctu *CreditTypeUpdate) AddAmount(i int32) *CreditTypeUpdate {
	ctu.mutation.AddAmount(i)
	return ctu
}

// SetStripeProductID sets the "stripe_product_id" field.
func (ctu *CreditTypeUpdate) SetStripeProductID(s string) *CreditTypeUpdate {
	ctu.mutation.SetStripeProductID(s)
	return ctu
}

// SetNillableStripeProductID sets the "stripe_product_id" field if the given value is not nil.
func (ctu *CreditTypeUpdate) SetNillableStripeProductID(s *string) *CreditTypeUpdate {
	if s != nil {
		ctu.SetStripeProductID(*s)
	}
	return ctu
}

// ClearStripeProductID clears the value of the "stripe_product_id" field.
func (ctu *CreditTypeUpdate) ClearStripeProductID() *CreditTypeUpdate {
	ctu.mutation.ClearStripeProductID()
	return ctu
}

// SetType sets the "type" field.
func (ctu *CreditTypeUpdate) SetType(c credittype.Type) *CreditTypeUpdate {
	ctu.mutation.SetType(c)
	return ctu
}

// SetUpdatedAt sets the "updated_at" field.
func (ctu *CreditTypeUpdate) SetUpdatedAt(t time.Time) *CreditTypeUpdate {
	ctu.mutation.SetUpdatedAt(t)
	return ctu
}

// AddCreditIDs adds the "credits" edge to the Credit entity by IDs.
func (ctu *CreditTypeUpdate) AddCreditIDs(ids ...uuid.UUID) *CreditTypeUpdate {
	ctu.mutation.AddCreditIDs(ids...)
	return ctu
}

// AddCredits adds the "credits" edges to the Credit entity.
func (ctu *CreditTypeUpdate) AddCredits(c ...*Credit) *CreditTypeUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ctu.AddCreditIDs(ids...)
}

// Mutation returns the CreditTypeMutation object of the builder.
func (ctu *CreditTypeUpdate) Mutation() *CreditTypeMutation {
	return ctu.mutation
}

// ClearCredits clears all "credits" edges to the Credit entity.
func (ctu *CreditTypeUpdate) ClearCredits() *CreditTypeUpdate {
	ctu.mutation.ClearCredits()
	return ctu
}

// RemoveCreditIDs removes the "credits" edge to Credit entities by IDs.
func (ctu *CreditTypeUpdate) RemoveCreditIDs(ids ...uuid.UUID) *CreditTypeUpdate {
	ctu.mutation.RemoveCreditIDs(ids...)
	return ctu
}

// RemoveCredits removes "credits" edges to Credit entities.
func (ctu *CreditTypeUpdate) RemoveCredits(c ...*Credit) *CreditTypeUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ctu.RemoveCreditIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ctu *CreditTypeUpdate) Save(ctx context.Context) (int, error) {
	ctu.defaults()
	return withHooks[int, CreditTypeMutation](ctx, ctu.sqlSave, ctu.mutation, ctu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ctu *CreditTypeUpdate) SaveX(ctx context.Context) int {
	affected, err := ctu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ctu *CreditTypeUpdate) Exec(ctx context.Context) error {
	_, err := ctu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctu *CreditTypeUpdate) ExecX(ctx context.Context) {
	if err := ctu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ctu *CreditTypeUpdate) defaults() {
	if _, ok := ctu.mutation.UpdatedAt(); !ok {
		v := credittype.UpdateDefaultUpdatedAt()
		ctu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ctu *CreditTypeUpdate) check() error {
	if v, ok := ctu.mutation.GetType(); ok {
		if err := credittype.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "CreditType.type": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ctu *CreditTypeUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CreditTypeUpdate {
	ctu.modifiers = append(ctu.modifiers, modifiers...)
	return ctu
}

func (ctu *CreditTypeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ctu.check(); err != nil {
		return n, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   credittype.Table,
			Columns: credittype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: credittype.FieldID,
			},
		},
	}
	if ps := ctu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ctu.mutation.Name(); ok {
		_spec.SetField(credittype.FieldName, field.TypeString, value)
	}
	if value, ok := ctu.mutation.Description(); ok {
		_spec.SetField(credittype.FieldDescription, field.TypeString, value)
	}
	if ctu.mutation.DescriptionCleared() {
		_spec.ClearField(credittype.FieldDescription, field.TypeString)
	}
	if value, ok := ctu.mutation.Amount(); ok {
		_spec.SetField(credittype.FieldAmount, field.TypeInt32, value)
	}
	if value, ok := ctu.mutation.AddedAmount(); ok {
		_spec.AddField(credittype.FieldAmount, field.TypeInt32, value)
	}
	if value, ok := ctu.mutation.StripeProductID(); ok {
		_spec.SetField(credittype.FieldStripeProductID, field.TypeString, value)
	}
	if ctu.mutation.StripeProductIDCleared() {
		_spec.ClearField(credittype.FieldStripeProductID, field.TypeString)
	}
	if value, ok := ctu.mutation.GetType(); ok {
		_spec.SetField(credittype.FieldType, field.TypeEnum, value)
	}
	if value, ok := ctu.mutation.UpdatedAt(); ok {
		_spec.SetField(credittype.FieldUpdatedAt, field.TypeTime, value)
	}
	if ctu.mutation.CreditsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   credittype.CreditsTable,
			Columns: []string{credittype.CreditsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: credit.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ctu.mutation.RemovedCreditsIDs(); len(nodes) > 0 && !ctu.mutation.CreditsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   credittype.CreditsTable,
			Columns: []string{credittype.CreditsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: credit.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ctu.mutation.CreditsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   credittype.CreditsTable,
			Columns: []string{credittype.CreditsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: credit.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(ctu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, ctu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{credittype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ctu.mutation.done = true
	return n, nil
}

// CreditTypeUpdateOne is the builder for updating a single CreditType entity.
type CreditTypeUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *CreditTypeMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetName sets the "name" field.
func (ctuo *CreditTypeUpdateOne) SetName(s string) *CreditTypeUpdateOne {
	ctuo.mutation.SetName(s)
	return ctuo
}

// SetDescription sets the "description" field.
func (ctuo *CreditTypeUpdateOne) SetDescription(s string) *CreditTypeUpdateOne {
	ctuo.mutation.SetDescription(s)
	return ctuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ctuo *CreditTypeUpdateOne) SetNillableDescription(s *string) *CreditTypeUpdateOne {
	if s != nil {
		ctuo.SetDescription(*s)
	}
	return ctuo
}

// ClearDescription clears the value of the "description" field.
func (ctuo *CreditTypeUpdateOne) ClearDescription() *CreditTypeUpdateOne {
	ctuo.mutation.ClearDescription()
	return ctuo
}

// SetAmount sets the "amount" field.
func (ctuo *CreditTypeUpdateOne) SetAmount(i int32) *CreditTypeUpdateOne {
	ctuo.mutation.ResetAmount()
	ctuo.mutation.SetAmount(i)
	return ctuo
}

// AddAmount adds i to the "amount" field.
func (ctuo *CreditTypeUpdateOne) AddAmount(i int32) *CreditTypeUpdateOne {
	ctuo.mutation.AddAmount(i)
	return ctuo
}

// SetStripeProductID sets the "stripe_product_id" field.
func (ctuo *CreditTypeUpdateOne) SetStripeProductID(s string) *CreditTypeUpdateOne {
	ctuo.mutation.SetStripeProductID(s)
	return ctuo
}

// SetNillableStripeProductID sets the "stripe_product_id" field if the given value is not nil.
func (ctuo *CreditTypeUpdateOne) SetNillableStripeProductID(s *string) *CreditTypeUpdateOne {
	if s != nil {
		ctuo.SetStripeProductID(*s)
	}
	return ctuo
}

// ClearStripeProductID clears the value of the "stripe_product_id" field.
func (ctuo *CreditTypeUpdateOne) ClearStripeProductID() *CreditTypeUpdateOne {
	ctuo.mutation.ClearStripeProductID()
	return ctuo
}

// SetType sets the "type" field.
func (ctuo *CreditTypeUpdateOne) SetType(c credittype.Type) *CreditTypeUpdateOne {
	ctuo.mutation.SetType(c)
	return ctuo
}

// SetUpdatedAt sets the "updated_at" field.
func (ctuo *CreditTypeUpdateOne) SetUpdatedAt(t time.Time) *CreditTypeUpdateOne {
	ctuo.mutation.SetUpdatedAt(t)
	return ctuo
}

// AddCreditIDs adds the "credits" edge to the Credit entity by IDs.
func (ctuo *CreditTypeUpdateOne) AddCreditIDs(ids ...uuid.UUID) *CreditTypeUpdateOne {
	ctuo.mutation.AddCreditIDs(ids...)
	return ctuo
}

// AddCredits adds the "credits" edges to the Credit entity.
func (ctuo *CreditTypeUpdateOne) AddCredits(c ...*Credit) *CreditTypeUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ctuo.AddCreditIDs(ids...)
}

// Mutation returns the CreditTypeMutation object of the builder.
func (ctuo *CreditTypeUpdateOne) Mutation() *CreditTypeMutation {
	return ctuo.mutation
}

// ClearCredits clears all "credits" edges to the Credit entity.
func (ctuo *CreditTypeUpdateOne) ClearCredits() *CreditTypeUpdateOne {
	ctuo.mutation.ClearCredits()
	return ctuo
}

// RemoveCreditIDs removes the "credits" edge to Credit entities by IDs.
func (ctuo *CreditTypeUpdateOne) RemoveCreditIDs(ids ...uuid.UUID) *CreditTypeUpdateOne {
	ctuo.mutation.RemoveCreditIDs(ids...)
	return ctuo
}

// RemoveCredits removes "credits" edges to Credit entities.
func (ctuo *CreditTypeUpdateOne) RemoveCredits(c ...*Credit) *CreditTypeUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ctuo.RemoveCreditIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ctuo *CreditTypeUpdateOne) Select(field string, fields ...string) *CreditTypeUpdateOne {
	ctuo.fields = append([]string{field}, fields...)
	return ctuo
}

// Save executes the query and returns the updated CreditType entity.
func (ctuo *CreditTypeUpdateOne) Save(ctx context.Context) (*CreditType, error) {
	ctuo.defaults()
	return withHooks[*CreditType, CreditTypeMutation](ctx, ctuo.sqlSave, ctuo.mutation, ctuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ctuo *CreditTypeUpdateOne) SaveX(ctx context.Context) *CreditType {
	node, err := ctuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ctuo *CreditTypeUpdateOne) Exec(ctx context.Context) error {
	_, err := ctuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctuo *CreditTypeUpdateOne) ExecX(ctx context.Context) {
	if err := ctuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ctuo *CreditTypeUpdateOne) defaults() {
	if _, ok := ctuo.mutation.UpdatedAt(); !ok {
		v := credittype.UpdateDefaultUpdatedAt()
		ctuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ctuo *CreditTypeUpdateOne) check() error {
	if v, ok := ctuo.mutation.GetType(); ok {
		if err := credittype.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "CreditType.type": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ctuo *CreditTypeUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CreditTypeUpdateOne {
	ctuo.modifiers = append(ctuo.modifiers, modifiers...)
	return ctuo
}

func (ctuo *CreditTypeUpdateOne) sqlSave(ctx context.Context) (_node *CreditType, err error) {
	if err := ctuo.check(); err != nil {
		return _node, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   credittype.Table,
			Columns: credittype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: credittype.FieldID,
			},
		},
	}
	id, ok := ctuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CreditType.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ctuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, credittype.FieldID)
		for _, f := range fields {
			if !credittype.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != credittype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ctuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ctuo.mutation.Name(); ok {
		_spec.SetField(credittype.FieldName, field.TypeString, value)
	}
	if value, ok := ctuo.mutation.Description(); ok {
		_spec.SetField(credittype.FieldDescription, field.TypeString, value)
	}
	if ctuo.mutation.DescriptionCleared() {
		_spec.ClearField(credittype.FieldDescription, field.TypeString)
	}
	if value, ok := ctuo.mutation.Amount(); ok {
		_spec.SetField(credittype.FieldAmount, field.TypeInt32, value)
	}
	if value, ok := ctuo.mutation.AddedAmount(); ok {
		_spec.AddField(credittype.FieldAmount, field.TypeInt32, value)
	}
	if value, ok := ctuo.mutation.StripeProductID(); ok {
		_spec.SetField(credittype.FieldStripeProductID, field.TypeString, value)
	}
	if ctuo.mutation.StripeProductIDCleared() {
		_spec.ClearField(credittype.FieldStripeProductID, field.TypeString)
	}
	if value, ok := ctuo.mutation.GetType(); ok {
		_spec.SetField(credittype.FieldType, field.TypeEnum, value)
	}
	if value, ok := ctuo.mutation.UpdatedAt(); ok {
		_spec.SetField(credittype.FieldUpdatedAt, field.TypeTime, value)
	}
	if ctuo.mutation.CreditsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   credittype.CreditsTable,
			Columns: []string{credittype.CreditsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: credit.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ctuo.mutation.RemovedCreditsIDs(); len(nodes) > 0 && !ctuo.mutation.CreditsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   credittype.CreditsTable,
			Columns: []string{credittype.CreditsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: credit.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ctuo.mutation.CreditsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   credittype.CreditsTable,
			Columns: []string{credittype.CreditsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: credit.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(ctuo.modifiers...)
	_node = &CreditType{config: ctuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ctuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{credittype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ctuo.mutation.done = true
	return _node, nil
}
