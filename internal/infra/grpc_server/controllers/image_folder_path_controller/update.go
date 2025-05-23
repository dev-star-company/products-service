package image_folder_path_controller

import (
	"context"
	"products-service/generated_protos/image_folder_path_proto"
	"products-service/internal/app/ent"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"
)

func (c *controller) Update(ctx context.Context, in *image_folder_path_proto.UpdateRequest) (*image_folder_path_proto.UpdateResponse, error) {
	if in.RequesterId == 0 {
		return nil, errs.ProductsNotFound(int(in.Id))
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	var image_folder_path *ent.ImageFolderPath

	image_folder_pathQ := tx.ImageFolderPath.UpdateOneID(int(in.Id))

	if in.ImageFolderSourceId != nil && *in.ImageFolderSourceId > 0 {
		image_folder_pathQ.SetImageFolderSourceID(int(*in.ImageFolderSourceId))
	}

	image_folder_pathQ.SetUpdatedBy(int(in.RequesterId))

	image_folder_path, err = image_folder_pathQ.Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, utils.Rollback(tx, errs.ProductsNotFound(int(in.Id)))
		}
		if ent.IsConstraintError(err) {
			return nil, utils.Rollback(tx, errs.InvalidForeignKey(err))
		}
		return nil, errs.SavingError("image_folder_path", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &image_folder_path_proto.UpdateResponse{
		RequesterId:         uint32(image_folder_path.CreatedBy),
		ImageFolderSourceId: uint32(*image_folder_path.ImageFolderSourceID),
	}, nil
}
