// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"products-service/internal/app/ent/brand"
	"products-service/internal/app/ent/category"
	"products-service/internal/app/ent/images"
	"products-service/internal/app/ent/producthasfeature"
	"products-service/internal/app/ent/producthasimage"
	"products-service/internal/app/ent/producthasinfo"
	"products-service/internal/app/ent/producthasproductreference"
	"products-service/internal/app/ent/productprices"
	"products-service/internal/app/ent/productreferences"
	"products-service/internal/app/ent/products"
	"products-service/internal/app/ent/promotionhasproduct"
	"products-service/internal/app/ent/toolhasproduct"
	"products-service/internal/app/ent/varianttype"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProductsCreate is the builder for creating a Products entity.
type ProductsCreate struct {
	config
	mutation *ProductsMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (pc *ProductsCreate) SetCreatedAt(t time.Time) *ProductsCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pc *ProductsCreate) SetNillableCreatedAt(t *time.Time) *ProductsCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetUpdatedAt sets the "updated_at" field.
func (pc *ProductsCreate) SetUpdatedAt(t time.Time) *ProductsCreate {
	pc.mutation.SetUpdatedAt(t)
	return pc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pc *ProductsCreate) SetNillableUpdatedAt(t *time.Time) *ProductsCreate {
	if t != nil {
		pc.SetUpdatedAt(*t)
	}
	return pc
}

// SetDeletedAt sets the "deleted_at" field.
func (pc *ProductsCreate) SetDeletedAt(t time.Time) *ProductsCreate {
	pc.mutation.SetDeletedAt(t)
	return pc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (pc *ProductsCreate) SetNillableDeletedAt(t *time.Time) *ProductsCreate {
	if t != nil {
		pc.SetDeletedAt(*t)
	}
	return pc
}

// SetCreatedBy sets the "created_by" field.
func (pc *ProductsCreate) SetCreatedBy(i int) *ProductsCreate {
	pc.mutation.SetCreatedBy(i)
	return pc
}

// SetUpdatedBy sets the "updated_by" field.
func (pc *ProductsCreate) SetUpdatedBy(i int) *ProductsCreate {
	pc.mutation.SetUpdatedBy(i)
	return pc
}

// SetDeletedBy sets the "deleted_by" field.
func (pc *ProductsCreate) SetDeletedBy(i int) *ProductsCreate {
	pc.mutation.SetDeletedBy(i)
	return pc
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (pc *ProductsCreate) SetNillableDeletedBy(i *int) *ProductsCreate {
	if i != nil {
		pc.SetDeletedBy(*i)
	}
	return pc
}

// SetCategoryID sets the "category_id" field.
func (pc *ProductsCreate) SetCategoryID(i int) *ProductsCreate {
	pc.mutation.SetCategoryID(i)
	return pc
}

// SetNillableCategoryID sets the "category_id" field if the given value is not nil.
func (pc *ProductsCreate) SetNillableCategoryID(i *int) *ProductsCreate {
	if i != nil {
		pc.SetCategoryID(*i)
	}
	return pc
}

// SetBrandID sets the "brand_id" field.
func (pc *ProductsCreate) SetBrandID(i int) *ProductsCreate {
	pc.mutation.SetBrandID(i)
	return pc
}

// SetNillableBrandID sets the "brand_id" field if the given value is not nil.
func (pc *ProductsCreate) SetNillableBrandID(i *int) *ProductsCreate {
	if i != nil {
		pc.SetBrandID(*i)
	}
	return pc
}

// SetVariantTypeID sets the "variant_type_id" field.
func (pc *ProductsCreate) SetVariantTypeID(i int) *ProductsCreate {
	pc.mutation.SetVariantTypeID(i)
	return pc
}

// SetNillableVariantTypeID sets the "variant_type_id" field if the given value is not nil.
func (pc *ProductsCreate) SetNillableVariantTypeID(i *int) *ProductsCreate {
	if i != nil {
		pc.SetVariantTypeID(*i)
	}
	return pc
}

// SetProductReferencesID sets the "product_references_id" field.
func (pc *ProductsCreate) SetProductReferencesID(i int) *ProductsCreate {
	pc.mutation.SetProductReferencesID(i)
	return pc
}

// SetNillableProductReferencesID sets the "product_references_id" field if the given value is not nil.
func (pc *ProductsCreate) SetNillableProductReferencesID(i *int) *ProductsCreate {
	if i != nil {
		pc.SetProductReferencesID(*i)
	}
	return pc
}

// SetImageID sets the "image_id" field.
func (pc *ProductsCreate) SetImageID(i int) *ProductsCreate {
	pc.mutation.SetImageID(i)
	return pc
}

// SetName sets the "name" field.
func (pc *ProductsCreate) SetName(s string) *ProductsCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetStock sets the "stock" field.
func (pc *ProductsCreate) SetStock(i int) *ProductsCreate {
	pc.mutation.SetStock(i)
	return pc
}

// SetCategory sets the "category" edge to the Category entity.
func (pc *ProductsCreate) SetCategory(c *Category) *ProductsCreate {
	return pc.SetCategoryID(c.ID)
}

// SetBrand sets the "brand" edge to the Brand entity.
func (pc *ProductsCreate) SetBrand(b *Brand) *ProductsCreate {
	return pc.SetBrandID(b.ID)
}

// SetVariantType sets the "variant_type" edge to the VariantType entity.
func (pc *ProductsCreate) SetVariantType(v *VariantType) *ProductsCreate {
	return pc.SetVariantTypeID(v.ID)
}

// AddProductReferenceIDs adds the "product_references" edge to the ProductReferences entity by IDs.
func (pc *ProductsCreate) AddProductReferenceIDs(ids ...int) *ProductsCreate {
	pc.mutation.AddProductReferenceIDs(ids...)
	return pc
}

// AddProductReferences adds the "product_references" edges to the ProductReferences entity.
func (pc *ProductsCreate) AddProductReferences(p ...*ProductReferences) *ProductsCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pc.AddProductReferenceIDs(ids...)
}

// AddImageIDs adds the "images" edge to the Images entity by IDs.
func (pc *ProductsCreate) AddImageIDs(ids ...int) *ProductsCreate {
	pc.mutation.AddImageIDs(ids...)
	return pc
}

// AddImages adds the "images" edges to the Images entity.
func (pc *ProductsCreate) AddImages(i ...*Images) *ProductsCreate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return pc.AddImageIDs(ids...)
}

