// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"products-service/internal/app/ent/imagefolderpath"
	"products-service/internal/app/ent/images"
	"products-service/internal/app/ent/producthasimage"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ImagesCreate is the builder for creating a Images entity.
type ImagesCreate struct {
	config
	mutation *ImagesMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (ic *ImagesCreate) SetCreatedAt(t time.Time) *ImagesCreate {
	ic.mutation.SetCreatedAt(t)
	return ic
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ic *ImagesCreate) SetNillableCreatedAt(t *time.Time) *ImagesCreate {
	if t != nil {
		ic.SetCreatedAt(*t)
	}
	return ic
}

// SetUpdatedAt sets the "updated_at" field.
func (ic *ImagesCreate) SetUpdatedAt(t time.Time) *ImagesCreate {
	ic.mutation.SetUpdatedAt(t)
	return ic
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ic *ImagesCreate) SetNillableUpdatedAt(t *time.Time) *ImagesCreate {
	if t != nil {
		ic.SetUpdatedAt(*t)
	}
	return ic
}

// SetDeletedAt sets the "deleted_at" field.
func (ic *ImagesCreate) SetDeletedAt(t time.Time) *ImagesCreate {
	ic.mutation.SetDeletedAt(t)
	return ic
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ic *ImagesCreate) SetNillableDeletedAt(t *time.Time) *ImagesCreate {
	if t != nil {
		ic.SetDeletedAt(*t)
	}
	return ic
}

// SetCreatedBy sets the "created_by" field.
func (ic *ImagesCreate) SetCreatedBy(i int) *ImagesCreate {
	ic.mutation.SetCreatedBy(i)
	return ic
}

// SetUpdatedBy sets the "updated_by" field.
func (ic *ImagesCreate) SetUpdatedBy(i int) *ImagesCreate {
	ic.mutation.SetUpdatedBy(i)
	return ic
}

// SetDeletedBy sets the "deleted_by" field.
func (ic *ImagesCreate) SetDeletedBy(i int) *ImagesCreate {
	ic.mutation.SetDeletedBy(i)
	return ic
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (ic *ImagesCreate) SetNillableDeletedBy(i *int) *ImagesCreate {
	if i != nil {
		ic.SetDeletedBy(*i)
	}
	return ic
}

// SetImageFolderPathID sets the "image_folder_path_id" field.
func (ic *ImagesCreate) SetImageFolderPathID(i int) *ImagesCreate {
	ic.mutation.SetImageFolderPathID(i)
	return ic
}

// SetContent sets the "content" field.
func (ic *ImagesCreate) SetContent(s string) *ImagesCreate {
	ic.mutation.SetContent(s)
	return ic
}

// SetPath sets the "path" field.
func (ic *ImagesCreate) SetPath(s string) *ImagesCreate {
	ic.mutation.SetPath(s)
	return ic
}

// SetImageFolderPath sets the "image_folder_path" edge to the ImageFolderPath entity.
func (ic *ImagesCreate) SetImageFolderPath(i *ImageFolderPath) *ImagesCreate {
	return ic.SetImageFolderPathID(i.ID)
}

// AddProductHasImageIDs adds the "product_has_image" edge to the ProductHasImage entity by IDs.
func (ic *ImagesCreate) AddProductHasImageIDs(ids ...int) *ImagesCreate {
	ic.mutation.AddProductHasImageIDs(ids...)
	return ic
}

// AddProductHasImage adds the "product_has_image" edges to the ProductHasImage entity.
func (ic *ImagesCreate) AddProductHasImage(p ...*ProductHasImage) *ImagesCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ic.AddProductHasImageIDs(ids...)
}

// Mutation returns the ImagesMutation object of the builder.
func (ic *ImagesCreate) Mutation() *ImagesMutation {
	return ic.mutation
}

// Save creates the Images in the database.
func (ic *ImagesCreate) Save(ctx context.Context) (*Images, error) {
	ic.defaults()
	return withHooks(ctx, ic.sqlSave, ic.mutation, ic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ic *ImagesCreate) SaveX(ctx context.Context) *Images {
	v, err := ic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ic *ImagesCreate) Exec(ctx context.Context) error {
	_, err := ic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ic *ImagesCreate) ExecX(ctx context.Context) {
	if err := ic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ic *ImagesCreate) defaults() {
	if _, ok := ic.mutation.CreatedAt(); !ok {
		v := images.DefaultCreatedAt()
		ic.mutation.SetCreatedAt(v)
	}
	if _, ok := ic.mutation.UpdatedAt(); !ok {
		v := images.DefaultUpdatedAt()
		ic.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ic *ImagesCreate) check() error {
	if _, ok := ic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Images.created_at"`)}
	}
	if _, ok := ic.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Images.updated_at"`)}
	}
	if _, ok := ic.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "Images.created_by"`)}
	}
	if v, ok := ic.mutation.CreatedBy(); ok {
		if err := images.CreatedByValidator(v); err != nil {
			return &ValidationError{Name: "created_by", err: fmt.Errorf(`ent: validator failed for field "Images.created_by": %w`, err)}
		}
	}
	if _, ok := ic.mutation.UpdatedBy(); !ok {
		return &ValidationError{Name: "updated_by", err: errors.New(`ent: missing required field "Images.updated_by"`)}
	}
	if v, ok := ic.mutation.UpdatedBy(); ok {
		if err := images.UpdatedByValidator(v); err != nil {
			return &ValidationError{Name: "updated_by", err: fmt.Errorf(`ent: validator failed for field "Images.updated_by": %w`, err)}
		}
	}
	if _, ok := ic.mutation.ImageFolderPathID(); !ok {
		return &ValidationError{Name: "image_folder_path_id", err: errors.New(`ent: missing required field "Images.image_folder_path_id"`)}
	}
	if _, ok := ic.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "Images.content"`)}
	}
	if _, ok := ic.mutation.Path(); !ok {
		return &ValidationError{Name: "path", err: errors.New(`ent: missing required field "Images.path"`)}
	}
	if len(ic.mutation.ImageFolderPathIDs()) == 0 {
		return &ValidationError{Name: "image_folder_path", err: errors.New(`ent: missing required edge "Images.image_folder_path"`)}
	}
	return nil
}

