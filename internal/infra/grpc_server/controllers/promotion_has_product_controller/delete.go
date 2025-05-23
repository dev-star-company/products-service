package promotion_has_product_controller

import (
	"context"
	"products-service/generated_protos/promotion_has_product_proto"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"
	"time"
)

func (c *controller) Delete(ctx context.Context, in *promotion_has_product_proto.DeleteRequest) (*promotion_has_product_proto.DeleteResponse, error) {
	if in.RequesterId == 0 {
		return nil, errs.PromotionHasProductNotFound(int(in.Id))
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	err = tx.PromotionHasProduct.UpdateOneID(int(in.Id)).
		SetDeletedAt(time.Now()).
		SetDeletedBy(int(in.RequesterId)).
		Exec(ctx)

	if err != nil {
		return nil, utils.Rollback(tx, errs.DeleteError("promotion_has_product", err))
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &promotion_has_product_proto.DeleteResponse{}, nil
}