// AddProductHasImageIDs adds the "product_has_image" edge to the ProductHasImage entity by IDs.
func (pc *ProductsCreate) AddProductHasImageIDs(ids ...int) *ProductsCreate {
	pc.mutation.AddProductHasImageIDs(ids...)
	return pc
}

// AddProductHasImage adds the "product_has_image" edges to the ProductHasImage entity.
func (pc *ProductsCreate) AddProductHasImage(p ...*ProductHasImage) *ProductsCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pc.AddProductHasImageIDs(ids...)
}

// AddPromotionHasProductIDs adds the "promotion_has_product" edge to the PromotionHasProduct entity by IDs.
func (pc *ProductsCreate) AddPromotionHasProductIDs(ids ...int) *ProductsCreate {
	pc.mutation.AddPromotionHasProductIDs(ids...)
	return pc
}

// AddPromotionHasProduct adds the "promotion_has_product" edges to the PromotionHasProduct entity.
func (pc *ProductsCreate) AddPromotionHasProduct(p ...*PromotionHasProduct) *ProductsCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pc.AddPromotionHasProductIDs(ids...)
}

// AddToolHasProductIDs adds the "tool_has_product" edge to the ToolHasProduct entity by IDs.
func (pc *ProductsCreate) AddToolHasProductIDs(ids ...int) *ProductsCreate {
	pc.mutation.AddToolHasProductIDs(ids...)
	return pc
}

// AddToolHasProduct adds the "tool_has_product" edges to the ToolHasProduct entity.
func (pc *ProductsCreate) AddToolHasProduct(t ...*ToolHasProduct) *ProductsCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return pc.AddToolHasProductIDs(ids...)
}