func (ic *ImagesCreate) sqlSave(ctx context.Context) (*Images, error) {
	if err := ic.check(); err != nil {
		return nil, err
	}
	_node, _spec := ic.createSpec()
	if err := sqlgraph.CreateNode(ctx, ic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ic.mutation.id = &_node.ID
	ic.mutation.done = true
	return _node, nil
}

func (ic *ImagesCreate) createSpec() (*Images, *sqlgraph.CreateSpec) {
	var (
		_node = &Images{config: ic.config}
		_spec = sqlgraph.NewCreateSpec(images.Table, sqlgraph.NewFieldSpec(images.FieldID, field.TypeInt))
	)
	if value, ok := ic.mutation.CreatedAt(); ok {
		_spec.SetField(images.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ic.mutation.UpdatedAt(); ok {
		_spec.SetField(images.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ic.mutation.DeletedAt(); ok {
		_spec.SetField(images.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if value, ok := ic.mutation.CreatedBy(); ok {
		_spec.SetField(images.FieldCreatedBy, field.TypeInt, value)
		_node.CreatedBy = value
	}
	if value, ok := ic.mutation.UpdatedBy(); ok {
		_spec.SetField(images.FieldUpdatedBy, field.TypeInt, value)
		_node.UpdatedBy = value
	}
	if value, ok := ic.mutation.DeletedBy(); ok {
		_spec.SetField(images.FieldDeletedBy, field.TypeInt, value)
		_node.DeletedBy = &value
	}
	if value, ok := ic.mutation.Content(); ok {
		_spec.SetField(images.FieldContent, field.TypeString, value)
		_node.Content = &value
	}
	if value, ok := ic.mutation.Path(); ok {
		_spec.SetField(images.FieldPath, field.TypeString, value)
		_node.Path = &value
	}
	if nodes := ic.mutation.ImageFolderPathIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   images.ImageFolderPathTable,
			Columns: []string{images.ImageFolderPathColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(imagefolderpath.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ImageFolderPathID = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ic.mutation.ProductHasImageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   images.ProductHasImageTable,
			Columns: []string{images.ProductHasImageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(producthasimage.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ImagesCreateBulk is the builder for creating many Images entities in bulk.
type ImagesCreateBulk struct {
	config
	err      error
	builders []*ImagesCreate
}

// Save creates the Images entities in the database.
func (icb *ImagesCreateBulk) Save(ctx context.Context) ([]*Images, error) {
	if icb.err != nil {
		return nil, icb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(icb.builders))
	nodes := make([]*Images, len(icb.builders))
	mutators := make([]Mutator, len(icb.builders))
	for i := range icb.builders {
		func(i int, root context.Context) {
			builder := icb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ImagesMutation)
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
					_, err = mutators[i+1].Mutate(root, icb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, icb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, icb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (icb *ImagesCreateBulk) SaveX(ctx context.Context) []*Images {
	v, err := icb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (icb *ImagesCreateBulk) Exec(ctx context.Context) error {
	_, err := icb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (icb *ImagesCreateBulk) ExecX(ctx context.Context) {
	if err := icb.Exec(ctx); err != nil {
		panic(err)
	}
}
