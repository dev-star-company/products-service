package product_has_info_controller

import (
	"context"
	"products-service/generated_protos/product_has_info_proto"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"
	"time"
)

func (c *controller) Delete(ctx context.Context, in *product_has_info_proto.DeleteRequest) (*product_has_info_proto.DeleteResponse, error) {
	if in.RequesterId == 0 {
		return nil, errs.ProductHasInfoNotFound(int(in.Id))
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	err = tx.ProductHasInfo.UpdateOneID(int(in.Id)).
		SetDeletedAt(time.Now()).
		SetDeletedBy(int(in.RequesterId)).
		Exec(ctx)

	if err != nil {
		return nil, utils.Rollback(tx, errs.DeleteError("product_has_info", err))
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &product_has_info_proto.DeleteResponse{}, nil
}
