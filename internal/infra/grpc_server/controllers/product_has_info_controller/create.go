package product_has_info_controller

import (
	"context"
	"products-service/generated_protos/product_has_info_proto"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"
)

func (c *controller) Create(ctx context.Context, in *product_has_info_proto.CreateRequest) (*product_has_info_proto.CreateResponse, error) {

	if in.RequesterId == 0 {
		return nil, errs.RequesterIdRequired()
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	create, err := c.Db.ProductHasInfo.Create().
		SetProductInfoID(int(in.ProductInfoId)).
		SetProductID(int(in.ProductsId)).
		SetCreatedBy(int(in.RequesterId)).
		SetUpdatedBy(int(in.RequesterId)).
		Save(ctx)

	if err != nil {
		return nil, errs.CreateError("product type", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &product_has_info_proto.CreateResponse{
		ProductInfoId: uint32(*create.ProductInfoID),
		ProductsId:    uint32(*create.ProductID),
	}, nil
}
