package image_folder_source_controller

import (
	"context"
	"products-service/internal/adapters/grpc_convertions"
	"products-service/internal/app/ent"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/image_folder_source_proto"
)

func (c *controller) Update(ctx context.Context, in *image_folder_source_proto.UpdateRequest) (*image_folder_source_proto.UpdateResponse, error) {

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	var image_folder_source *ent.ImageFolderSource

	image_folder_sourceQ := tx.ImageFolderSource.UpdateOneID(int(in.Id))

	if in.Name != nil && *in.Name != "" {
		image_folder_sourceQ.SetName(string(*in.Name))
	}

	image_folder_source, err = image_folder_sourceQ.Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, utils.Rollback(tx, errs.ProductsNotFound(int(in.Id)))
		}
		if ent.IsConstraintError(err) {
			return nil, utils.Rollback(tx, errs.InvalidForeignKey(err))
		}
		return nil, errs.SavingError("image_folder_source", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &image_folder_source_proto.UpdateResponse{
		Imagefoldersource: grpc_convertions.ImageFolderSourceToProto(image_folder_source),
	}, nil
}
