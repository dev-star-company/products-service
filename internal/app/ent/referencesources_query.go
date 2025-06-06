// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"products-service/internal/app/ent/predicate"
	"products-service/internal/app/ent/productreferences"
	"products-service/internal/app/ent/referencesources"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ReferenceSourcesQuery is the builder for querying ReferenceSources entities.
type ReferenceSourcesQuery struct {
	config
	ctx                   *QueryContext
	order                 []referencesources.OrderOption
	inters                []Interceptor
	predicates            []predicate.ReferenceSources
	withProductReferences *ProductReferencesQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ReferenceSourcesQuery builder.
func (rsq *ReferenceSourcesQuery) Where(ps ...predicate.ReferenceSources) *ReferenceSourcesQuery {
	rsq.predicates = append(rsq.predicates, ps...)
	return rsq
}

// Limit the number of records to be returned by this query.
func (rsq *ReferenceSourcesQuery) Limit(limit int) *ReferenceSourcesQuery {
	rsq.ctx.Limit = &limit
	return rsq
}

// Offset to start from.
func (rsq *ReferenceSourcesQuery) Offset(offset int) *ReferenceSourcesQuery {
	rsq.ctx.Offset = &offset
	return rsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rsq *ReferenceSourcesQuery) Unique(unique bool) *ReferenceSourcesQuery {
	rsq.ctx.Unique = &unique
	return rsq
}

// Order specifies how the records should be ordered.
func (rsq *ReferenceSourcesQuery) Order(o ...referencesources.OrderOption) *ReferenceSourcesQuery {
	rsq.order = append(rsq.order, o...)
	return rsq
}

// QueryProductReferences chains the current query on the "product_references" edge.
func (rsq *ReferenceSourcesQuery) QueryProductReferences() *ProductReferencesQuery {
	query := (&ProductReferencesClient{config: rsq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rsq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(referencesources.Table, referencesources.FieldID, selector),
			sqlgraph.To(productreferences.Table, productreferences.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, referencesources.ProductReferencesTable, referencesources.ProductReferencesColumn),
		)
		fromU = sqlgraph.SetNeighbors(rsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ReferenceSources entity from the query.
// Returns a *NotFoundError when no ReferenceSources was found.
func (rsq *ReferenceSourcesQuery) First(ctx context.Context) (*ReferenceSources, error) {
	nodes, err := rsq.Limit(1).All(setContextOp(ctx, rsq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{referencesources.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rsq *ReferenceSourcesQuery) FirstX(ctx context.Context) *ReferenceSources {
	node, err := rsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ReferenceSources ID from the query.
// Returns a *NotFoundError when no ReferenceSources ID was found.
func (rsq *ReferenceSourcesQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rsq.Limit(1).IDs(setContextOp(ctx, rsq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{referencesources.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rsq *ReferenceSourcesQuery) FirstIDX(ctx context.Context) int {
	id, err := rsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ReferenceSources entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ReferenceSources entity is found.
// Returns a *NotFoundError when no ReferenceSources entities are found.
func (rsq *ReferenceSourcesQuery) Only(ctx context.Context) (*ReferenceSources, error) {
	nodes, err := rsq.Limit(2).All(setContextOp(ctx, rsq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{referencesources.Label}
	default:
		return nil, &NotSingularError{referencesources.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rsq *ReferenceSourcesQuery) OnlyX(ctx context.Context) *ReferenceSources {
	node, err := rsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ReferenceSources ID in the query.
// Returns a *NotSingularError when more than one ReferenceSources ID is found.
// Returns a *NotFoundError when no entities are found.
func (rsq *ReferenceSourcesQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rsq.Limit(2).IDs(setContextOp(ctx, rsq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{referencesources.Label}
	default:
		err = &NotSingularError{referencesources.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rsq *ReferenceSourcesQuery) OnlyIDX(ctx context.Context) int {
	id, err := rsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ReferenceSourcesSlice.
func (rsq *ReferenceSourcesQuery) All(ctx context.Context) ([]*ReferenceSources, error) {
	ctx = setContextOp(ctx, rsq.ctx, ent.OpQueryAll)
	if err := rsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ReferenceSources, *ReferenceSourcesQuery]()
	return withInterceptors[[]*ReferenceSources](ctx, rsq, qr, rsq.inters)
}

// AllX is like All, but panics if an error occurs.
func (rsq *ReferenceSourcesQuery) AllX(ctx context.Context) []*ReferenceSources {
	nodes, err := rsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ReferenceSources IDs.
func (rsq *ReferenceSourcesQuery) IDs(ctx context.Context) (ids []int, err error) {
	if rsq.ctx.Unique == nil && rsq.path != nil {
		rsq.Unique(true)
	}
	ctx = setContextOp(ctx, rsq.ctx, ent.OpQueryIDs)
	if err = rsq.Select(referencesources.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rsq *ReferenceSourcesQuery) IDsX(ctx context.Context) []int {
	ids, err := rsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rsq *ReferenceSourcesQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, rsq.ctx, ent.OpQueryCount)
	if err := rsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, rsq, querierCount[*ReferenceSourcesQuery](), rsq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (rsq *ReferenceSourcesQuery) CountX(ctx context.Context) int {
	count, err := rsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rsq *ReferenceSourcesQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, rsq.ctx, ent.OpQueryExist)
	switch _, err := rsq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (rsq *ReferenceSourcesQuery) ExistX(ctx context.Context) bool {
	exist, err := rsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ReferenceSourcesQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rsq *ReferenceSourcesQuery) Clone() *ReferenceSourcesQuery {
	if rsq == nil {
		return nil
	}
	return &ReferenceSourcesQuery{
		config:                rsq.config,
		ctx:                   rsq.ctx.Clone(),
		order:                 append([]referencesources.OrderOption{}, rsq.order...),
		inters:                append([]Interceptor{}, rsq.inters...),
		predicates:            append([]predicate.ReferenceSources{}, rsq.predicates...),
		withProductReferences: rsq.withProductReferences.Clone(),
		// clone intermediate query.
		sql:  rsq.sql.Clone(),
		path: rsq.path,
	}
}

// WithProductReferences tells the query-builder to eager-load the nodes that are connected to
// the "product_references" edge. The optional arguments are used to configure the query builder of the edge.
func (rsq *ReferenceSourcesQuery) WithProductReferences(opts ...func(*ProductReferencesQuery)) *ReferenceSourcesQuery {
	query := (&ProductReferencesClient{config: rsq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rsq.withProductReferences = query
	return rsq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ReferenceSources.Query().
//		GroupBy(referencesources.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (rsq *ReferenceSourcesQuery) GroupBy(field string, fields ...string) *ReferenceSourcesGroupBy {
	rsq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ReferenceSourcesGroupBy{build: rsq}
	grbuild.flds = &rsq.ctx.Fields
	grbuild.label = referencesources.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.ReferenceSources.Query().
//		Select(referencesources.FieldCreatedAt).
//		Scan(ctx, &v)
func (rsq *ReferenceSourcesQuery) Select(fields ...string) *ReferenceSourcesSelect {
	rsq.ctx.Fields = append(rsq.ctx.Fields, fields...)
	sbuild := &ReferenceSourcesSelect{ReferenceSourcesQuery: rsq}
	sbuild.label = referencesources.Label
	sbuild.flds, sbuild.scan = &rsq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ReferenceSourcesSelect configured with the given aggregations.
func (rsq *ReferenceSourcesQuery) Aggregate(fns ...AggregateFunc) *ReferenceSourcesSelect {
	return rsq.Select().Aggregate(fns...)
}

func (rsq *ReferenceSourcesQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range rsq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, rsq); err != nil {
				return err
			}
		}
	}
	for _, f := range rsq.ctx.Fields {
		if !referencesources.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if rsq.path != nil {
		prev, err := rsq.path(ctx)
		if err != nil {
			return err
		}
		rsq.sql = prev
	}
	return nil
}

func (rsq *ReferenceSourcesQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ReferenceSources, error) {
	var (
		nodes       = []*ReferenceSources{}
		_spec       = rsq.querySpec()
		loadedTypes = [1]bool{
			rsq.withProductReferences != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ReferenceSources).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ReferenceSources{config: rsq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, rsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := rsq.withProductReferences; query != nil {
		if err := rsq.loadProductReferences(ctx, query, nodes,
			func(n *ReferenceSources) { n.Edges.ProductReferences = []*ProductReferences{} },
			func(n *ReferenceSources, e *ProductReferences) {
				n.Edges.ProductReferences = append(n.Edges.ProductReferences, e)
			}); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (rsq *ReferenceSourcesQuery) loadProductReferences(ctx context.Context, query *ProductReferencesQuery, nodes []*ReferenceSources, init func(*ReferenceSources), assign func(*ReferenceSources, *ProductReferences)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*ReferenceSources)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.ProductReferences(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(referencesources.ProductReferencesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.reference_sources_product_references
		if fk == nil {
			return fmt.Errorf(`foreign-key "reference_sources_product_references" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "reference_sources_product_references" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (rsq *ReferenceSourcesQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rsq.querySpec()
	_spec.Node.Columns = rsq.ctx.Fields
	if len(rsq.ctx.Fields) > 0 {
		_spec.Unique = rsq.ctx.Unique != nil && *rsq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, rsq.driver, _spec)
}

func (rsq *ReferenceSourcesQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(referencesources.Table, referencesources.Columns, sqlgraph.NewFieldSpec(referencesources.FieldID, field.TypeInt))
	_spec.From = rsq.sql
	if unique := rsq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if rsq.path != nil {
		_spec.Unique = true
	}
	if fields := rsq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, referencesources.FieldID)
		for i := range fields {
			if fields[i] != referencesources.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := rsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rsq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rsq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rsq *ReferenceSourcesQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rsq.driver.Dialect())
	t1 := builder.Table(referencesources.Table)
	columns := rsq.ctx.Fields
	if len(columns) == 0 {
		columns = referencesources.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rsq.sql != nil {
		selector = rsq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rsq.ctx.Unique != nil && *rsq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range rsq.predicates {
		p(selector)
	}
	for _, p := range rsq.order {
		p(selector)
	}
	if offset := rsq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rsq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ReferenceSourcesGroupBy is the group-by builder for ReferenceSources entities.
type ReferenceSourcesGroupBy struct {
	selector
	build *ReferenceSourcesQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rsgb *ReferenceSourcesGroupBy) Aggregate(fns ...AggregateFunc) *ReferenceSourcesGroupBy {
	rsgb.fns = append(rsgb.fns, fns...)
	return rsgb
}

// Scan applies the selector query and scans the result into the given value.
func (rsgb *ReferenceSourcesGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rsgb.build.ctx, ent.OpQueryGroupBy)
	if err := rsgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ReferenceSourcesQuery, *ReferenceSourcesGroupBy](ctx, rsgb.build, rsgb, rsgb.build.inters, v)
}

func (rsgb *ReferenceSourcesGroupBy) sqlScan(ctx context.Context, root *ReferenceSourcesQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(rsgb.fns))
	for _, fn := range rsgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*rsgb.flds)+len(rsgb.fns))
		for _, f := range *rsgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*rsgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rsgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ReferenceSourcesSelect is the builder for selecting fields of ReferenceSources entities.
type ReferenceSourcesSelect struct {
	*ReferenceSourcesQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rss *ReferenceSourcesSelect) Aggregate(fns ...AggregateFunc) *ReferenceSourcesSelect {
	rss.fns = append(rss.fns, fns...)
	return rss
}

// Scan applies the selector query and scans the result into the given value.
func (rss *ReferenceSourcesSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rss.ctx, ent.OpQuerySelect)
	if err := rss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ReferenceSourcesQuery, *ReferenceSourcesSelect](ctx, rss.ReferenceSourcesQuery, rss, rss.inters, v)
}

func (rss *ReferenceSourcesSelect) sqlScan(ctx context.Context, root *ReferenceSourcesQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(rss.fns))
	for _, fn := range rss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*rss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