// AddProductHasFeatureIDs adds the "product_has_feature" edge to the ProductHasFeature entity by IDs.
func (pc *ProductsCreate) AddProductHasFeatureIDs(ids ...int) *ProductsCreate {
	pc.mutation.AddProductHasFeatureIDs(ids...)
	return pc
}

// AddProductHasFeature adds the "product_has_feature" edges to the ProductHasFeature entity.
func (pc *ProductsCreate) AddProductHasFeature(p ...*ProductHasFeature) *ProductsCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pc.AddProductHasFeatureIDs(ids...)
}

// AddProductHasInfoIDs adds the "product_has_info" edge to the ProductHasInfo entity by IDs.
func (pc *ProductsCreate) AddProductHasInfoIDs(ids ...int) *ProductsCreate {
	pc.mutation.AddProductHasInfoIDs(ids...)
	return pc
}

// AddProductHasInfo adds the "product_has_info" edges to the ProductHasInfo entity.
func (pc *ProductsCreate) AddProductHasInfo(p ...*ProductHasInfo) *ProductsCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pc.AddProductHasInfoIDs(ids...)
}

// AddProductHasProductReferenceIDs adds the "product_has_product_reference" edge to the ProductHasProductReference entity by IDs.
func (pc *ProductsCreate) AddProductHasProductReferenceIDs(ids ...int) *ProductsCreate {
	pc.mutation.AddProductHasProductReferenceIDs(ids...)
	return pc
}

// AddProductHasProductReference adds the "product_has_product_reference" edges to the ProductHasProductReference entity.
func (pc *ProductsCreate) AddProductHasProductReference(p ...*ProductHasProductReference) *ProductsCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pc.AddProductHasProductReferenceIDs(ids...)
}

// AddProductPriceIDs adds the "product_prices" edge to the ProductPrices entity by IDs.
func (pc *ProductsCreate) AddProductPriceIDs(ids ...int) *ProductsCreate {
	pc.mutation.AddProductPriceIDs(ids...)
	return pc
}

// AddProductPrices adds the "product_prices" edges to the ProductPrices entity.
func (pc *ProductsCreate) AddProductPrices(p ...*ProductPrices) *ProductsCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pc.AddProductPriceIDs(ids...)
}

// Mutation returns the ProductsMutation object of the builder.
func (pc *ProductsCreate) Mutation() *ProductsMutation {
	return pc.mutation
}

// Save creates the Products in the database.
func (pc *ProductsCreate) Save(ctx context.Context) (*Products, error) {
	pc.defaults()
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *ProductsCreate) SaveX(ctx context.Context) *Products {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *ProductsCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *ProductsCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *ProductsCreate) defaults() {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		v := products.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		v := products.DefaultUpdatedAt()
		pc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *ProductsCreate) check() error {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Products.created_at"`)}
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Products.updated_at"`)}
	}
	if _, ok := pc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "Products.created_by"`)}
	}
	if v, ok := pc.mutation.CreatedBy(); ok {
		if err := products.CreatedByValidator(v); err != nil {
			return &ValidationError{Name: "created_by", err: fmt.Errorf(`ent: validator failed for field "Products.created_by": %w`, err)}
		}
	}
	if _, ok := pc.mutation.UpdatedBy(); !ok {
		return &ValidationError{Name: "updated_by", err: errors.New(`ent: missing required field "Products.updated_by"`)}
	}
	if v, ok := pc.mutation.UpdatedBy(); ok {
		if err := products.UpdatedByValidator(v); err != nil {
			return &ValidationError{Name: "updated_by", err: fmt.Errorf(`ent: validator failed for field "Products.updated_by": %w`, err)}
		}
	}
	if _, ok := pc.mutation.ImageID(); !ok {
		return &ValidationError{Name: "image_id", err: errors.New(`ent: missing required field "Products.image_id"`)}
	}
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Products.name"`)}
	}
	if _, ok := pc.mutation.Stock(); !ok {
		return &ValidationError{Name: "stock", err: errors.New(`ent: missing required field "Products.stock"`)}
	}
	return nil
}

