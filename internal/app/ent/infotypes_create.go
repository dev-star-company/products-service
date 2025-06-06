// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"products-service/internal/app/ent/infotypes"
	"products-service/internal/app/ent/productinfo"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// InfoTypesCreate is the builder for creating a InfoTypes entity.
type InfoTypesCreate struct {
	config
	mutation *InfoTypesMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (itc *InfoTypesCreate) SetCreatedAt(t time.Time) *InfoTypesCreate {
	itc.mutation.SetCreatedAt(t)
	return itc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (itc *InfoTypesCreate) SetNillableCreatedAt(t *time.Time) *InfoTypesCreate {
	if t != nil {
		itc.SetCreatedAt(*t)
	}
	return itc
}

// SetUpdatedAt sets the "updated_at" field.
func (itc *InfoTypesCreate) SetUpdatedAt(t time.Time) *InfoTypesCreate {
	itc.mutation.SetUpdatedAt(t)
	return itc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (itc *InfoTypesCreate) SetNillableUpdatedAt(t *time.Time) *InfoTypesCreate {
	if t != nil {
		itc.SetUpdatedAt(*t)
	}
	return itc
}

// SetDeletedAt sets the "deleted_at" field.
func (itc *InfoTypesCreate) SetDeletedAt(t time.Time) *InfoTypesCreate {
	itc.mutation.SetDeletedAt(t)
	return itc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (itc *InfoTypesCreate) SetNillableDeletedAt(t *time.Time) *InfoTypesCreate {
	if t != nil {
		itc.SetDeletedAt(*t)
	}
	return itc
}

// SetCreatedBy sets the "created_by" field.
func (itc *InfoTypesCreate) SetCreatedBy(i int) *InfoTypesCreate {
	itc.mutation.SetCreatedBy(i)
	return itc
}

// SetUpdatedBy sets the "updated_by" field.
func (itc *InfoTypesCreate) SetUpdatedBy(i int) *InfoTypesCreate {
	itc.mutation.SetUpdatedBy(i)
	return itc
}

// SetDeletedBy sets the "deleted_by" field.
func (itc *InfoTypesCreate) SetDeletedBy(i int) *InfoTypesCreate {
	itc.mutation.SetDeletedBy(i)
	return itc
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (itc *InfoTypesCreate) SetNillableDeletedBy(i *int) *InfoTypesCreate {
	if i != nil {
		itc.SetDeletedBy(*i)
	}
	return itc
}

// SetName sets the "name" field.
func (itc *InfoTypesCreate) SetName(s string) *InfoTypesCreate {
	itc.mutation.SetName(s)
	return itc
}

// AddProductInfoIDs adds the "product_info" edge to the ProductInfo entity by IDs.
func (itc *InfoTypesCreate) AddProductInfoIDs(ids ...int) *InfoTypesCreate {
	itc.mutation.AddProductInfoIDs(ids...)
	return itc
}

// AddProductInfo adds the "product_info" edges to the ProductInfo entity.
func (itc *InfoTypesCreate) AddProductInfo(p ...*ProductInfo) *InfoTypesCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return itc.AddProductInfoIDs(ids...)
}

// Mutation returns the InfoTypesMutation object of the builder.
func (itc *InfoTypesCreate) Mutation() *InfoTypesMutation {
	return itc.mutation
}

// Save creates the InfoTypes in the database.
func (itc *InfoTypesCreate) Save(ctx context.Context) (*InfoTypes, error) {
	itc.defaults()
	return withHooks(ctx, itc.sqlSave, itc.mutation, itc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (itc *InfoTypesCreate) SaveX(ctx context.Context) *InfoTypes {
	v, err := itc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (itc *InfoTypesCreate) Exec(ctx context.Context) error {
	_, err := itc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (itc *InfoTypesCreate) ExecX(ctx context.Context) {
	if err := itc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (itc *InfoTypesCreate) defaults() {
	if _, ok := itc.mutation.CreatedAt(); !ok {
		v := infotypes.DefaultCreatedAt()
		itc.mutation.SetCreatedAt(v)
	}
	if _, ok := itc.mutation.UpdatedAt(); !ok {
		v := infotypes.DefaultUpdatedAt()
		itc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (itc *InfoTypesCreate) check() error {
	if _, ok := itc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "InfoTypes.created_at"`)}
	}
	if _, ok := itc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "InfoTypes.updated_at"`)}
	}
	if _, ok := itc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "InfoTypes.created_by"`)}
	}
	if v, ok := itc.mutation.CreatedBy(); ok {
		if err := infotypes.CreatedByValidator(v); err != nil {
			return &ValidationError{Name: "created_by", err: fmt.Errorf(`ent: validator failed for field "InfoTypes.created_by": %w`, err)}
		}
	}
	if _, ok := itc.mutation.UpdatedBy(); !ok {
		return &ValidationError{Name: "updated_by", err: errors.New(`ent: missing required field "InfoTypes.updated_by"`)}
	}
	if v, ok := itc.mutation.UpdatedBy(); ok {
		if err := infotypes.UpdatedByValidator(v); err != nil {
			return &ValidationError{Name: "updated_by", err: fmt.Errorf(`ent: validator failed for field "InfoTypes.updated_by": %w`, err)}
		}
	}
	if _, ok := itc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "InfoTypes.name"`)}
	}
	return nil
}

func (itc *InfoTypesCreate) sqlSave(ctx context.Context) (*InfoTypes, error) {
	if err := itc.check(); err != nil {
		return nil, err
	}
	_node, _spec := itc.createSpec()
	if err := sqlgraph.CreateNode(ctx, itc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	itc.mutation.id = &_node.ID
	itc.mutation.done = true
	return _node, nil
}

func (itc *InfoTypesCreate) createSpec() (*InfoTypes, *sqlgraph.CreateSpec) {
	var (
		_node = &InfoTypes{config: itc.config}
		_spec = sqlgraph.NewCreateSpec(infotypes.Table, sqlgraph.NewFieldSpec(infotypes.FieldID, field.TypeInt))
	)
	if value, ok := itc.mutation.CreatedAt(); ok {
		_spec.SetField(infotypes.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := itc.mutation.UpdatedAt(); ok {
		_spec.SetField(infotypes.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := itc.mutation.DeletedAt(); ok {
		_spec.SetField(infotypes.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if value, ok := itc.mutation.CreatedBy(); ok {
		_spec.SetField(infotypes.FieldCreatedBy, field.TypeInt, value)
		_node.CreatedBy = value
	}
	if value, ok := itc.mutation.UpdatedBy(); ok {
		_spec.SetField(infotypes.FieldUpdatedBy, field.TypeInt, value)
		_node.UpdatedBy = value
	}
	if value, ok := itc.mutation.DeletedBy(); ok {
		_spec.SetField(infotypes.FieldDeletedBy, field.TypeInt, value)
		_node.DeletedBy = &value
	}
	if value, ok := itc.mutation.Name(); ok {
		_spec.SetField(infotypes.FieldName, field.TypeString, value)
		_node.Name = &value
	}
	if nodes := itc.mutation.ProductInfoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   infotypes.ProductInfoTable,
			Columns: []string{infotypes.ProductInfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(productinfo.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// InfoTypesCreateBulk is the builder for creating many InfoTypes entities in bulk.
type InfoTypesCreateBulk struct {
	config
	err      error
	builders []*InfoTypesCreate
}

// Save creates the InfoTypes entities in the database.
func (itcb *InfoTypesCreateBulk) Save(ctx context.Context) ([]*InfoTypes, error) {
	if itcb.err != nil {
		return nil, itcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(itcb.builders))
	nodes := make([]*InfoTypes, len(itcb.builders))
	mutators := make([]Mutator, len(itcb.builders))
	for i := range itcb.builders {
		func(i int, root context.Context) {
			builder := itcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*InfoTypesMutation)
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
					_, err = mutators[i+1].Mutate(root, itcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, itcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, itcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (itcb *InfoTypesCreateBulk) SaveX(ctx context.Context) []*InfoTypes {
	v, err := itcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (itcb *InfoTypesCreateBulk) Exec(ctx context.Context) error {
	_, err := itcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (itcb *InfoTypesCreateBulk) ExecX(ctx context.Context) {
	if err := itcb.Exec(ctx); err != nil {
		panic(err)
	}
}
