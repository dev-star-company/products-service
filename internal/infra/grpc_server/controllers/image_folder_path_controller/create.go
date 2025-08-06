package image_folder_path_controller

import (
	"context"
	"products-service/internal/adapters/grpc_convertions"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/image_folder_path_proto"
)

func (c *controller) Create(ctx context.Context, in *image_folder_path_proto.CreateRequest) (*image_folder_path_proto.CreateResponse, error) {

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	create, err := c.Db.ImageFolderPath.Create().
		SetImageFolderSourceID(int(in.ImageFolderSourceId)).
		Save(ctx)

	if err != nil {
		return nil, errs.CreateError("product type", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &image_folder_path_proto.CreateResponse{
		Imagefolderpath: grpc_convertions.ImageFolderPathToProto(create),
	}, nil
}
