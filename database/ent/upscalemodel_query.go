// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/stablecog/go-apps/database/ent/predicate"
	"github.com/stablecog/go-apps/database/ent/upscale"
	"github.com/stablecog/go-apps/database/ent/upscalemodel"
)

// UpscaleModelQuery is the builder for querying UpscaleModel entities.
type UpscaleModelQuery struct {
	config
	limit        *int
	offset       *int
	unique       *bool
	order        []OrderFunc
	fields       []string
	inters       []Interceptor
	predicates   []predicate.UpscaleModel
	withUpscales *UpscaleQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UpscaleModelQuery builder.
func (umq *UpscaleModelQuery) Where(ps ...predicate.UpscaleModel) *UpscaleModelQuery {
	umq.predicates = append(umq.predicates, ps...)
	return umq
}

// Limit the number of records to be returned by this query.
func (umq *UpscaleModelQuery) Limit(limit int) *UpscaleModelQuery {
	umq.limit = &limit
	return umq
}

// Offset to start from.
func (umq *UpscaleModelQuery) Offset(offset int) *UpscaleModelQuery {
	umq.offset = &offset
	return umq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (umq *UpscaleModelQuery) Unique(unique bool) *UpscaleModelQuery {
	umq.unique = &unique
	return umq
}

// Order specifies how the records should be ordered.
func (umq *UpscaleModelQuery) Order(o ...OrderFunc) *UpscaleModelQuery {
	umq.order = append(umq.order, o...)
	return umq
}

// QueryUpscales chains the current query on the "upscales" edge.
func (umq *UpscaleModelQuery) QueryUpscales() *UpscaleQuery {
	query := (&UpscaleClient{config: umq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := umq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := umq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(upscalemodel.Table, upscalemodel.FieldID, selector),
			sqlgraph.To(upscale.Table, upscale.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, upscalemodel.UpscalesTable, upscalemodel.UpscalesColumn),
		)
		fromU = sqlgraph.SetNeighbors(umq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UpscaleModel entity from the query.
// Returns a *NotFoundError when no UpscaleModel was found.
func (umq *UpscaleModelQuery) First(ctx context.Context) (*UpscaleModel, error) {
	nodes, err := umq.Limit(1).All(newQueryContext(ctx, TypeUpscaleModel, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{upscalemodel.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (umq *UpscaleModelQuery) FirstX(ctx context.Context) *UpscaleModel {
	node, err := umq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UpscaleModel ID from the query.
// Returns a *NotFoundError when no UpscaleModel ID was found.
func (umq *UpscaleModelQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = umq.Limit(1).IDs(newQueryContext(ctx, TypeUpscaleModel, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{upscalemodel.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (umq *UpscaleModelQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := umq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UpscaleModel entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UpscaleModel entity is found.
// Returns a *NotFoundError when no UpscaleModel entities are found.
func (umq *UpscaleModelQuery) Only(ctx context.Context) (*UpscaleModel, error) {
	nodes, err := umq.Limit(2).All(newQueryContext(ctx, TypeUpscaleModel, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{upscalemodel.Label}
	default:
		return nil, &NotSingularError{upscalemodel.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (umq *UpscaleModelQuery) OnlyX(ctx context.Context) *UpscaleModel {
	node, err := umq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UpscaleModel ID in the query.
// Returns a *NotSingularError when more than one UpscaleModel ID is found.
// Returns a *NotFoundError when no entities are found.
func (umq *UpscaleModelQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = umq.Limit(2).IDs(newQueryContext(ctx, TypeUpscaleModel, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{upscalemodel.Label}
	default:
		err = &NotSingularError{upscalemodel.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (umq *UpscaleModelQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := umq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UpscaleModels.
func (umq *UpscaleModelQuery) All(ctx context.Context) ([]*UpscaleModel, error) {
	ctx = newQueryContext(ctx, TypeUpscaleModel, "All")
	if err := umq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*UpscaleModel, *UpscaleModelQuery]()
	return withInterceptors[[]*UpscaleModel](ctx, umq, qr, umq.inters)
}

// AllX is like All, but panics if an error occurs.
func (umq *UpscaleModelQuery) AllX(ctx context.Context) []*UpscaleModel {
	nodes, err := umq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UpscaleModel IDs.
func (umq *UpscaleModelQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	ctx = newQueryContext(ctx, TypeUpscaleModel, "IDs")
	if err := umq.Select(upscalemodel.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (umq *UpscaleModelQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := umq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (umq *UpscaleModelQuery) Count(ctx context.Context) (int, error) {
	ctx = newQueryContext(ctx, TypeUpscaleModel, "Count")
	if err := umq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, umq, querierCount[*UpscaleModelQuery](), umq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (umq *UpscaleModelQuery) CountX(ctx context.Context) int {
	count, err := umq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (umq *UpscaleModelQuery) Exist(ctx context.Context) (bool, error) {
	ctx = newQueryContext(ctx, TypeUpscaleModel, "Exist")
	switch _, err := umq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (umq *UpscaleModelQuery) ExistX(ctx context.Context) bool {
	exist, err := umq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UpscaleModelQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (umq *UpscaleModelQuery) Clone() *UpscaleModelQuery {
	if umq == nil {
		return nil
	}
	return &UpscaleModelQuery{
		config:       umq.config,
		limit:        umq.limit,
		offset:       umq.offset,
		order:        append([]OrderFunc{}, umq.order...),
		inters:       append([]Interceptor{}, umq.inters...),
		predicates:   append([]predicate.UpscaleModel{}, umq.predicates...),
		withUpscales: umq.withUpscales.Clone(),
		// clone intermediate query.
		sql:    umq.sql.Clone(),
		path:   umq.path,
		unique: umq.unique,
	}
}

// WithUpscales tells the query-builder to eager-load the nodes that are connected to
// the "upscales" edge. The optional arguments are used to configure the query builder of the edge.
func (umq *UpscaleModelQuery) WithUpscales(opts ...func(*UpscaleQuery)) *UpscaleModelQuery {
	query := (&UpscaleClient{config: umq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	umq.withUpscales = query
	return umq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UpscaleModel.Query().
//		GroupBy(upscalemodel.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (umq *UpscaleModelQuery) GroupBy(field string, fields ...string) *UpscaleModelGroupBy {
	umq.fields = append([]string{field}, fields...)
	grbuild := &UpscaleModelGroupBy{build: umq}
	grbuild.flds = &umq.fields
	grbuild.label = upscalemodel.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.UpscaleModel.Query().
//		Select(upscalemodel.FieldName).
//		Scan(ctx, &v)
func (umq *UpscaleModelQuery) Select(fields ...string) *UpscaleModelSelect {
	umq.fields = append(umq.fields, fields...)
	sbuild := &UpscaleModelSelect{UpscaleModelQuery: umq}
	sbuild.label = upscalemodel.Label
	sbuild.flds, sbuild.scan = &umq.fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UpscaleModelSelect configured with the given aggregations.
func (umq *UpscaleModelQuery) Aggregate(fns ...AggregateFunc) *UpscaleModelSelect {
	return umq.Select().Aggregate(fns...)
}

func (umq *UpscaleModelQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range umq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, umq); err != nil {
				return err
			}
		}
	}
	for _, f := range umq.fields {
		if !upscalemodel.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if umq.path != nil {
		prev, err := umq.path(ctx)
		if err != nil {
			return err
		}
		umq.sql = prev
	}
	return nil
}

func (umq *UpscaleModelQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UpscaleModel, error) {
	var (
		nodes       = []*UpscaleModel{}
		_spec       = umq.querySpec()
		loadedTypes = [1]bool{
			umq.withUpscales != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*UpscaleModel).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &UpscaleModel{config: umq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, umq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := umq.withUpscales; query != nil {
		if err := umq.loadUpscales(ctx, query, nodes,
			func(n *UpscaleModel) { n.Edges.Upscales = []*Upscale{} },
			func(n *UpscaleModel, e *Upscale) { n.Edges.Upscales = append(n.Edges.Upscales, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (umq *UpscaleModelQuery) loadUpscales(ctx context.Context, query *UpscaleQuery, nodes []*UpscaleModel, init func(*UpscaleModel), assign func(*UpscaleModel, *Upscale)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*UpscaleModel)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.Upscale(func(s *sql.Selector) {
		s.Where(sql.InValues(upscalemodel.UpscalesColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ModelID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "model_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (umq *UpscaleModelQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := umq.querySpec()
	_spec.Node.Columns = umq.fields
	if len(umq.fields) > 0 {
		_spec.Unique = umq.unique != nil && *umq.unique
	}
	return sqlgraph.CountNodes(ctx, umq.driver, _spec)
}

func (umq *UpscaleModelQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   upscalemodel.Table,
			Columns: upscalemodel.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: upscalemodel.FieldID,
			},
		},
		From:   umq.sql,
		Unique: true,
	}
	if unique := umq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := umq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, upscalemodel.FieldID)
		for i := range fields {
			if fields[i] != upscalemodel.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := umq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := umq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := umq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := umq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (umq *UpscaleModelQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(umq.driver.Dialect())
	t1 := builder.Table(upscalemodel.Table)
	columns := umq.fields
	if len(columns) == 0 {
		columns = upscalemodel.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if umq.sql != nil {
		selector = umq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if umq.unique != nil && *umq.unique {
		selector.Distinct()
	}
	for _, p := range umq.predicates {
		p(selector)
	}
	for _, p := range umq.order {
		p(selector)
	}
	if offset := umq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := umq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UpscaleModelGroupBy is the group-by builder for UpscaleModel entities.
type UpscaleModelGroupBy struct {
	selector
	build *UpscaleModelQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (umgb *UpscaleModelGroupBy) Aggregate(fns ...AggregateFunc) *UpscaleModelGroupBy {
	umgb.fns = append(umgb.fns, fns...)
	return umgb
}

// Scan applies the selector query and scans the result into the given value.
func (umgb *UpscaleModelGroupBy) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeUpscaleModel, "GroupBy")
	if err := umgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UpscaleModelQuery, *UpscaleModelGroupBy](ctx, umgb.build, umgb, umgb.build.inters, v)
}

func (umgb *UpscaleModelGroupBy) sqlScan(ctx context.Context, root *UpscaleModelQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(umgb.fns))
	for _, fn := range umgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*umgb.flds)+len(umgb.fns))
		for _, f := range *umgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*umgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := umgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UpscaleModelSelect is the builder for selecting fields of UpscaleModel entities.
type UpscaleModelSelect struct {
	*UpscaleModelQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ums *UpscaleModelSelect) Aggregate(fns ...AggregateFunc) *UpscaleModelSelect {
	ums.fns = append(ums.fns, fns...)
	return ums
}

// Scan applies the selector query and scans the result into the given value.
func (ums *UpscaleModelSelect) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeUpscaleModel, "Select")
	if err := ums.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UpscaleModelQuery, *UpscaleModelSelect](ctx, ums.UpscaleModelQuery, ums, ums.inters, v)
}

func (ums *UpscaleModelSelect) sqlScan(ctx context.Context, root *UpscaleModelQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ums.fns))
	for _, fn := range ums.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ums.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ums.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
