// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/morning-night-dream/platform/pkg/ent/article"
	"github.com/morning-night-dream/platform/pkg/ent/predicate"
	"github.com/morning-night-dream/platform/pkg/ent/readarticle"
)

// ReadArticleQuery is the builder for querying ReadArticle entities.
type ReadArticleQuery struct {
	config
	limit       *int
	offset      *int
	unique      *bool
	order       []OrderFunc
	fields      []string
	inters      []Interceptor
	predicates  []predicate.ReadArticle
	withArticle *ArticleQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ReadArticleQuery builder.
func (raq *ReadArticleQuery) Where(ps ...predicate.ReadArticle) *ReadArticleQuery {
	raq.predicates = append(raq.predicates, ps...)
	return raq
}

// Limit the number of records to be returned by this query.
func (raq *ReadArticleQuery) Limit(limit int) *ReadArticleQuery {
	raq.limit = &limit
	return raq
}

// Offset to start from.
func (raq *ReadArticleQuery) Offset(offset int) *ReadArticleQuery {
	raq.offset = &offset
	return raq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (raq *ReadArticleQuery) Unique(unique bool) *ReadArticleQuery {
	raq.unique = &unique
	return raq
}

// Order specifies how the records should be ordered.
func (raq *ReadArticleQuery) Order(o ...OrderFunc) *ReadArticleQuery {
	raq.order = append(raq.order, o...)
	return raq
}

// QueryArticle chains the current query on the "article" edge.
func (raq *ReadArticleQuery) QueryArticle() *ArticleQuery {
	query := (&ArticleClient{config: raq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := raq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := raq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(readarticle.Table, readarticle.FieldID, selector),
			sqlgraph.To(article.Table, article.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, readarticle.ArticleTable, readarticle.ArticleColumn),
		)
		fromU = sqlgraph.SetNeighbors(raq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ReadArticle entity from the query.
// Returns a *NotFoundError when no ReadArticle was found.
func (raq *ReadArticleQuery) First(ctx context.Context) (*ReadArticle, error) {
	nodes, err := raq.Limit(1).All(newQueryContext(ctx, TypeReadArticle, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{readarticle.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (raq *ReadArticleQuery) FirstX(ctx context.Context) *ReadArticle {
	node, err := raq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ReadArticle ID from the query.
// Returns a *NotFoundError when no ReadArticle ID was found.
func (raq *ReadArticleQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = raq.Limit(1).IDs(newQueryContext(ctx, TypeReadArticle, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{readarticle.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (raq *ReadArticleQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := raq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ReadArticle entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ReadArticle entity is found.
// Returns a *NotFoundError when no ReadArticle entities are found.
func (raq *ReadArticleQuery) Only(ctx context.Context) (*ReadArticle, error) {
	nodes, err := raq.Limit(2).All(newQueryContext(ctx, TypeReadArticle, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{readarticle.Label}
	default:
		return nil, &NotSingularError{readarticle.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (raq *ReadArticleQuery) OnlyX(ctx context.Context) *ReadArticle {
	node, err := raq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ReadArticle ID in the query.
// Returns a *NotSingularError when more than one ReadArticle ID is found.
// Returns a *NotFoundError when no entities are found.
func (raq *ReadArticleQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = raq.Limit(2).IDs(newQueryContext(ctx, TypeReadArticle, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{readarticle.Label}
	default:
		err = &NotSingularError{readarticle.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (raq *ReadArticleQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := raq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ReadArticles.
func (raq *ReadArticleQuery) All(ctx context.Context) ([]*ReadArticle, error) {
	ctx = newQueryContext(ctx, TypeReadArticle, "All")
	if err := raq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ReadArticle, *ReadArticleQuery]()
	return withInterceptors[[]*ReadArticle](ctx, raq, qr, raq.inters)
}

// AllX is like All, but panics if an error occurs.
func (raq *ReadArticleQuery) AllX(ctx context.Context) []*ReadArticle {
	nodes, err := raq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ReadArticle IDs.
func (raq *ReadArticleQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	ctx = newQueryContext(ctx, TypeReadArticle, "IDs")
	if err := raq.Select(readarticle.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (raq *ReadArticleQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := raq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (raq *ReadArticleQuery) Count(ctx context.Context) (int, error) {
	ctx = newQueryContext(ctx, TypeReadArticle, "Count")
	if err := raq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, raq, querierCount[*ReadArticleQuery](), raq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (raq *ReadArticleQuery) CountX(ctx context.Context) int {
	count, err := raq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (raq *ReadArticleQuery) Exist(ctx context.Context) (bool, error) {
	ctx = newQueryContext(ctx, TypeReadArticle, "Exist")
	switch _, err := raq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (raq *ReadArticleQuery) ExistX(ctx context.Context) bool {
	exist, err := raq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ReadArticleQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (raq *ReadArticleQuery) Clone() *ReadArticleQuery {
	if raq == nil {
		return nil
	}
	return &ReadArticleQuery{
		config:      raq.config,
		limit:       raq.limit,
		offset:      raq.offset,
		order:       append([]OrderFunc{}, raq.order...),
		inters:      append([]Interceptor{}, raq.inters...),
		predicates:  append([]predicate.ReadArticle{}, raq.predicates...),
		withArticle: raq.withArticle.Clone(),
		// clone intermediate query.
		sql:    raq.sql.Clone(),
		path:   raq.path,
		unique: raq.unique,
	}
}

// WithArticle tells the query-builder to eager-load the nodes that are connected to
// the "article" edge. The optional arguments are used to configure the query builder of the edge.
func (raq *ReadArticleQuery) WithArticle(opts ...func(*ArticleQuery)) *ReadArticleQuery {
	query := (&ArticleClient{config: raq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	raq.withArticle = query
	return raq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ArticleID uuid.UUID `json:"article_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ReadArticle.Query().
//		GroupBy(readarticle.FieldArticleID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (raq *ReadArticleQuery) GroupBy(field string, fields ...string) *ReadArticleGroupBy {
	raq.fields = append([]string{field}, fields...)
	grbuild := &ReadArticleGroupBy{build: raq}
	grbuild.flds = &raq.fields
	grbuild.label = readarticle.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ArticleID uuid.UUID `json:"article_id,omitempty"`
//	}
//
//	client.ReadArticle.Query().
//		Select(readarticle.FieldArticleID).
//		Scan(ctx, &v)
func (raq *ReadArticleQuery) Select(fields ...string) *ReadArticleSelect {
	raq.fields = append(raq.fields, fields...)
	sbuild := &ReadArticleSelect{ReadArticleQuery: raq}
	sbuild.label = readarticle.Label
	sbuild.flds, sbuild.scan = &raq.fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ReadArticleSelect configured with the given aggregations.
func (raq *ReadArticleQuery) Aggregate(fns ...AggregateFunc) *ReadArticleSelect {
	return raq.Select().Aggregate(fns...)
}

func (raq *ReadArticleQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range raq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, raq); err != nil {
				return err
			}
		}
	}
	for _, f := range raq.fields {
		if !readarticle.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if raq.path != nil {
		prev, err := raq.path(ctx)
		if err != nil {
			return err
		}
		raq.sql = prev
	}
	return nil
}

func (raq *ReadArticleQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ReadArticle, error) {
	var (
		nodes       = []*ReadArticle{}
		_spec       = raq.querySpec()
		loadedTypes = [1]bool{
			raq.withArticle != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ReadArticle).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ReadArticle{config: raq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, raq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := raq.withArticle; query != nil {
		if err := raq.loadArticle(ctx, query, nodes, nil,
			func(n *ReadArticle, e *Article) { n.Edges.Article = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (raq *ReadArticleQuery) loadArticle(ctx context.Context, query *ArticleQuery, nodes []*ReadArticle, init func(*ReadArticle), assign func(*ReadArticle, *Article)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*ReadArticle)
	for i := range nodes {
		fk := nodes[i].ArticleID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(article.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "article_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (raq *ReadArticleQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := raq.querySpec()
	_spec.Node.Columns = raq.fields
	if len(raq.fields) > 0 {
		_spec.Unique = raq.unique != nil && *raq.unique
	}
	return sqlgraph.CountNodes(ctx, raq.driver, _spec)
}

func (raq *ReadArticleQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   readarticle.Table,
			Columns: readarticle.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: readarticle.FieldID,
			},
		},
		From:   raq.sql,
		Unique: true,
	}
	if unique := raq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := raq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, readarticle.FieldID)
		for i := range fields {
			if fields[i] != readarticle.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := raq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := raq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := raq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := raq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (raq *ReadArticleQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(raq.driver.Dialect())
	t1 := builder.Table(readarticle.Table)
	columns := raq.fields
	if len(columns) == 0 {
		columns = readarticle.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if raq.sql != nil {
		selector = raq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if raq.unique != nil && *raq.unique {
		selector.Distinct()
	}
	for _, p := range raq.predicates {
		p(selector)
	}
	for _, p := range raq.order {
		p(selector)
	}
	if offset := raq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := raq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ReadArticleGroupBy is the group-by builder for ReadArticle entities.
type ReadArticleGroupBy struct {
	selector
	build *ReadArticleQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ragb *ReadArticleGroupBy) Aggregate(fns ...AggregateFunc) *ReadArticleGroupBy {
	ragb.fns = append(ragb.fns, fns...)
	return ragb
}

// Scan applies the selector query and scans the result into the given value.
func (ragb *ReadArticleGroupBy) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeReadArticle, "GroupBy")
	if err := ragb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ReadArticleQuery, *ReadArticleGroupBy](ctx, ragb.build, ragb, ragb.build.inters, v)
}

func (ragb *ReadArticleGroupBy) sqlScan(ctx context.Context, root *ReadArticleQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ragb.fns))
	for _, fn := range ragb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ragb.flds)+len(ragb.fns))
		for _, f := range *ragb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ragb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ragb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ReadArticleSelect is the builder for selecting fields of ReadArticle entities.
type ReadArticleSelect struct {
	*ReadArticleQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ras *ReadArticleSelect) Aggregate(fns ...AggregateFunc) *ReadArticleSelect {
	ras.fns = append(ras.fns, fns...)
	return ras
}

// Scan applies the selector query and scans the result into the given value.
func (ras *ReadArticleSelect) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeReadArticle, "Select")
	if err := ras.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ReadArticleQuery, *ReadArticleSelect](ctx, ras.ReadArticleQuery, ras, ras.inters, v)
}

func (ras *ReadArticleSelect) sqlScan(ctx context.Context, root *ReadArticleQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ras.fns))
	for _, fn := range ras.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ras.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ras.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
