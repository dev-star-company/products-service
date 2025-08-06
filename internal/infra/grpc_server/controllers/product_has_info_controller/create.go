package product_has_info_controller

import (
	"context"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/product_has_info_proto"
)

func (c *controller) Create(ctx context.Context, in *product_has_info_proto.CreateRequest) (*product_has_info_proto.CreateResponse, error) {

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	create, err := c.Db.ProductHasInfo.Create().
		SetProductInfoID(int(in.ProductInfoId)).
		SetProductID(int(in.ProductsId)).
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
