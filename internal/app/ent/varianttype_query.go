// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"products-service/internal/app/ent/predicate"
	"products-service/internal/app/ent/products"
	"products-service/internal/app/ent/varianttype"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// VariantTypeQuery is the builder for querying VariantType entities.
type VariantTypeQuery struct {
	config
	ctx          *QueryContext
	order        []varianttype.OrderOption
	inters       []Interceptor
	predicates   []predicate.VariantType
	withProducts *ProductsQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the VariantTypeQuery builder.
func (vtq *VariantTypeQuery) Where(ps ...predicate.VariantType) *VariantTypeQuery {
	vtq.predicates = append(vtq.predicates, ps...)
	return vtq
}

// Limit the number of records to be returned by this query.
func (vtq *VariantTypeQuery) Limit(limit int) *VariantTypeQuery {
	vtq.ctx.Limit = &limit
	return vtq
}

// Offset to start from.
func (vtq *VariantTypeQuery) Offset(offset int) *VariantTypeQuery {
	vtq.ctx.Offset = &offset
	return vtq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (vtq *VariantTypeQuery) Unique(unique bool) *VariantTypeQuery {
	vtq.ctx.Unique = &unique
	return vtq
}

// Order specifies how the records should be ordered.
func (vtq *VariantTypeQuery) Order(o ...varianttype.OrderOption) *VariantTypeQuery {
	vtq.order = append(vtq.order, o...)
	return vtq
}

// QueryProducts chains the current query on the "products" edge.
func (vtq *VariantTypeQuery) QueryProducts() *ProductsQuery {
	query := (&ProductsClient{config: vtq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vtq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vtq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(varianttype.Table, varianttype.FieldID, selector),
			sqlgraph.To(products.Table, products.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, varianttype.ProductsTable, varianttype.ProductsColumn),
		)
		fromU = sqlgraph.SetNeighbors(vtq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first VariantType entity from the query.
// Returns a *NotFoundError when no VariantType was found.
func (vtq *VariantTypeQuery) First(ctx context.Context) (*VariantType, error) {
	nodes, err := vtq.Limit(1).All(setContextOp(ctx, vtq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{varianttype.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (vtq *VariantTypeQuery) FirstX(ctx context.Context) *VariantType {
	node, err := vtq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first VariantType ID from the query.
// Returns a *NotFoundError when no VariantType ID was found.
func (vtq *VariantTypeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = vtq.Limit(1).IDs(setContextOp(ctx, vtq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{varianttype.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (vtq *VariantTypeQuery) FirstIDX(ctx context.Context) int {
	id, err := vtq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single VariantType entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one VariantType entity is found.
// Returns a *NotFoundError when no VariantType entities are found.
func (vtq *VariantTypeQuery) Only(ctx context.Context) (*VariantType, error) {
	nodes, err := vtq.Limit(2).All(setContextOp(ctx, vtq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{varianttype.Label}
	default:
		return nil, &NotSingularError{varianttype.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (vtq *VariantTypeQuery) OnlyX(ctx context.Context) *VariantType {
	node, err := vtq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only VariantType ID in the query.
// Returns a *NotSingularError when more than one VariantType ID is found.
// Returns a *NotFoundError when no entities are found.
func (vtq *VariantTypeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = vtq.Limit(2).IDs(setContextOp(ctx, vtq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{varianttype.Label}
	default:
		err = &NotSingularError{varianttype.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (vtq *VariantTypeQuery) OnlyIDX(ctx context.Context) int {
	id, err := vtq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of VariantTypes.
func (vtq *VariantTypeQuery) All(ctx context.Context) ([]*VariantType, error) {
	ctx = setContextOp(ctx, vtq.ctx, ent.OpQueryAll)
	if err := vtq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*VariantType, *VariantTypeQuery]()
	return withInterceptors[[]*VariantType](ctx, vtq, qr, vtq.inters)
}

// AllX is like All, but panics if an error occurs.
func (vtq *VariantTypeQuery) AllX(ctx context.Context) []*VariantType {
	nodes, err := vtq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of VariantType IDs.
func (vtq *VariantTypeQuery) IDs(ctx context.Context) (ids []int, err error) {
	if vtq.ctx.Unique == nil && vtq.path != nil {
		vtq.Unique(true)
	}
	ctx = setContextOp(ctx, vtq.ctx, ent.OpQueryIDs)
	if err = vtq.Select(varianttype.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (vtq *VariantTypeQuery) IDsX(ctx context.Context) []int {
	ids, err := vtq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (vtq *VariantTypeQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, vtq.ctx, ent.OpQueryCount)
	if err := vtq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, vtq, querierCount[*VariantTypeQuery](), vtq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (vtq *VariantTypeQuery) CountX(ctx context.Context) int {
	count, err := vtq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (vtq *VariantTypeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, vtq.ctx, ent.OpQueryExist)
	switch _, err := vtq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (vtq *VariantTypeQuery) ExistX(ctx context.Context) bool {
	exist, err := vtq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the VariantTypeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (vtq *VariantTypeQuery) Clone() *VariantTypeQuery {
	if vtq == nil {
		return nil
	}
	return &VariantTypeQuery{
		config:       vtq.config,
		ctx:          vtq.ctx.Clone(),
		order:        append([]varianttype.OrderOption{}, vtq.order...),
		inters:       append([]Interceptor{}, vtq.inters...),
		predicates:   append([]predicate.VariantType{}, vtq.predicates...),
		withProducts: vtq.withProducts.Clone(),
		// clone intermediate query.
		sql:  vtq.sql.Clone(),
		path: vtq.path,
	}
}

// WithProducts tells the query-builder to eager-load the nodes that are connected to
// the "products" edge. The optional arguments are used to configure the query builder of the edge.
func (vtq *VariantTypeQuery) WithProducts(opts ...func(*ProductsQuery)) *VariantTypeQuery {
	query := (&ProductsClient{config: vtq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	vtq.withProducts = query
	return vtq
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
//	client.VariantType.Query().
//		GroupBy(varianttype.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (vtq *VariantTypeQuery) GroupBy(field string, fields ...string) *VariantTypeGroupBy {
	vtq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &VariantTypeGroupBy{build: vtq}
	grbuild.flds = &vtq.ctx.Fields
	grbuild.label = varianttype.Label
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
//	client.VariantType.Query().
//		Select(varianttype.FieldCreatedAt).
//		Scan(ctx, &v)
func (vtq *VariantTypeQuery) Select(fields ...string) *VariantTypeSelect {
	vtq.ctx.Fields = append(vtq.ctx.Fields, fields...)
	sbuild := &VariantTypeSelect{VariantTypeQuery: vtq}
	sbuild.label = varianttype.Label
	sbuild.flds, sbuild.scan = &vtq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a VariantTypeSelect configured with the given aggregations.
func (vtq *VariantTypeQuery) Aggregate(fns ...AggregateFunc) *VariantTypeSelect {
	return vtq.Select().Aggregate(fns...)
}

func (vtq *VariantTypeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range vtq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, vtq); err != nil {
				return err
			}
		}
	}
	for _, f := range vtq.ctx.Fields {
		if !varianttype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if vtq.path != nil {
		prev, err := vtq.path(ctx)
		if err != nil {
			return err
		}
		vtq.sql = prev
	}
	return nil
}

func (vtq *VariantTypeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*VariantType, error) {
	var (
		nodes       = []*VariantType{}
		_spec       = vtq.querySpec()
		loadedTypes = [1]bool{
			vtq.withProducts != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*VariantType).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &VariantType{config: vtq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, vtq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := vtq.withProducts; query != nil {
		if err := vtq.loadProducts(ctx, query, nodes,
			func(n *VariantType) { n.Edges.Products = []*Products{} },
			func(n *VariantType, e *Products) { n.Edges.Products = append(n.Edges.Products, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (vtq *VariantTypeQuery) loadProducts(ctx context.Context, query *ProductsQuery, nodes []*VariantType, init func(*VariantType), assign func(*VariantType, *Products)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*VariantType)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Products(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(varianttype.ProductsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.variant_type_products
		if fk == nil {
			return fmt.Errorf(`foreign-key "variant_type_products" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "variant_type_products" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (vtq *VariantTypeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := vtq.querySpec()
	_spec.Node.Columns = vtq.ctx.Fields
	if len(vtq.ctx.Fields) > 0 {
		_spec.Unique = vtq.ctx.Unique != nil && *vtq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, vtq.driver, _spec)
}

func (vtq *VariantTypeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(varianttype.Table, varianttype.Columns, sqlgraph.NewFieldSpec(varianttype.FieldID, field.TypeInt))
	_spec.From = vtq.sql
	if unique := vtq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if vtq.path != nil {
		_spec.Unique = true
	}
	if fields := vtq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, varianttype.FieldID)
		for i := range fields {
			if fields[i] != varianttype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := vtq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := vtq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := vtq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := vtq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (vtq *VariantTypeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(vtq.driver.Dialect())
	t1 := builder.Table(varianttype.Table)
	columns := vtq.ctx.Fields
	if len(columns) == 0 {
		columns = varianttype.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if vtq.sql != nil {
		selector = vtq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if vtq.ctx.Unique != nil && *vtq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range vtq.predicates {
		p(selector)
	}
	for _, p := range vtq.order {
		p(selector)
	}
	if offset := vtq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := vtq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// VariantTypeGroupBy is the group-by builder for VariantType entities.
type VariantTypeGroupBy struct {
	selector
	build *VariantTypeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (vtgb *VariantTypeGroupBy) Aggregate(fns ...AggregateFunc) *VariantTypeGroupBy {
	vtgb.fns = append(vtgb.fns, fns...)
	return vtgb
}

// Scan applies the selector query and scans the result into the given value.
func (vtgb *VariantTypeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vtgb.build.ctx, ent.OpQueryGroupBy)
	if err := vtgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VariantTypeQuery, *VariantTypeGroupBy](ctx, vtgb.build, vtgb, vtgb.build.inters, v)
}

func (vtgb *VariantTypeGroupBy) sqlScan(ctx context.Context, root *VariantTypeQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(vtgb.fns))
	for _, fn := range vtgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*vtgb.flds)+len(vtgb.fns))
		for _, f := range *vtgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*vtgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vtgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// VariantTypeSelect is the builder for selecting fields of VariantType entities.
type VariantTypeSelect struct {
	*VariantTypeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (vts *VariantTypeSelect) Aggregate(fns ...AggregateFunc) *VariantTypeSelect {
	vts.fns = append(vts.fns, fns...)
	return vts
}

// Scan applies the selector query and scans the result into the given value.
func (vts *VariantTypeSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vts.ctx, ent.OpQuerySelect)
	if err := vts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VariantTypeQuery, *VariantTypeSelect](ctx, vts.VariantTypeQuery, vts, vts.inters, v)
}

func (vts *VariantTypeSelect) sqlScan(ctx context.Context, root *VariantTypeQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(vts.fns))
	for _, fn := range vts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*vts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
