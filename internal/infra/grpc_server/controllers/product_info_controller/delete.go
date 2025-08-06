package product_info_controller

import (
	"context"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"
	"time"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/product_info_proto"
)

func (c *controller) Delete(ctx context.Context, in *product_info_proto.DeleteRequest) (*product_info_proto.DeleteResponse, error) {

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	err = tx.ProductInfo.UpdateOneID(int(in.Id)).
		SetDeletedAt(time.Now()).
		Exec(ctx)

	if err != nil {
		return nil, utils.Rollback(tx, errs.DeleteError("product_info", err))
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &product_info_proto.DeleteResponse{}, nil
}
