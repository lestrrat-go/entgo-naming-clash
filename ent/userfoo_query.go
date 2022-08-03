// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/bug/ent/predicate"
	"entgo.io/bug/ent/user"
	"entgo.io/bug/ent/userfoo"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserFooQuery is the builder for querying UserFoo entities.
type UserFooQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.UserFoo
	// eager-loading edges.
	withParent *UserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserFooQuery builder.
func (ufq *UserFooQuery) Where(ps ...predicate.UserFoo) *UserFooQuery {
	ufq.predicates = append(ufq.predicates, ps...)
	return ufq
}

// Limit adds a limit step to the query.
func (ufq *UserFooQuery) Limit(limit int) *UserFooQuery {
	ufq.limit = &limit
	return ufq
}

// Offset adds an offset step to the query.
func (ufq *UserFooQuery) Offset(offset int) *UserFooQuery {
	ufq.offset = &offset
	return ufq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ufq *UserFooQuery) Unique(unique bool) *UserFooQuery {
	ufq.unique = &unique
	return ufq
}

// Order adds an order step to the query.
func (ufq *UserFooQuery) Order(o ...OrderFunc) *UserFooQuery {
	ufq.order = append(ufq.order, o...)
	return ufq
}

// QueryParent chains the current query on the "parent" edge.
func (ufq *UserFooQuery) QueryParent() *UserQuery {
	query := &UserQuery{config: ufq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ufq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ufq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(userfoo.Table, userfoo.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, userfoo.ParentTable, userfoo.ParentPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(ufq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UserFoo entity from the query.
// Returns a *NotFoundError when no UserFoo was found.
func (ufq *UserFooQuery) First(ctx context.Context) (*UserFoo, error) {
	nodes, err := ufq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{userfoo.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ufq *UserFooQuery) FirstX(ctx context.Context) *UserFoo {
	node, err := ufq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserFoo ID from the query.
// Returns a *NotFoundError when no UserFoo ID was found.
func (ufq *UserFooQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ufq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{userfoo.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ufq *UserFooQuery) FirstIDX(ctx context.Context) int {
	id, err := ufq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UserFoo entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UserFoo entity is found.
// Returns a *NotFoundError when no UserFoo entities are found.
func (ufq *UserFooQuery) Only(ctx context.Context) (*UserFoo, error) {
	nodes, err := ufq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{userfoo.Label}
	default:
		return nil, &NotSingularError{userfoo.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ufq *UserFooQuery) OnlyX(ctx context.Context) *UserFoo {
	node, err := ufq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UserFoo ID in the query.
// Returns a *NotSingularError when more than one UserFoo ID is found.
// Returns a *NotFoundError when no entities are found.
func (ufq *UserFooQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ufq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{userfoo.Label}
	default:
		err = &NotSingularError{userfoo.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ufq *UserFooQuery) OnlyIDX(ctx context.Context) int {
	id, err := ufq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserFoos.
func (ufq *UserFooQuery) All(ctx context.Context) ([]*UserFoo, error) {
	if err := ufq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ufq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ufq *UserFooQuery) AllX(ctx context.Context) []*UserFoo {
	nodes, err := ufq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserFoo IDs.
func (ufq *UserFooQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := ufq.Select(userfoo.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ufq *UserFooQuery) IDsX(ctx context.Context) []int {
	ids, err := ufq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ufq *UserFooQuery) Count(ctx context.Context) (int, error) {
	if err := ufq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ufq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ufq *UserFooQuery) CountX(ctx context.Context) int {
	count, err := ufq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ufq *UserFooQuery) Exist(ctx context.Context) (bool, error) {
	if err := ufq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ufq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ufq *UserFooQuery) ExistX(ctx context.Context) bool {
	exist, err := ufq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserFooQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ufq *UserFooQuery) Clone() *UserFooQuery {
	if ufq == nil {
		return nil
	}
	return &UserFooQuery{
		config:     ufq.config,
		limit:      ufq.limit,
		offset:     ufq.offset,
		order:      append([]OrderFunc{}, ufq.order...),
		predicates: append([]predicate.UserFoo{}, ufq.predicates...),
		withParent: ufq.withParent.Clone(),
		// clone intermediate query.
		sql:    ufq.sql.Clone(),
		path:   ufq.path,
		unique: ufq.unique,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (ufq *UserFooQuery) WithParent(opts ...func(*UserQuery)) *UserFooQuery {
	query := &UserQuery{config: ufq.config}
	for _, opt := range opts {
		opt(query)
	}
	ufq.withParent = query
	return ufq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Dummy string `json:"dummy,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UserFoo.Query().
//		GroupBy(userfoo.FieldDummy).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ufq *UserFooQuery) GroupBy(field string, fields ...string) *UserFooGroupBy {
	grbuild := &UserFooGroupBy{config: ufq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ufq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ufq.sqlQuery(ctx), nil
	}
	grbuild.label = userfoo.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Dummy string `json:"dummy,omitempty"`
//	}
//
//	client.UserFoo.Query().
//		Select(userfoo.FieldDummy).
//		Scan(ctx, &v)
//
func (ufq *UserFooQuery) Select(fields ...string) *UserFooSelect {
	ufq.fields = append(ufq.fields, fields...)
	selbuild := &UserFooSelect{UserFooQuery: ufq}
	selbuild.label = userfoo.Label
	selbuild.flds, selbuild.scan = &ufq.fields, selbuild.Scan
	return selbuild
}

func (ufq *UserFooQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ufq.fields {
		if !userfoo.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ufq.path != nil {
		prev, err := ufq.path(ctx)
		if err != nil {
			return err
		}
		ufq.sql = prev
	}
	return nil
}

func (ufq *UserFooQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UserFoo, error) {
	var (
		nodes       = []*UserFoo{}
		_spec       = ufq.querySpec()
		loadedTypes = [1]bool{
			ufq.withParent != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*UserFoo).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &UserFoo{config: ufq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ufq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := ufq.withParent; query != nil {
		edgeids := make([]driver.Value, len(nodes))
		byid := make(map[int]*UserFoo)
		nids := make(map[int]map[*UserFoo]struct{})
		for i, node := range nodes {
			edgeids[i] = node.ID
			byid[node.ID] = node
			node.Edges.Parent = []*User{}
		}
		query.Where(func(s *sql.Selector) {
			joinT := sql.Table(userfoo.ParentTable)
			s.Join(joinT).On(s.C(user.FieldID), joinT.C(userfoo.ParentPrimaryKey[0]))
			s.Where(sql.InValues(joinT.C(userfoo.ParentPrimaryKey[1]), edgeids...))
			columns := s.SelectedColumns()
			s.Select(joinT.C(userfoo.ParentPrimaryKey[1]))
			s.AppendSelect(columns...)
			s.SetDistinct(false)
		})
		neighbors, err := query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]interface{}, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]interface{}{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []interface{}) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*UserFoo]struct{}{byid[outValue]: struct{}{}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byid[outValue]] = struct{}{}
				return nil
			}
		})
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "parent" node returned %v`, n.ID)
			}
			for kn := range nodes {
				kn.Edges.Parent = append(kn.Edges.Parent, n)
			}
		}
	}

	return nodes, nil
}

func (ufq *UserFooQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ufq.querySpec()
	_spec.Node.Columns = ufq.fields
	if len(ufq.fields) > 0 {
		_spec.Unique = ufq.unique != nil && *ufq.unique
	}
	return sqlgraph.CountNodes(ctx, ufq.driver, _spec)
}

func (ufq *UserFooQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ufq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (ufq *UserFooQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userfoo.Table,
			Columns: userfoo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: userfoo.FieldID,
			},
		},
		From:   ufq.sql,
		Unique: true,
	}
	if unique := ufq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ufq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userfoo.FieldID)
		for i := range fields {
			if fields[i] != userfoo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ufq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ufq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ufq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ufq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ufq *UserFooQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ufq.driver.Dialect())
	t1 := builder.Table(userfoo.Table)
	columns := ufq.fields
	if len(columns) == 0 {
		columns = userfoo.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ufq.sql != nil {
		selector = ufq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ufq.unique != nil && *ufq.unique {
		selector.Distinct()
	}
	for _, p := range ufq.predicates {
		p(selector)
	}
	for _, p := range ufq.order {
		p(selector)
	}
	if offset := ufq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ufq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserFooGroupBy is the group-by builder for UserFoo entities.
type UserFooGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ufgb *UserFooGroupBy) Aggregate(fns ...AggregateFunc) *UserFooGroupBy {
	ufgb.fns = append(ufgb.fns, fns...)
	return ufgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ufgb *UserFooGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ufgb.path(ctx)
	if err != nil {
		return err
	}
	ufgb.sql = query
	return ufgb.sqlScan(ctx, v)
}

func (ufgb *UserFooGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ufgb.fields {
		if !userfoo.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ufgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ufgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ufgb *UserFooGroupBy) sqlQuery() *sql.Selector {
	selector := ufgb.sql.Select()
	aggregation := make([]string, 0, len(ufgb.fns))
	for _, fn := range ufgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ufgb.fields)+len(ufgb.fns))
		for _, f := range ufgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ufgb.fields...)...)
}

// UserFooSelect is the builder for selecting fields of UserFoo entities.
type UserFooSelect struct {
	*UserFooQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ufs *UserFooSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ufs.prepareQuery(ctx); err != nil {
		return err
	}
	ufs.sql = ufs.UserFooQuery.sqlQuery(ctx)
	return ufs.sqlScan(ctx, v)
}

func (ufs *UserFooSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ufs.sql.Query()
	if err := ufs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
