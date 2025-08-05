package tools_controller

import (
	"context"
	"products-service/generated_protos/tools_proto"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"
	"time"
)

func (c *controller) Delete(ctx context.Context, in *tools_proto.DeleteRequest) (*tools_proto.DeleteResponse, error) {

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	err = tx.Brand.UpdateOneID(int(in.Id)).
		SetDeletedAt(time.Now()).
		Exec(ctx)

	if err != nil {
		return nil, utils.Rollback(tx, errs.DeleteError("tools", err))
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &tools_proto.DeleteResponse{}, nil
}
