package products_controller

import (
	"context"
	"products-service/generated_protos/products_proto"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"
)

func (c *controller) Create(ctx context.Context, in *products_proto.CreateRequest) (*products_proto.CreateResponse, error) {

	if in.RequesterId == 0 {
		return nil, errs.RequesterIdRequired()
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	create, err := c.Db.Products.Create().
		SetCategoryID(int(*in.CategoryId)).
		SetBrandID(int(*in.BrandId)).
		SetVariantTypeID(int(*in.VariantTypeId)).
		SetProductReferencesID(int(*in.ProductReferencesId)).
		SetImageID(int(*in.ImageId)).
		SetName(in.Name).
		SetStock(int(in.Stock)).
		SetCreatedBy(int(in.RequesterId)).
		SetUpdatedBy(int(in.RequesterId)).
		Save(ctx)

	if err != nil {
		return nil, errs.CreateError("product type", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &products_proto.CreateResponse{
		CategoryId:          uint32(*create.CategoryID),
		BrandId:             uint32(*create.BrandID),
		VariantTypeId:       uint32(*create.VariantTypeID),
		ProductReferencesId: uint32(*create.ProductReferencesID),
		ImageId:             uint32(*create.ImageID),
		Name:                string(*create.Name),
		Stock:               uint32(create.Stock),
	}, nil
}
