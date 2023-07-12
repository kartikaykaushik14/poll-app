// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"poll-app/ent/poll"
	"poll-app/ent/polloption"
	"poll-app/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PollQuery is the builder for querying Poll entities.
type PollQuery struct {
	config
	ctx             *QueryContext
	order           []poll.OrderOption
	inters          []Interceptor
	predicates      []predicate.Poll
	withPollOptions *PollOptionQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PollQuery builder.
func (pq *PollQuery) Where(ps ...predicate.Poll) *PollQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit the number of records to be returned by this query.
func (pq *PollQuery) Limit(limit int) *PollQuery {
	pq.ctx.Limit = &limit
	return pq
}

// Offset to start from.
func (pq *PollQuery) Offset(offset int) *PollQuery {
	pq.ctx.Offset = &offset
	return pq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pq *PollQuery) Unique(unique bool) *PollQuery {
	pq.ctx.Unique = &unique
	return pq
}

// Order specifies how the records should be ordered.
func (pq *PollQuery) Order(o ...poll.OrderOption) *PollQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// QueryPollOptions chains the current query on the "poll_options" edge.
func (pq *PollQuery) QueryPollOptions() *PollOptionQuery {
	query := (&PollOptionClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(poll.Table, poll.FieldID, selector),
			sqlgraph.To(polloption.Table, polloption.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, poll.PollOptionsTable, poll.PollOptionsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Poll entity from the query.
// Returns a *NotFoundError when no Poll was found.
func (pq *PollQuery) First(ctx context.Context) (*Poll, error) {
	nodes, err := pq.Limit(1).All(setContextOp(ctx, pq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{poll.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *PollQuery) FirstX(ctx context.Context) *Poll {
	node, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Poll ID from the query.
// Returns a *NotFoundError when no Poll ID was found.
func (pq *PollQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(1).IDs(setContextOp(ctx, pq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{poll.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pq *PollQuery) FirstIDX(ctx context.Context) int {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Poll entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Poll entity is found.
// Returns a *NotFoundError when no Poll entities are found.
func (pq *PollQuery) Only(ctx context.Context) (*Poll, error) {
	nodes, err := pq.Limit(2).All(setContextOp(ctx, pq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{poll.Label}
	default:
		return nil, &NotSingularError{poll.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *PollQuery) OnlyX(ctx context.Context) *Poll {
	node, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Poll ID in the query.
// Returns a *NotSingularError when more than one Poll ID is found.
// Returns a *NotFoundError when no entities are found.
func (pq *PollQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(2).IDs(setContextOp(ctx, pq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{poll.Label}
	default:
		err = &NotSingularError{poll.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *PollQuery) OnlyIDX(ctx context.Context) int {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Polls.
func (pq *PollQuery) All(ctx context.Context) ([]*Poll, error) {
	ctx = setContextOp(ctx, pq.ctx, "All")
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Poll, *PollQuery]()
	return withInterceptors[[]*Poll](ctx, pq, qr, pq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pq *PollQuery) AllX(ctx context.Context) []*Poll {
	nodes, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Poll IDs.
func (pq *PollQuery) IDs(ctx context.Context) (ids []int, err error) {
	if pq.ctx.Unique == nil && pq.path != nil {
		pq.Unique(true)
	}
	ctx = setContextOp(ctx, pq.ctx, "IDs")
	if err = pq.Select(poll.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *PollQuery) IDsX(ctx context.Context) []int {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *PollQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pq.ctx, "Count")
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pq, querierCount[*PollQuery](), pq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pq *PollQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *PollQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pq.ctx, "Exist")
	switch _, err := pq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *PollQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PollQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *PollQuery) Clone() *PollQuery {
	if pq == nil {
		return nil
	}
	return &PollQuery{
		config:          pq.config,
		ctx:             pq.ctx.Clone(),
		order:           append([]poll.OrderOption{}, pq.order...),
		inters:          append([]Interceptor{}, pq.inters...),
		predicates:      append([]predicate.Poll{}, pq.predicates...),
		withPollOptions: pq.withPollOptions.Clone(),
		// clone intermediate query.
		sql:  pq.sql.Clone(),
		path: pq.path,
	}
}

// WithPollOptions tells the query-builder to eager-load the nodes that are connected to
// the "poll_options" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PollQuery) WithPollOptions(opts ...func(*PollOptionQuery)) *PollQuery {
	query := (&PollOptionClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withPollOptions = query
	return pq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Question string `json:"question,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Poll.Query().
//		GroupBy(poll.FieldQuestion).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pq *PollQuery) GroupBy(field string, fields ...string) *PollGroupBy {
	pq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PollGroupBy{build: pq}
	grbuild.flds = &pq.ctx.Fields
	grbuild.label = poll.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Question string `json:"question,omitempty"`
//	}
//
//	client.Poll.Query().
//		Select(poll.FieldQuestion).
//		Scan(ctx, &v)
func (pq *PollQuery) Select(fields ...string) *PollSelect {
	pq.ctx.Fields = append(pq.ctx.Fields, fields...)
	sbuild := &PollSelect{PollQuery: pq}
	sbuild.label = poll.Label
	sbuild.flds, sbuild.scan = &pq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PollSelect configured with the given aggregations.
func (pq *PollQuery) Aggregate(fns ...AggregateFunc) *PollSelect {
	return pq.Select().Aggregate(fns...)
}

func (pq *PollQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pq); err != nil {
				return err
			}
		}
	}
	for _, f := range pq.ctx.Fields {
		if !poll.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pq.path != nil {
		prev, err := pq.path(ctx)
		if err != nil {
			return err
		}
		pq.sql = prev
	}
	return nil
}

func (pq *PollQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Poll, error) {
	var (
		nodes       = []*Poll{}
		_spec       = pq.querySpec()
		loadedTypes = [1]bool{
			pq.withPollOptions != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Poll).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Poll{config: pq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pq.withPollOptions; query != nil {
		if err := pq.loadPollOptions(ctx, query, nodes,
			func(n *Poll) { n.Edges.PollOptions = []*PollOption{} },
			func(n *Poll, e *PollOption) { n.Edges.PollOptions = append(n.Edges.PollOptions, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pq *PollQuery) loadPollOptions(ctx context.Context, query *PollOptionQuery, nodes []*Poll, init func(*Poll), assign func(*Poll, *PollOption)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Poll)
	nids := make(map[int]map[*Poll]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(poll.PollOptionsTable)
		s.Join(joinT).On(s.C(polloption.FieldID), joinT.C(poll.PollOptionsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(poll.PollOptionsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(poll.PollOptionsPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*Poll]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*PollOption](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "poll_options" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (pq *PollQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	_spec.Node.Columns = pq.ctx.Fields
	if len(pq.ctx.Fields) > 0 {
		_spec.Unique = pq.ctx.Unique != nil && *pq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *PollQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(poll.Table, poll.Columns, sqlgraph.NewFieldSpec(poll.FieldID, field.TypeInt))
	_spec.From = pq.sql
	if unique := pq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pq.path != nil {
		_spec.Unique = true
	}
	if fields := pq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, poll.FieldID)
		for i := range fields {
			if fields[i] != poll.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pq *PollQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(poll.Table)
	columns := pq.ctx.Fields
	if len(columns) == 0 {
		columns = poll.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pq.ctx.Unique != nil && *pq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range pq.predicates {
		p(selector)
	}
	for _, p := range pq.order {
		p(selector)
	}
	if offset := pq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PollGroupBy is the group-by builder for Poll entities.
type PollGroupBy struct {
	selector
	build *PollQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *PollGroupBy) Aggregate(fns ...AggregateFunc) *PollGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the selector query and scans the result into the given value.
func (pgb *PollGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pgb.build.ctx, "GroupBy")
	if err := pgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PollQuery, *PollGroupBy](ctx, pgb.build, pgb, pgb.build.inters, v)
}

func (pgb *PollGroupBy) sqlScan(ctx context.Context, root *PollQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pgb.fns))
	for _, fn := range pgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pgb.flds)+len(pgb.fns))
		for _, f := range *pgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PollSelect is the builder for selecting fields of Poll entities.
type PollSelect struct {
	*PollQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ps *PollSelect) Aggregate(fns ...AggregateFunc) *PollSelect {
	ps.fns = append(ps.fns, fns...)
	return ps
}

// Scan applies the selector query and scans the result into the given value.
func (ps *PollSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ps.ctx, "Select")
	if err := ps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PollQuery, *PollSelect](ctx, ps.PollQuery, ps, ps.inters, v)
}

func (ps *PollSelect) sqlScan(ctx context.Context, root *PollQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ps.fns))
	for _, fn := range ps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}