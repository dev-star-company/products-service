package products_controller

import (
	"context"
	"products-service/generated_protos/products_proto"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/products"
	"products-service/internal/pkg/errs"
)

func (c *controller) Get(ctx context.Context, in *products_proto.GetRequest) (*products_proto.GetResponse, error) {
	products, err := c.Db.Products.
		Query().
		Where(products.ID(int(in.Id))).
		WithCategory().
		WithBrand().
		WithVariantType().
		WithProductReferences().
		WithImages().
		Only(ctx)

	if ent.IsNotFound(err) {
		return nil, errs.ProductsNotFound(int(in.Id))
	}

	return &products_proto.GetResponse{
		RequesterId:         uint32(products.CreatedBy),
		CategoryId:          uint32(*products.CategoryID),
		BrandId:             uint32(*products.BrandID),
		VariantTypeId:       uint32(*products.VariantTypeID),
		ProductReferencesId: uint32(*products.ProductReferencesID),
		ImageId:             uint32(*products.ImageID),
		Name:                *products.Name,
		Stock:               uint32(products.Stock),
	}, nil
}
