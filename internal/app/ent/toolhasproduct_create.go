// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"products-service/internal/app/ent/products"
	"products-service/internal/app/ent/toolhasproduct"
	"products-service/internal/app/ent/tools"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ToolHasProductCreate is the builder for creating a ToolHasProduct entity.
type ToolHasProductCreate struct {
	config
	mutation *ToolHasProductMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (thpc *ToolHasProductCreate) SetCreatedAt(t time.Time) *ToolHasProductCreate {
	thpc.mutation.SetCreatedAt(t)
	return thpc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (thpc *ToolHasProductCreate) SetNillableCreatedAt(t *time.Time) *ToolHasProductCreate {
	if t != nil {
		thpc.SetCreatedAt(*t)
	}
	return thpc
}

// SetUpdatedAt sets the "updated_at" field.
func (thpc *ToolHasProductCreate) SetUpdatedAt(t time.Time) *ToolHasProductCreate {
	thpc.mutation.SetUpdatedAt(t)
	return thpc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (thpc *ToolHasProductCreate) SetNillableUpdatedAt(t *time.Time) *ToolHasProductCreate {
	if t != nil {
		thpc.SetUpdatedAt(*t)
	}
	return thpc
}

// SetDeletedAt sets the "deleted_at" field.
func (thpc *ToolHasProductCreate) SetDeletedAt(t time.Time) *ToolHasProductCreate {
	thpc.mutation.SetDeletedAt(t)
	return thpc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (thpc *ToolHasProductCreate) SetNillableDeletedAt(t *time.Time) *ToolHasProductCreate {
	if t != nil {
		thpc.SetDeletedAt(*t)
	}
	return thpc
}

// SetCreatedBy sets the "created_by" field.
func (thpc *ToolHasProductCreate) SetCreatedBy(i int) *ToolHasProductCreate {
	thpc.mutation.SetCreatedBy(i)
	return thpc
}

// SetUpdatedBy sets the "updated_by" field.
func (thpc *ToolHasProductCreate) SetUpdatedBy(i int) *ToolHasProductCreate {
	thpc.mutation.SetUpdatedBy(i)
	return thpc
}

// SetDeletedBy sets the "deleted_by" field.
func (thpc *ToolHasProductCreate) SetDeletedBy(i int) *ToolHasProductCreate {
	thpc.mutation.SetDeletedBy(i)
	return thpc
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (thpc *ToolHasProductCreate) SetNillableDeletedBy(i *int) *ToolHasProductCreate {
	if i != nil {
		thpc.SetDeletedBy(*i)
	}
	return thpc
}

// SetToolID sets the "tool_id" field.
func (thpc *ToolHasProductCreate) SetToolID(i int) *ToolHasProductCreate {
	thpc.mutation.SetToolID(i)
	return thpc
}

// SetProductID sets the "product_id" field.
func (thpc *ToolHasProductCreate) SetProductID(i int) *ToolHasProductCreate {
	thpc.mutation.SetProductID(i)
	return thpc
}

// SetNillableProductID sets the "product_id" field if the given value is not nil.
func (thpc *ToolHasProductCreate) SetNillableProductID(i *int) *ToolHasProductCreate {
	if i != nil {
		thpc.SetProductID(*i)
	}
	return thpc
}

// SetProduct sets the "product" edge to the Products entity.
func (thpc *ToolHasProductCreate) SetProduct(p *Products) *ToolHasProductCreate {
	return thpc.SetProductID(p.ID)
}

// SetTool sets the "tool" edge to the Tools entity.
func (thpc *ToolHasProductCreate) SetTool(t *Tools) *ToolHasProductCreate {
	return thpc.SetToolID(t.ID)
}

// Mutation returns the ToolHasProductMutation object of the builder.
func (thpc *ToolHasProductCreate) Mutation() *ToolHasProductMutation {
	return thpc.mutation
}

// Save creates the ToolHasProduct in the database.
func (thpc *ToolHasProductCreate) Save(ctx context.Context) (*ToolHasProduct, error) {
	thpc.defaults()
	return withHooks(ctx, thpc.sqlSave, thpc.mutation, thpc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (thpc *ToolHasProductCreate) SaveX(ctx context.Context) *ToolHasProduct {
	v, err := thpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (thpc *ToolHasProductCreate) Exec(ctx context.Context) error {
	_, err := thpc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (thpc *ToolHasProductCreate) ExecX(ctx context.Context) {
	if err := thpc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (thpc *ToolHasProductCreate) defaults() {
	if _, ok := thpc.mutation.CreatedAt(); !ok {
		v := toolhasproduct.DefaultCreatedAt()
		thpc.mutation.SetCreatedAt(v)
	}
	if _, ok := thpc.mutation.UpdatedAt(); !ok {
		v := toolhasproduct.DefaultUpdatedAt()
		thpc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (thpc *ToolHasProductCreate) check() error {
	if _, ok := thpc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "ToolHasProduct.created_at"`)}
	}
	if _, ok := thpc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "ToolHasProduct.updated_at"`)}
	}
	if _, ok := thpc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "ToolHasProduct.created_by"`)}
	}
	if v, ok := thpc.mutation.CreatedBy(); ok {
		if err := toolhasproduct.CreatedByValidator(v); err != nil {
			return &ValidationError{Name: "created_by", err: fmt.Errorf(`ent: validator failed for field "ToolHasProduct.created_by": %w`, err)}
		}
	}
	if _, ok := thpc.mutation.UpdatedBy(); !ok {
		return &ValidationError{Name: "updated_by", err: errors.New(`ent: missing required field "ToolHasProduct.updated_by"`)}
	}
	if v, ok := thpc.mutation.UpdatedBy(); ok {
		if err := toolhasproduct.UpdatedByValidator(v); err != nil {
			return &ValidationError{Name: "updated_by", err: fmt.Errorf(`ent: validator failed for field "ToolHasProduct.updated_by": %w`, err)}
		}
	}
	if _, ok := thpc.mutation.ToolID(); !ok {
		return &ValidationError{Name: "tool_id", err: errors.New(`ent: missing required field "ToolHasProduct.tool_id"`)}
	}
	if len(thpc.mutation.ProductIDs()) == 0 {
		return &ValidationError{Name: "product", err: errors.New(`ent: missing required edge "ToolHasProduct.product"`)}
	}
	if len(thpc.mutation.ToolIDs()) == 0 {
		return &ValidationError{Name: "tool", err: errors.New(`ent: missing required edge "ToolHasProduct.tool"`)}
	}
	return nil
}

func (thpc *ToolHasProductCreate) sqlSave(ctx context.Context) (*ToolHasProduct, error) {
	if err := thpc.check(); err != nil {
		return nil, err
	}
	_node, _spec := thpc.createSpec()
	if err := sqlgraph.CreateNode(ctx, thpc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	thpc.mutation.id = &_node.ID
	thpc.mutation.done = true
	return _node, nil
}

func (thpc *ToolHasProductCreate) createSpec() (*ToolHasProduct, *sqlgraph.CreateSpec) {
	var (
		_node = &ToolHasProduct{config: thpc.config}
		_spec = sqlgraph.NewCreateSpec(toolhasproduct.Table, sqlgraph.NewFieldSpec(toolhasproduct.FieldID, field.TypeInt))
	)
	if value, ok := thpc.mutation.CreatedAt(); ok {
		_spec.SetField(toolhasproduct.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := thpc.mutation.UpdatedAt(); ok {
		_spec.SetField(toolhasproduct.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := thpc.mutation.DeletedAt(); ok {
		_spec.SetField(toolhasproduct.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if value, ok := thpc.mutation.CreatedBy(); ok {
		_spec.SetField(toolhasproduct.FieldCreatedBy, field.TypeInt, value)
		_node.CreatedBy = value
	}
	if value, ok := thpc.mutation.UpdatedBy(); ok {
		_spec.SetField(toolhasproduct.FieldUpdatedBy, field.TypeInt, value)
		_node.UpdatedBy = value
	}
	if value, ok := thpc.mutation.DeletedBy(); ok {
		_spec.SetField(toolhasproduct.FieldDeletedBy, field.TypeInt, value)
		_node.DeletedBy = &value
	}
	if value, ok := thpc.mutation.ToolID(); ok {
		_spec.SetField(toolhasproduct.FieldToolID, field.TypeInt, value)
		_node.ToolID = &value
	}
	if value, ok := thpc.mutation.ProductID(); ok {
		_spec.SetField(toolhasproduct.FieldProductID, field.TypeInt, value)
		_node.ProductID = &value
	}
	if nodes := thpc.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   toolhasproduct.ProductTable,
			Columns: []string{toolhasproduct.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(products.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.products_tool_has_product = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := thpc.mutation.ToolIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   toolhasproduct.ToolTable,
			Columns: []string{toolhasproduct.ToolColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tools.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.tools_tool_has_product = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ToolHasProductCreateBulk is the builder for creating many ToolHasProduct entities in bulk.
type ToolHasProductCreateBulk struct {
	config
	err      error
	builders []*ToolHasProductCreate
}

// Save creates the ToolHasProduct entities in the database.
func (thpcb *ToolHasProductCreateBulk) Save(ctx context.Context) ([]*ToolHasProduct, error) {
	if thpcb.err != nil {
		return nil, thpcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(thpcb.builders))
	nodes := make([]*ToolHasProduct, len(thpcb.builders))
	mutators := make([]Mutator, len(thpcb.builders))
	for i := range thpcb.builders {
		func(i int, root context.Context) {
			builder := thpcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ToolHasProductMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, thpcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, thpcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, thpcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (thpcb *ToolHasProductCreateBulk) SaveX(ctx context.Context) []*ToolHasProduct {
	v, err := thpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (thpcb *ToolHasProductCreateBulk) Exec(ctx context.Context) error {
	_, err := thpcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (thpcb *ToolHasProductCreateBulk) ExecX(ctx context.Context) {
	if err := thpcb.Exec(ctx); err != nil {
		panic(err)
	}
}
