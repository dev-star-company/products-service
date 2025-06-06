package products_controller

import (
	"context"
	"products-service/generated_protos/products_proto"
	"products-service/internal/app/ent"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"
)

func (c *controller) Update(ctx context.Context, in *products_proto.UpdateRequest) (*products_proto.UpdateResponse, error) {
	if in.RequesterId == 0 {
		return nil, errs.ProductsNotFound(int(in.Id))
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	var products *ent.Products

	productsQ := tx.Products.UpdateOneID(int(in.Id))

	if in.CategoryId != nil && *in.CategoryId > 0 {
		productsQ.SetCategoryID(int(*in.CategoryId))
	}

	if in.BrandId != nil && *in.BrandId > 0 {
		productsQ.SetBrandID(int(*in.BrandId))
	}

	if in.VariantTypeId != nil && *in.VariantTypeId > 0 {
		productsQ.SetVariantTypeID(int(*in.VariantTypeId))
	}

	if in.ProductReferencesId != nil && *in.ProductReferencesId > 0 {
		productsQ.SetProductReferencesID(int(*in.ProductReferencesId))
	}

	if in.ImageId != nil && *in.ImageId > 0 {
		productsQ.SetImageID(int(*in.ImageId))
	}

	if in.Name != nil && *in.Name != "" {
		productsQ.SetName(string(*in.Name))
	}

	if in.Stock != nil && *in.Stock > 0 {
		productsQ.SetStock(int(*in.Stock))
	}

	productsQ.SetUpdatedBy(int(in.RequesterId))

	products, err = productsQ.Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, utils.Rollback(tx, errs.ProductsNotFound(int(in.Id)))
		}
		if ent.IsConstraintError(err) {
			return nil, utils.Rollback(tx, errs.InvalidForeignKey(err))
		}
		return nil, errs.SavingError("products", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &products_proto.UpdateResponse{
		RequesterId:         uint32(products.CreatedBy),
		CategoryId:          uint32(*products.CategoryID),
		BrandId:             uint32(*products.BrandID),
		VariantTypeId:       uint32(*products.VariantTypeID),
		ProductReferencesId: uint32(*products.ProductReferencesID),
		ImageId:             uint32(*products.ImageID),
		Name:                string(*products.Name),
		Stock:               uint32(products.Stock),
	}, nil
}
