package promotion_has_product_controller

import (
	"context"
	"products-service/generated_protos/promotion_has_product_proto"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"
)

func (c *controller) Create(ctx context.Context, in *promotion_has_product_proto.CreateRequest) (*promotion_has_product_proto.CreateResponse, error) {

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	create, err := c.Db.PromotionHasProduct.Create().
		SetPromotionsID(int(in.PromotionId)).
		SetProductsID(int(in.ProductsId)).
		Save(ctx)

	if err != nil {
		return nil, errs.CreateError("product type", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &promotion_has_product_proto.CreateResponse{
		PromotionId: uint32(*create.PromotionsID),
		ProductsId:  uint32(*create.ProductsID),
	}, nil
}
