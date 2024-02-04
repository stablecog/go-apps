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
	pgvector "github.com/pgvector/pgvector-go"
	"github.com/stablecog/sc-go/database/ent/generationoutput"
	"github.com/stablecog/sc-go/database/ent/generationoutputembed"
	"github.com/stablecog/sc-go/database/ent/predicate"
)

// GenerationOutputEmbedUpdate is the builder for updating GenerationOutputEmbed entities.
type GenerationOutputEmbedUpdate struct {
	config
	hooks     []Hook
	mutation  *GenerationOutputEmbedMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the GenerationOutputEmbedUpdate builder.
func (goeu *GenerationOutputEmbedUpdate) Where(ps ...predicate.GenerationOutputEmbed) *GenerationOutputEmbedUpdate {
	goeu.mutation.Where(ps...)
	return goeu
}

// SetPromptEmbedding sets the "prompt_embedding" field.
func (goeu *GenerationOutputEmbedUpdate) SetPromptEmbedding(pg pgvector.Vector) *GenerationOutputEmbedUpdate {
	goeu.mutation.SetPromptEmbedding(pg)
	return goeu
}

// SetImageEmbedding sets the "image_embedding" field.
func (goeu *GenerationOutputEmbedUpdate) SetImageEmbedding(pg pgvector.Vector) *GenerationOutputEmbedUpdate {
	goeu.mutation.SetImageEmbedding(pg)
	return goeu
}

// SetOutputID sets the "output_id" field.
func (goeu *GenerationOutputEmbedUpdate) SetOutputID(u uuid.UUID) *GenerationOutputEmbedUpdate {
	goeu.mutation.SetOutputID(u)
	return goeu
}

// SetUpdatedAt sets the "updated_at" field.
func (goeu *GenerationOutputEmbedUpdate) SetUpdatedAt(t time.Time) *GenerationOutputEmbedUpdate {
	goeu.mutation.SetUpdatedAt(t)
	return goeu
}

// SetGenerationOutputsID sets the "generation_outputs" edge to the GenerationOutput entity by ID.
func (goeu *GenerationOutputEmbedUpdate) SetGenerationOutputsID(id uuid.UUID) *GenerationOutputEmbedUpdate {
	goeu.mutation.SetGenerationOutputsID(id)
	return goeu
}

// SetGenerationOutputs sets the "generation_outputs" edge to the GenerationOutput entity.
func (goeu *GenerationOutputEmbedUpdate) SetGenerationOutputs(g *GenerationOutput) *GenerationOutputEmbedUpdate {
	return goeu.SetGenerationOutputsID(g.ID)
}

// Mutation returns the GenerationOutputEmbedMutation object of the builder.
func (goeu *GenerationOutputEmbedUpdate) Mutation() *GenerationOutputEmbedMutation {
	return goeu.mutation
}

// ClearGenerationOutputs clears the "generation_outputs" edge to the GenerationOutput entity.
func (goeu *GenerationOutputEmbedUpdate) ClearGenerationOutputs() *GenerationOutputEmbedUpdate {
	goeu.mutation.ClearGenerationOutputs()
	return goeu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (goeu *GenerationOutputEmbedUpdate) Save(ctx context.Context) (int, error) {
	goeu.defaults()
	return withHooks[int, GenerationOutputEmbedMutation](ctx, goeu.sqlSave, goeu.mutation, goeu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (goeu *GenerationOutputEmbedUpdate) SaveX(ctx context.Context) int {
	affected, err := goeu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (goeu *GenerationOutputEmbedUpdate) Exec(ctx context.Context) error {
	_, err := goeu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (goeu *GenerationOutputEmbedUpdate) ExecX(ctx context.Context) {
	if err := goeu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (goeu *GenerationOutputEmbedUpdate) defaults() {
	if _, ok := goeu.mutation.UpdatedAt(); !ok {
		v := generationoutputembed.UpdateDefaultUpdatedAt()
		goeu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (goeu *GenerationOutputEmbedUpdate) check() error {
	if _, ok := goeu.mutation.GenerationOutputsID(); goeu.mutation.GenerationOutputsCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "GenerationOutputEmbed.generation_outputs"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (goeu *GenerationOutputEmbedUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *GenerationOutputEmbedUpdate {
	goeu.modifiers = append(goeu.modifiers, modifiers...)
	return goeu
}

func (goeu *GenerationOutputEmbedUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := goeu.check(); err != nil {
		return n, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   generationoutputembed.Table,
			Columns: generationoutputembed.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: generationoutputembed.FieldID,
			},
		},
	}
	if ps := goeu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := goeu.mutation.PromptEmbedding(); ok {
		_spec.SetField(generationoutputembed.FieldPromptEmbedding, field.TypeOther, value)
	}
	if value, ok := goeu.mutation.ImageEmbedding(); ok {
		_spec.SetField(generationoutputembed.FieldImageEmbedding, field.TypeOther, value)
	}
	if value, ok := goeu.mutation.UpdatedAt(); ok {
		_spec.SetField(generationoutputembed.FieldUpdatedAt, field.TypeTime, value)
	}
	if goeu.mutation.GenerationOutputsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := goeu.mutation.GenerationOutputsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(goeu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, goeu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{generationoutputembed.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	goeu.mutation.done = true
	return n, nil
}

// GenerationOutputEmbedUpdateOne is the builder for updating a single GenerationOutputEmbed entity.
type GenerationOutputEmbedUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *GenerationOutputEmbedMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetPromptEmbedding sets the "prompt_embedding" field.
func (goeuo *GenerationOutputEmbedUpdateOne) SetPromptEmbedding(pg pgvector.Vector) *GenerationOutputEmbedUpdateOne {
	goeuo.mutation.SetPromptEmbedding(pg)
	return goeuo
}

// SetImageEmbedding sets the "image_embedding" field.
func (goeuo *GenerationOutputEmbedUpdateOne) SetImageEmbedding(pg pgvector.Vector) *GenerationOutputEmbedUpdateOne {
	goeuo.mutation.SetImageEmbedding(pg)
	return goeuo
}

// SetOutputID sets the "output_id" field.
func (goeuo *GenerationOutputEmbedUpdateOne) SetOutputID(u uuid.UUID) *GenerationOutputEmbedUpdateOne {
	goeuo.mutation.SetOutputID(u)
	return goeuo
}

// SetUpdatedAt sets the "updated_at" field.
func (goeuo *GenerationOutputEmbedUpdateOne) SetUpdatedAt(t time.Time) *GenerationOutputEmbedUpdateOne {
	goeuo.mutation.SetUpdatedAt(t)
	return goeuo
}

// SetGenerationOutputsID sets the "generation_outputs" edge to the GenerationOutput entity by ID.
func (goeuo *GenerationOutputEmbedUpdateOne) SetGenerationOutputsID(id uuid.UUID) *GenerationOutputEmbedUpdateOne {
	goeuo.mutation.SetGenerationOutputsID(id)
	return goeuo
}

// SetGenerationOutputs sets the "generation_outputs" edge to the GenerationOutput entity.
func (goeuo *GenerationOutputEmbedUpdateOne) SetGenerationOutputs(g *GenerationOutput) *GenerationOutputEmbedUpdateOne {
	return goeuo.SetGenerationOutputsID(g.ID)
}

// Mutation returns the GenerationOutputEmbedMutation object of the builder.
func (goeuo *GenerationOutputEmbedUpdateOne) Mutation() *GenerationOutputEmbedMutation {
	return goeuo.mutation
}

// ClearGenerationOutputs clears the "generation_outputs" edge to the GenerationOutput entity.
func (goeuo *GenerationOutputEmbedUpdateOne) ClearGenerationOutputs() *GenerationOutputEmbedUpdateOne {
	goeuo.mutation.ClearGenerationOutputs()
	return goeuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (goeuo *GenerationOutputEmbedUpdateOne) Select(field string, fields ...string) *GenerationOutputEmbedUpdateOne {
	goeuo.fields = append([]string{field}, fields...)
	return goeuo
}

// Save executes the query and returns the updated GenerationOutputEmbed entity.
func (goeuo *GenerationOutputEmbedUpdateOne) Save(ctx context.Context) (*GenerationOutputEmbed, error) {
	goeuo.defaults()
	return withHooks[*GenerationOutputEmbed, GenerationOutputEmbedMutation](ctx, goeuo.sqlSave, goeuo.mutation, goeuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (goeuo *GenerationOutputEmbedUpdateOne) SaveX(ctx context.Context) *GenerationOutputEmbed {
	node, err := goeuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (goeuo *GenerationOutputEmbedUpdateOne) Exec(ctx context.Context) error {
	_, err := goeuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (goeuo *GenerationOutputEmbedUpdateOne) ExecX(ctx context.Context) {
	if err := goeuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (goeuo *GenerationOutputEmbedUpdateOne) defaults() {
	if _, ok := goeuo.mutation.UpdatedAt(); !ok {
		v := generationoutputembed.UpdateDefaultUpdatedAt()
		goeuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (goeuo *GenerationOutputEmbedUpdateOne) check() error {
	if _, ok := goeuo.mutation.GenerationOutputsID(); goeuo.mutation.GenerationOutputsCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "GenerationOutputEmbed.generation_outputs"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (goeuo *GenerationOutputEmbedUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *GenerationOutputEmbedUpdateOne {
	goeuo.modifiers = append(goeuo.modifiers, modifiers...)
	return goeuo
}

func (goeuo *GenerationOutputEmbedUpdateOne) sqlSave(ctx context.Context) (_node *GenerationOutputEmbed, err error) {
	if err := goeuo.check(); err != nil {
		return _node, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   generationoutputembed.Table,
			Columns: generationoutputembed.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: generationoutputembed.FieldID,
			},
		},
	}
	id, ok := goeuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "GenerationOutputEmbed.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := goeuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, generationoutputembed.FieldID)
		for _, f := range fields {
			if !generationoutputembed.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != generationoutputembed.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := goeuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := goeuo.mutation.PromptEmbedding(); ok {
		_spec.SetField(generationoutputembed.FieldPromptEmbedding, field.TypeOther, value)
	}
	if value, ok := goeuo.mutation.ImageEmbedding(); ok {
		_spec.SetField(generationoutputembed.FieldImageEmbedding, field.TypeOther, value)
	}
	if value, ok := goeuo.mutation.UpdatedAt(); ok {
		_spec.SetField(generationoutputembed.FieldUpdatedAt, field.TypeTime, value)
	}
	if goeuo.mutation.GenerationOutputsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := goeuo.mutation.GenerationOutputsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(goeuo.modifiers...)
	_node = &GenerationOutputEmbed{config: goeuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, goeuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{generationoutputembed.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	goeuo.mutation.done = true
	return _node, nil
}