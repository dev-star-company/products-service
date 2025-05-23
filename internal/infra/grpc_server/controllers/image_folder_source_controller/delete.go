package image_folder_source_controller

import (
	"context"
	"products-service/generated_protos/image_folder_source_proto"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"
	"time"
)

func (c *controller) Delete(ctx context.Context, in *image_folder_source_proto.DeleteRequest) (*image_folder_source_proto.DeleteResponse, error) {
	if in.RequesterId == 0 {
		return nil, errs.ImageFolderSourceNotFound(int(in.Id))
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	err = tx.ImageFolderSource.UpdateOneID(int(in.Id)).
		SetDeletedAt(time.Now()).
		SetDeletedBy(int(in.RequesterId)).
		Exec(ctx)

	if err != nil {
		return nil, utils.Rollback(tx, errs.DeleteError("image_folder_source", err))
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &image_folder_source_proto.DeleteResponse{}, nil
}
