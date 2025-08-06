package product_prices_controller

import (
	"context"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"
	"time"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/product_prices_proto"
)

func (c *controller) Delete(ctx context.Context, in *product_prices_proto.DeleteRequest) (*product_prices_proto.DeleteResponse, error) {

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	err = tx.ProductPrices.UpdateOneID(int(in.Id)).
		SetDeletedAt(time.Now()).
		Exec(ctx)

	if err != nil {
		return nil, utils.Rollback(tx, errs.DeleteError("product_prices", err))
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &product_prices_proto.DeleteResponse{}, nil
}
