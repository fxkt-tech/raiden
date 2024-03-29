// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/fxkt-tech/raiden/app/mainsite/service/internal/data/ent/predicate"
	"github.com/fxkt-tech/raiden/app/mainsite/service/internal/data/ent/rduser"
)

// RdUserQuery is the builder for querying RdUser entities.
type RdUserQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.RdUser
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RdUserQuery builder.
func (ruq *RdUserQuery) Where(ps ...predicate.RdUser) *RdUserQuery {
	ruq.predicates = append(ruq.predicates, ps...)
	return ruq
}

// Limit the number of records to be returned by this query.
func (ruq *RdUserQuery) Limit(limit int) *RdUserQuery {
	ruq.ctx.Limit = &limit
	return ruq
}

// Offset to start from.
func (ruq *RdUserQuery) Offset(offset int) *RdUserQuery {
	ruq.ctx.Offset = &offset
	return ruq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ruq *RdUserQuery) Unique(unique bool) *RdUserQuery {
	ruq.ctx.Unique = &unique
	return ruq
}

// Order specifies how the records should be ordered.
func (ruq *RdUserQuery) Order(o ...OrderFunc) *RdUserQuery {
	ruq.order = append(ruq.order, o...)
	return ruq
}

// First returns the first RdUser entity from the query.
// Returns a *NotFoundError when no RdUser was found.
func (ruq *RdUserQuery) First(ctx context.Context) (*RdUser, error) {
	nodes, err := ruq.Limit(1).All(setContextOp(ctx, ruq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{rduser.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ruq *RdUserQuery) FirstX(ctx context.Context) *RdUser {
	node, err := ruq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first RdUser ID from the query.
// Returns a *NotFoundError when no RdUser ID was found.
func (ruq *RdUserQuery) FirstID(ctx context.Context) (id int32, err error) {
	var ids []int32
	if ids, err = ruq.Limit(1).IDs(setContextOp(ctx, ruq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{rduser.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ruq *RdUserQuery) FirstIDX(ctx context.Context) int32 {
	id, err := ruq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single RdUser entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one RdUser entity is found.
// Returns a *NotFoundError when no RdUser entities are found.
func (ruq *RdUserQuery) Only(ctx context.Context) (*RdUser, error) {
	nodes, err := ruq.Limit(2).All(setContextOp(ctx, ruq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{rduser.Label}
	default:
		return nil, &NotSingularError{rduser.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ruq *RdUserQuery) OnlyX(ctx context.Context) *RdUser {
	node, err := ruq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only RdUser ID in the query.
// Returns a *NotSingularError when more than one RdUser ID is found.
// Returns a *NotFoundError when no entities are found.
func (ruq *RdUserQuery) OnlyID(ctx context.Context) (id int32, err error) {
	var ids []int32
	if ids, err = ruq.Limit(2).IDs(setContextOp(ctx, ruq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{rduser.Label}
	default:
		err = &NotSingularError{rduser.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ruq *RdUserQuery) OnlyIDX(ctx context.Context) int32 {
	id, err := ruq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of RdUsers.
func (ruq *RdUserQuery) All(ctx context.Context) ([]*RdUser, error) {
	ctx = setContextOp(ctx, ruq.ctx, "All")
	if err := ruq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*RdUser, *RdUserQuery]()
	return withInterceptors[[]*RdUser](ctx, ruq, qr, ruq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ruq *RdUserQuery) AllX(ctx context.Context) []*RdUser {
	nodes, err := ruq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of RdUser IDs.
func (ruq *RdUserQuery) IDs(ctx context.Context) ([]int32, error) {
	var ids []int32
	ctx = setContextOp(ctx, ruq.ctx, "IDs")
	if err := ruq.Select(rduser.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ruq *RdUserQuery) IDsX(ctx context.Context) []int32 {
	ids, err := ruq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ruq *RdUserQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ruq.ctx, "Count")
	if err := ruq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ruq, querierCount[*RdUserQuery](), ruq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ruq *RdUserQuery) CountX(ctx context.Context) int {
	count, err := ruq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ruq *RdUserQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ruq.ctx, "Exist")
	switch _, err := ruq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ruq *RdUserQuery) ExistX(ctx context.Context) bool {
	exist, err := ruq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RdUserQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ruq *RdUserQuery) Clone() *RdUserQuery {
	if ruq == nil {
		return nil
	}
	return &RdUserQuery{
		config:     ruq.config,
		ctx:        ruq.ctx.Clone(),
		order:      append([]OrderFunc{}, ruq.order...),
		inters:     append([]Interceptor{}, ruq.inters...),
		predicates: append([]predicate.RdUser{}, ruq.predicates...),
		// clone intermediate query.
		sql:  ruq.sql.Clone(),
		path: ruq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Nick string `json:"nick,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.RdUser.Query().
//		GroupBy(rduser.FieldNick).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ruq *RdUserQuery) GroupBy(field string, fields ...string) *RdUserGroupBy {
	ruq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &RdUserGroupBy{build: ruq}
	grbuild.flds = &ruq.ctx.Fields
	grbuild.label = rduser.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Nick string `json:"nick,omitempty"`
//	}
//
//	client.RdUser.Query().
//		Select(rduser.FieldNick).
//		Scan(ctx, &v)
func (ruq *RdUserQuery) Select(fields ...string) *RdUserSelect {
	ruq.ctx.Fields = append(ruq.ctx.Fields, fields...)
	sbuild := &RdUserSelect{RdUserQuery: ruq}
	sbuild.label = rduser.Label
	sbuild.flds, sbuild.scan = &ruq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a RdUserSelect configured with the given aggregations.
func (ruq *RdUserQuery) Aggregate(fns ...AggregateFunc) *RdUserSelect {
	return ruq.Select().Aggregate(fns...)
}

func (ruq *RdUserQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ruq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ruq); err != nil {
				return err
			}
		}
	}
	for _, f := range ruq.ctx.Fields {
		if !rduser.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ruq.path != nil {
		prev, err := ruq.path(ctx)
		if err != nil {
			return err
		}
		ruq.sql = prev
	}
	return nil
}

func (ruq *RdUserQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*RdUser, error) {
	var (
		nodes = []*RdUser{}
		_spec = ruq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*RdUser).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &RdUser{config: ruq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ruq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (ruq *RdUserQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ruq.querySpec()
	_spec.Node.Columns = ruq.ctx.Fields
	if len(ruq.ctx.Fields) > 0 {
		_spec.Unique = ruq.ctx.Unique != nil && *ruq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ruq.driver, _spec)
}

func (ruq *RdUserQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   rduser.Table,
			Columns: rduser.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt32,
				Column: rduser.FieldID,
			},
		},
		From:   ruq.sql,
		Unique: true,
	}
	if unique := ruq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ruq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, rduser.FieldID)
		for i := range fields {
			if fields[i] != rduser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ruq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ruq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ruq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ruq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ruq *RdUserQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ruq.driver.Dialect())
	t1 := builder.Table(rduser.Table)
	columns := ruq.ctx.Fields
	if len(columns) == 0 {
		columns = rduser.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ruq.sql != nil {
		selector = ruq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ruq.ctx.Unique != nil && *ruq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ruq.predicates {
		p(selector)
	}
	for _, p := range ruq.order {
		p(selector)
	}
	if offset := ruq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ruq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// RdUserGroupBy is the group-by builder for RdUser entities.
type RdUserGroupBy struct {
	selector
	build *RdUserQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rugb *RdUserGroupBy) Aggregate(fns ...AggregateFunc) *RdUserGroupBy {
	rugb.fns = append(rugb.fns, fns...)
	return rugb
}

// Scan applies the selector query and scans the result into the given value.
func (rugb *RdUserGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rugb.build.ctx, "GroupBy")
	if err := rugb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RdUserQuery, *RdUserGroupBy](ctx, rugb.build, rugb, rugb.build.inters, v)
}

func (rugb *RdUserGroupBy) sqlScan(ctx context.Context, root *RdUserQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(rugb.fns))
	for _, fn := range rugb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*rugb.flds)+len(rugb.fns))
		for _, f := range *rugb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*rugb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rugb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// RdUserSelect is the builder for selecting fields of RdUser entities.
type RdUserSelect struct {
	*RdUserQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rus *RdUserSelect) Aggregate(fns ...AggregateFunc) *RdUserSelect {
	rus.fns = append(rus.fns, fns...)
	return rus
}

// Scan applies the selector query and scans the result into the given value.
func (rus *RdUserSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rus.ctx, "Select")
	if err := rus.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RdUserQuery, *RdUserSelect](ctx, rus.RdUserQuery, rus, rus.inters, v)
}

func (rus *RdUserSelect) sqlScan(ctx context.Context, root *RdUserQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(rus.fns))
	for _, fn := range rus.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*rus.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rus.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
