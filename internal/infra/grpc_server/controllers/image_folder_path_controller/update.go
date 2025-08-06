package image_folder_path_controller

import (
	"context"
	"products-service/internal/adapters/grpc_convertions"
	"products-service/internal/app/ent"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/image_folder_path_proto"
)

func (c *controller) Update(ctx context.Context, in *image_folder_path_proto.UpdateRequest) (*image_folder_path_proto.UpdateResponse, error) {

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	var image_folder_path *ent.ImageFolderPath

	image_folder_pathQ := tx.ImageFolderPath.UpdateOneID(int(in.Id))

	if in.ImageFolderSourceId != nil && *in.ImageFolderSourceId > 0 {
		image_folder_pathQ.SetImageFolderSourceID(int(*in.ImageFolderSourceId))
	}

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
		Imagefolderpath: grpc_convertions.ImageFolderPathToProto(image_folder_path),
	}, nil
}
