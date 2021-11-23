// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/goodfee"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// GoodFeeQuery is the builder for querying GoodFee entities.
type GoodFeeQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.GoodFee
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GoodFeeQuery builder.
func (gfq *GoodFeeQuery) Where(ps ...predicate.GoodFee) *GoodFeeQuery {
	gfq.predicates = append(gfq.predicates, ps...)
	return gfq
}

// Limit adds a limit step to the query.
func (gfq *GoodFeeQuery) Limit(limit int) *GoodFeeQuery {
	gfq.limit = &limit
	return gfq
}

// Offset adds an offset step to the query.
func (gfq *GoodFeeQuery) Offset(offset int) *GoodFeeQuery {
	gfq.offset = &offset
	return gfq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (gfq *GoodFeeQuery) Unique(unique bool) *GoodFeeQuery {
	gfq.unique = &unique
	return gfq
}

// Order adds an order step to the query.
func (gfq *GoodFeeQuery) Order(o ...OrderFunc) *GoodFeeQuery {
	gfq.order = append(gfq.order, o...)
	return gfq
}

// First returns the first GoodFee entity from the query.
// Returns a *NotFoundError when no GoodFee was found.
func (gfq *GoodFeeQuery) First(ctx context.Context) (*GoodFee, error) {
	nodes, err := gfq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{goodfee.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (gfq *GoodFeeQuery) FirstX(ctx context.Context) *GoodFee {
	node, err := gfq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first GoodFee ID from the query.
// Returns a *NotFoundError when no GoodFee ID was found.
func (gfq *GoodFeeQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = gfq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{goodfee.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (gfq *GoodFeeQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := gfq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single GoodFee entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one GoodFee entity is not found.
// Returns a *NotFoundError when no GoodFee entities are found.
func (gfq *GoodFeeQuery) Only(ctx context.Context) (*GoodFee, error) {
	nodes, err := gfq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{goodfee.Label}
	default:
		return nil, &NotSingularError{goodfee.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (gfq *GoodFeeQuery) OnlyX(ctx context.Context) *GoodFee {
	node, err := gfq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only GoodFee ID in the query.
// Returns a *NotSingularError when exactly one GoodFee ID is not found.
// Returns a *NotFoundError when no entities are found.
func (gfq *GoodFeeQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = gfq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{goodfee.Label}
	default:
		err = &NotSingularError{goodfee.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (gfq *GoodFeeQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := gfq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GoodFees.
func (gfq *GoodFeeQuery) All(ctx context.Context) ([]*GoodFee, error) {
	if err := gfq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return gfq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (gfq *GoodFeeQuery) AllX(ctx context.Context) []*GoodFee {
	nodes, err := gfq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of GoodFee IDs.
func (gfq *GoodFeeQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := gfq.Select(goodfee.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (gfq *GoodFeeQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := gfq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (gfq *GoodFeeQuery) Count(ctx context.Context) (int, error) {
	if err := gfq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return gfq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (gfq *GoodFeeQuery) CountX(ctx context.Context) int {
	count, err := gfq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (gfq *GoodFeeQuery) Exist(ctx context.Context) (bool, error) {
	if err := gfq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return gfq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (gfq *GoodFeeQuery) ExistX(ctx context.Context) bool {
	exist, err := gfq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GoodFeeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (gfq *GoodFeeQuery) Clone() *GoodFeeQuery {
	if gfq == nil {
		return nil
	}
	return &GoodFeeQuery{
		config:     gfq.config,
		limit:      gfq.limit,
		offset:     gfq.offset,
		order:      append([]OrderFunc{}, gfq.order...),
		predicates: append([]predicate.GoodFee{}, gfq.predicates...),
		// clone intermediate query.
		sql:  gfq.sql.Clone(),
		path: gfq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		AppID uuid.UUID `json:"app_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.GoodFee.Query().
//		GroupBy(goodfee.FieldAppID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (gfq *GoodFeeQuery) GroupBy(field string, fields ...string) *GoodFeeGroupBy {
	group := &GoodFeeGroupBy{config: gfq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := gfq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return gfq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		AppID uuid.UUID `json:"app_id,omitempty"`
//	}
//
//	client.GoodFee.Query().
//		Select(goodfee.FieldAppID).
//		Scan(ctx, &v)
//
func (gfq *GoodFeeQuery) Select(fields ...string) *GoodFeeSelect {
	gfq.fields = append(gfq.fields, fields...)
	return &GoodFeeSelect{GoodFeeQuery: gfq}
}

func (gfq *GoodFeeQuery) prepareQuery(ctx context.Context) error {
	for _, f := range gfq.fields {
		if !goodfee.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if gfq.path != nil {
		prev, err := gfq.path(ctx)
		if err != nil {
			return err
		}
		gfq.sql = prev
	}
	return nil
}

func (gfq *GoodFeeQuery) sqlAll(ctx context.Context) ([]*GoodFee, error) {
	var (
		nodes = []*GoodFee{}
		_spec = gfq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &GoodFee{config: gfq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, gfq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (gfq *GoodFeeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := gfq.querySpec()
	return sqlgraph.CountNodes(ctx, gfq.driver, _spec)
}

func (gfq *GoodFeeQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := gfq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (gfq *GoodFeeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   goodfee.Table,
			Columns: goodfee.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: goodfee.FieldID,
			},
		},
		From:   gfq.sql,
		Unique: true,
	}
	if unique := gfq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := gfq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, goodfee.FieldID)
		for i := range fields {
			if fields[i] != goodfee.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := gfq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := gfq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := gfq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := gfq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (gfq *GoodFeeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(gfq.driver.Dialect())
	t1 := builder.Table(goodfee.Table)
	columns := gfq.fields
	if len(columns) == 0 {
		columns = goodfee.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if gfq.sql != nil {
		selector = gfq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range gfq.predicates {
		p(selector)
	}
	for _, p := range gfq.order {
		p(selector)
	}
	if offset := gfq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := gfq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// GoodFeeGroupBy is the group-by builder for GoodFee entities.
type GoodFeeGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (gfgb *GoodFeeGroupBy) Aggregate(fns ...AggregateFunc) *GoodFeeGroupBy {
	gfgb.fns = append(gfgb.fns, fns...)
	return gfgb
}

// Scan applies the group-by query and scans the result into the given value.
func (gfgb *GoodFeeGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := gfgb.path(ctx)
	if err != nil {
		return err
	}
	gfgb.sql = query
	return gfgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (gfgb *GoodFeeGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := gfgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (gfgb *GoodFeeGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(gfgb.fields) > 1 {
		return nil, errors.New("ent: GoodFeeGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := gfgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (gfgb *GoodFeeGroupBy) StringsX(ctx context.Context) []string {
	v, err := gfgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (gfgb *GoodFeeGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = gfgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goodfee.Label}
	default:
		err = fmt.Errorf("ent: GoodFeeGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (gfgb *GoodFeeGroupBy) StringX(ctx context.Context) string {
	v, err := gfgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (gfgb *GoodFeeGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(gfgb.fields) > 1 {
		return nil, errors.New("ent: GoodFeeGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := gfgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (gfgb *GoodFeeGroupBy) IntsX(ctx context.Context) []int {
	v, err := gfgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (gfgb *GoodFeeGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = gfgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goodfee.Label}
	default:
		err = fmt.Errorf("ent: GoodFeeGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (gfgb *GoodFeeGroupBy) IntX(ctx context.Context) int {
	v, err := gfgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (gfgb *GoodFeeGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(gfgb.fields) > 1 {
		return nil, errors.New("ent: GoodFeeGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := gfgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (gfgb *GoodFeeGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := gfgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (gfgb *GoodFeeGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = gfgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goodfee.Label}
	default:
		err = fmt.Errorf("ent: GoodFeeGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (gfgb *GoodFeeGroupBy) Float64X(ctx context.Context) float64 {
	v, err := gfgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (gfgb *GoodFeeGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(gfgb.fields) > 1 {
		return nil, errors.New("ent: GoodFeeGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := gfgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (gfgb *GoodFeeGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := gfgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (gfgb *GoodFeeGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = gfgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goodfee.Label}
	default:
		err = fmt.Errorf("ent: GoodFeeGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (gfgb *GoodFeeGroupBy) BoolX(ctx context.Context) bool {
	v, err := gfgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (gfgb *GoodFeeGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range gfgb.fields {
		if !goodfee.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := gfgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gfgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (gfgb *GoodFeeGroupBy) sqlQuery() *sql.Selector {
	selector := gfgb.sql.Select()
	aggregation := make([]string, 0, len(gfgb.fns))
	for _, fn := range gfgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(gfgb.fields)+len(gfgb.fns))
		for _, f := range gfgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(gfgb.fields...)...)
}

// GoodFeeSelect is the builder for selecting fields of GoodFee entities.
type GoodFeeSelect struct {
	*GoodFeeQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (gfs *GoodFeeSelect) Scan(ctx context.Context, v interface{}) error {
	if err := gfs.prepareQuery(ctx); err != nil {
		return err
	}
	gfs.sql = gfs.GoodFeeQuery.sqlQuery(ctx)
	return gfs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (gfs *GoodFeeSelect) ScanX(ctx context.Context, v interface{}) {
	if err := gfs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (gfs *GoodFeeSelect) Strings(ctx context.Context) ([]string, error) {
	if len(gfs.fields) > 1 {
		return nil, errors.New("ent: GoodFeeSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := gfs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (gfs *GoodFeeSelect) StringsX(ctx context.Context) []string {
	v, err := gfs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (gfs *GoodFeeSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = gfs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goodfee.Label}
	default:
		err = fmt.Errorf("ent: GoodFeeSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (gfs *GoodFeeSelect) StringX(ctx context.Context) string {
	v, err := gfs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (gfs *GoodFeeSelect) Ints(ctx context.Context) ([]int, error) {
	if len(gfs.fields) > 1 {
		return nil, errors.New("ent: GoodFeeSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := gfs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (gfs *GoodFeeSelect) IntsX(ctx context.Context) []int {
	v, err := gfs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (gfs *GoodFeeSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = gfs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goodfee.Label}
	default:
		err = fmt.Errorf("ent: GoodFeeSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (gfs *GoodFeeSelect) IntX(ctx context.Context) int {
	v, err := gfs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (gfs *GoodFeeSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(gfs.fields) > 1 {
		return nil, errors.New("ent: GoodFeeSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := gfs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (gfs *GoodFeeSelect) Float64sX(ctx context.Context) []float64 {
	v, err := gfs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (gfs *GoodFeeSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = gfs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goodfee.Label}
	default:
		err = fmt.Errorf("ent: GoodFeeSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (gfs *GoodFeeSelect) Float64X(ctx context.Context) float64 {
	v, err := gfs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (gfs *GoodFeeSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(gfs.fields) > 1 {
		return nil, errors.New("ent: GoodFeeSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := gfs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (gfs *GoodFeeSelect) BoolsX(ctx context.Context) []bool {
	v, err := gfs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (gfs *GoodFeeSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = gfs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goodfee.Label}
	default:
		err = fmt.Errorf("ent: GoodFeeSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (gfs *GoodFeeSelect) BoolX(ctx context.Context) bool {
	v, err := gfs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (gfs *GoodFeeSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := gfs.sql.Query()
	if err := gfs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
