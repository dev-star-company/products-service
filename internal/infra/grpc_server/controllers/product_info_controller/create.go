package product_info_controller

import (
	"context"
	"products-service/generated_protos/product_info_proto"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"
)

func (c *controller) Create(ctx context.Context, in *product_info_proto.CreateRequest) (*product_info_proto.CreateResponse, error) {

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	create, err := c.Db.ProductInfo.Create().
		SetInfoTypesID(int(in.InfoTypeId)).
		SetValue(in.Value).
		Save(ctx)

	if err != nil {
		return nil, errs.CreateError("product type", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &product_info_proto.CreateResponse{
		InfoTypeId: uint32(*create.InfoTypesID),
		Value:      string(*create.Value),
	}, nil
}
