package images_controller

import (
	"context"
	"products-service/generated_protos/images_proto"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"
	"time"
)

func (c *controller) Delete(ctx context.Context, in *images_proto.DeleteRequest) (*images_proto.DeleteResponse, error) {

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	err = tx.Category.UpdateOneID(int(in.Id)).
		SetDeletedAt(time.Now()).
		Exec(ctx)

	if err != nil {
		return nil, utils.Rollback(tx, errs.DeleteError("images", err))
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &images_proto.DeleteResponse{}, nil
}
