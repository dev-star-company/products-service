package product_references_controller

import (
	"context"
	"products-service/generated_protos/product_references_proto"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"
	"time"
)

func (c *controller) Delete(ctx context.Context, in *product_references_proto.DeleteRequest) (*product_references_proto.DeleteResponse, error) {

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	err = tx.ProductReferences.UpdateOneID(int(in.Id)).
		SetDeletedAt(time.Now()).
		Exec(ctx)

	if err != nil {
		return nil, utils.Rollback(tx, errs.DeleteError("product_references", err))
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &product_references_proto.DeleteResponse{}, nil
}
