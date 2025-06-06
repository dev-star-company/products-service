// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"products-service/internal/app/ent/predicate"
	"products-service/internal/app/ent/pricetype"
	"products-service/internal/app/ent/productprices"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PriceTypeQuery is the builder for querying PriceType entities.
type PriceTypeQuery struct {
	config
	ctx               *QueryContext
	order             []pricetype.OrderOption
	inters            []Interceptor
	predicates        []predicate.PriceType
	withProductPrices *ProductPricesQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PriceTypeQuery builder.
func (ptq *PriceTypeQuery) Where(ps ...predicate.PriceType) *PriceTypeQuery {
	ptq.predicates = append(ptq.predicates, ps...)
	return ptq
}

// Limit the number of records to be returned by this query.
func (ptq *PriceTypeQuery) Limit(limit int) *PriceTypeQuery {
	ptq.ctx.Limit = &limit
	return ptq
}

// Offset to start from.
func (ptq *PriceTypeQuery) Offset(offset int) *PriceTypeQuery {
	ptq.ctx.Offset = &offset
	return ptq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ptq *PriceTypeQuery) Unique(unique bool) *PriceTypeQuery {
	ptq.ctx.Unique = &unique
	return ptq
}

// Order specifies how the records should be ordered.
func (ptq *PriceTypeQuery) Order(o ...pricetype.OrderOption) *PriceTypeQuery {
	ptq.order = append(ptq.order, o...)
	return ptq
}

// QueryProductPrices chains the current query on the "product_prices" edge.
func (ptq *PriceTypeQuery) QueryProductPrices() *ProductPricesQuery {
	query := (&ProductPricesClient{config: ptq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ptq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ptq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(pricetype.Table, pricetype.FieldID, selector),
			sqlgraph.To(productprices.Table, productprices.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, pricetype.ProductPricesTable, pricetype.ProductPricesColumn),
		)
		fromU = sqlgraph.SetNeighbors(ptq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first PriceType entity from the query.
// Returns a *NotFoundError when no PriceType was found.
func (ptq *PriceTypeQuery) First(ctx context.Context) (*PriceType, error) {
	nodes, err := ptq.Limit(1).All(setContextOp(ctx, ptq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{pricetype.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ptq *PriceTypeQuery) FirstX(ctx context.Context) *PriceType {
	node, err := ptq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PriceType ID from the query.
// Returns a *NotFoundError when no PriceType ID was found.
func (ptq *PriceTypeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ptq.Limit(1).IDs(setContextOp(ctx, ptq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{pricetype.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ptq *PriceTypeQuery) FirstIDX(ctx context.Context) int {
	id, err := ptq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PriceType entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PriceType entity is found.
// Returns a *NotFoundError when no PriceType entities are found.
func (ptq *PriceTypeQuery) Only(ctx context.Context) (*PriceType, error) {
	nodes, err := ptq.Limit(2).All(setContextOp(ctx, ptq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{pricetype.Label}
	default:
		return nil, &NotSingularError{pricetype.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ptq *PriceTypeQuery) OnlyX(ctx context.Context) *PriceType {
	node, err := ptq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PriceType ID in the query.
// Returns a *NotSingularError when more than one PriceType ID is found.
// Returns a *NotFoundError when no entities are found.
func (ptq *PriceTypeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ptq.Limit(2).IDs(setContextOp(ctx, ptq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{pricetype.Label}
	default:
		err = &NotSingularError{pricetype.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ptq *PriceTypeQuery) OnlyIDX(ctx context.Context) int {
	id, err := ptq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PriceTypes.
func (ptq *PriceTypeQuery) All(ctx context.Context) ([]*PriceType, error) {
	ctx = setContextOp(ctx, ptq.ctx, ent.OpQueryAll)
	if err := ptq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*PriceType, *PriceTypeQuery]()
	return withInterceptors[[]*PriceType](ctx, ptq, qr, ptq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ptq *PriceTypeQuery) AllX(ctx context.Context) []*PriceType {
	nodes, err := ptq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PriceType IDs.
func (ptq *PriceTypeQuery) IDs(ctx context.Context) (ids []int, err error) {
	if ptq.ctx.Unique == nil && ptq.path != nil {
		ptq.Unique(true)
	}
	ctx = setContextOp(ctx, ptq.ctx, ent.OpQueryIDs)
	if err = ptq.Select(pricetype.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ptq *PriceTypeQuery) IDsX(ctx context.Context) []int {
	ids, err := ptq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ptq *PriceTypeQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ptq.ctx, ent.OpQueryCount)
	if err := ptq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ptq, querierCount[*PriceTypeQuery](), ptq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ptq *PriceTypeQuery) CountX(ctx context.Context) int {
	count, err := ptq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ptq *PriceTypeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ptq.ctx, ent.OpQueryExist)
	switch _, err := ptq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ptq *PriceTypeQuery) ExistX(ctx context.Context) bool {
	exist, err := ptq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PriceTypeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ptq *PriceTypeQuery) Clone() *PriceTypeQuery {
	if ptq == nil {
		return nil
	}
	return &PriceTypeQuery{
		config:            ptq.config,
		ctx:               ptq.ctx.Clone(),
		order:             append([]pricetype.OrderOption{}, ptq.order...),
		inters:            append([]Interceptor{}, ptq.inters...),
		predicates:        append([]predicate.PriceType{}, ptq.predicates...),
		withProductPrices: ptq.withProductPrices.Clone(),
		// clone intermediate query.
		sql:  ptq.sql.Clone(),
		path: ptq.path,
	}
}

// WithProductPrices tells the query-builder to eager-load the nodes that are connected to
// the "product_prices" edge. The optional arguments are used to configure the query builder of the edge.
func (ptq *PriceTypeQuery) WithProductPrices(opts ...func(*ProductPricesQuery)) *PriceTypeQuery {
	query := (&ProductPricesClient{config: ptq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ptq.withProductPrices = query
	return ptq
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
//	client.PriceType.Query().
//		GroupBy(pricetype.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ptq *PriceTypeQuery) GroupBy(field string, fields ...string) *PriceTypeGroupBy {
	ptq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PriceTypeGroupBy{build: ptq}
	grbuild.flds = &ptq.ctx.Fields
	grbuild.label = pricetype.Label
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
//	client.PriceType.Query().
//		Select(pricetype.FieldCreatedAt).
//		Scan(ctx, &v)
func (ptq *PriceTypeQuery) Select(fields ...string) *PriceTypeSelect {
	ptq.ctx.Fields = append(ptq.ctx.Fields, fields...)
	sbuild := &PriceTypeSelect{PriceTypeQuery: ptq}
	sbuild.label = pricetype.Label
	sbuild.flds, sbuild.scan = &ptq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PriceTypeSelect configured with the given aggregations.
func (ptq *PriceTypeQuery) Aggregate(fns ...AggregateFunc) *PriceTypeSelect {
	return ptq.Select().Aggregate(fns...)
}

func (ptq *PriceTypeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ptq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ptq); err != nil {
				return err
			}
		}
	}
	for _, f := range ptq.ctx.Fields {
		if !pricetype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ptq.path != nil {
		prev, err := ptq.path(ctx)
		if err != nil {
			return err
		}
		ptq.sql = prev
	}
	return nil
}

func (ptq *PriceTypeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PriceType, error) {
	var (
		nodes       = []*PriceType{}
		_spec       = ptq.querySpec()
		loadedTypes = [1]bool{
			ptq.withProductPrices != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*PriceType).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &PriceType{config: ptq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ptq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ptq.withProductPrices; query != nil {
		if err := ptq.loadProductPrices(ctx, query, nodes,
			func(n *PriceType) { n.Edges.ProductPrices = []*ProductPrices{} },
			func(n *PriceType, e *ProductPrices) { n.Edges.ProductPrices = append(n.Edges.ProductPrices, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ptq *PriceTypeQuery) loadProductPrices(ctx context.Context, query *ProductPricesQuery, nodes []*PriceType, init func(*PriceType), assign func(*PriceType, *ProductPrices)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*PriceType)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(productprices.FieldPriceTypeID)
	}
	query.Where(predicate.ProductPrices(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(pricetype.ProductPricesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.PriceTypeID
		if fk == nil {
			return fmt.Errorf(`foreign-key "price_type_id" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "price_type_id" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (ptq *PriceTypeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ptq.querySpec()
	_spec.Node.Columns = ptq.ctx.Fields
	if len(ptq.ctx.Fields) > 0 {
		_spec.Unique = ptq.ctx.Unique != nil && *ptq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ptq.driver, _spec)
}

func (ptq *PriceTypeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(pricetype.Table, pricetype.Columns, sqlgraph.NewFieldSpec(pricetype.FieldID, field.TypeInt))
	_spec.From = ptq.sql
	if unique := ptq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ptq.path != nil {
		_spec.Unique = true
	}
	if fields := ptq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, pricetype.FieldID)
		for i := range fields {
			if fields[i] != pricetype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ptq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ptq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ptq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ptq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ptq *PriceTypeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ptq.driver.Dialect())
	t1 := builder.Table(pricetype.Table)
	columns := ptq.ctx.Fields
	if len(columns) == 0 {
		columns = pricetype.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ptq.sql != nil {
		selector = ptq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ptq.ctx.Unique != nil && *ptq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ptq.predicates {
		p(selector)
	}
	for _, p := range ptq.order {
		p(selector)
	}
	if offset := ptq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ptq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PriceTypeGroupBy is the group-by builder for PriceType entities.
type PriceTypeGroupBy struct {
	selector
	build *PriceTypeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ptgb *PriceTypeGroupBy) Aggregate(fns ...AggregateFunc) *PriceTypeGroupBy {
	ptgb.fns = append(ptgb.fns, fns...)
	return ptgb
}

// Scan applies the selector query and scans the result into the given value.
func (ptgb *PriceTypeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ptgb.build.ctx, ent.OpQueryGroupBy)
	if err := ptgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PriceTypeQuery, *PriceTypeGroupBy](ctx, ptgb.build, ptgb, ptgb.build.inters, v)
}

func (ptgb *PriceTypeGroupBy) sqlScan(ctx context.Context, root *PriceTypeQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ptgb.fns))
	for _, fn := range ptgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ptgb.flds)+len(ptgb.fns))
		for _, f := range *ptgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ptgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ptgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PriceTypeSelect is the builder for selecting fields of PriceType entities.
type PriceTypeSelect struct {
	*PriceTypeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (pts *PriceTypeSelect) Aggregate(fns ...AggregateFunc) *PriceTypeSelect {
	pts.fns = append(pts.fns, fns...)
	return pts
}

// Scan applies the selector query and scans the result into the given value.
func (pts *PriceTypeSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pts.ctx, ent.OpQuerySelect)
	if err := pts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PriceTypeQuery, *PriceTypeSelect](ctx, pts.PriceTypeQuery, pts, pts.inters, v)
}

func (pts *PriceTypeSelect) sqlScan(ctx context.Context, root *PriceTypeQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(pts.fns))
	for _, fn := range pts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*pts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