func (pc *ProductsCreate) sqlSave(ctx context.Context) (*Products, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *ProductsCreate) createSpec() (*Products, *sqlgraph.CreateSpec) {
	var (
		_node = &Products{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(products.Table, sqlgraph.NewFieldSpec(products.FieldID, field.TypeInt))
	)
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.SetField(products.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.SetField(products.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := pc.mutation.DeletedAt(); ok {
		_spec.SetField(products.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if value, ok := pc.mutation.CreatedBy(); ok {
		_spec.SetField(products.FieldCreatedBy, field.TypeInt, value)
		_node.CreatedBy = value
	}
	if value, ok := pc.mutation.UpdatedBy(); ok {
		_spec.SetField(products.FieldUpdatedBy, field.TypeInt, value)
		_node.UpdatedBy = value
	}
	if value, ok := pc.mutation.DeletedBy(); ok {
		_spec.SetField(products.FieldDeletedBy, field.TypeInt, value)
		_node.DeletedBy = &value
	}
	if value, ok := pc.mutation.CategoryID(); ok {
		_spec.SetField(products.FieldCategoryID, field.TypeInt, value)
		_node.CategoryID = &value
	}
	if value, ok := pc.mutation.BrandID(); ok {
		_spec.SetField(products.FieldBrandID, field.TypeInt, value)
		_node.BrandID = &value
	}
	if value, ok := pc.mutation.VariantTypeID(); ok {
		_spec.SetField(products.FieldVariantTypeID, field.TypeInt, value)
		_node.VariantTypeID = &value
	}
	if value, ok := pc.mutation.ProductReferencesID(); ok {
		_spec.SetField(products.FieldProductReferencesID, field.TypeInt, value)
		_node.ProductReferencesID = &value
	}
	if value, ok := pc.mutation.ImageID(); ok {
		_spec.SetField(products.FieldImageID, field.TypeInt, value)
		_node.ImageID = &value
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.SetField(products.FieldName, field.TypeString, value)
		_node.Name = &value
	}
	if value, ok := pc.mutation.Stock(); ok {
		_spec.SetField(products.FieldStock, field.TypeInt, value)
		_node.Stock = value
	}
	if nodes := pc.mutation.CategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   products.CategoryTable,
			Columns: []string{products.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.category_products = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.BrandIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   products.BrandTable,
			Columns: []string{products.BrandColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(brand.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.brand_products = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.VariantTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   products.VariantTypeTable,
			Columns: []string{products.VariantTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(varianttype.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.variant_type_products = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.ProductReferencesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   products.ProductReferencesTable,
			Columns: []string{products.ProductReferencesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(productreferences.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.ImagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   products.ImagesTable,
			Columns: []string{products.ImagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(images.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.ProductHasImageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   products.ProductHasImageTable,
			Columns: []string{products.ProductHasImageColumn},
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
	if nodes := pc.mutation.PromotionHasProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   products.PromotionHasProductTable,
			Columns: []string{products.PromotionHasProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(promotionhasproduct.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.ToolHasProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   products.ToolHasProductTable,
			Columns: []string{products.ToolHasProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(toolhasproduct.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.ProductHasFeatureIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   products.ProductHasFeatureTable,
			Columns: []string{products.ProductHasFeatureColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(producthasfeature.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.ProductHasInfoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   products.ProductHasInfoTable,
			Columns: []string{products.ProductHasInfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(producthasinfo.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.ProductHasProductReferenceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   products.ProductHasProductReferenceTable,
			Columns: []string{products.ProductHasProductReferenceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(producthasproductreference.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.ProductPricesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   products.ProductPricesTable,
			Columns: []string{products.ProductPricesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(productprices.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ProductsCreateBulk is the builder for creating many Products entities in bulk.
type ProductsCreateBulk struct {
	config
	err      error
	builders []*ProductsCreate
}

// Save creates the Products entities in the database.
func (pcb *ProductsCreateBulk) Save(ctx context.Context) ([]*Products, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Products, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProductsMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *ProductsCreateBulk) SaveX(ctx context.Context) []*Products {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *ProductsCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *ProductsCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
