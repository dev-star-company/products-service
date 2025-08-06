package image_folder_path_controller

import (
	"context"
	"products-service/internal/adapters/grpc_convertions"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/imagefolderpath"
	"products-service/internal/pkg/errs"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/image_folder_path_proto"
)

func (c *controller) Get(ctx context.Context, in *image_folder_path_proto.GetRequest) (*image_folder_path_proto.GetResponse, error) {
	image_folder_path, err := c.Db.ImageFolderPath.
		Query().
		Where(imagefolderpath.ID(int(in.Id))).
		Only(ctx)

	if ent.IsNotFound(err) {
		return nil, errs.ImageFolderPathNotFound(int(in.Id))
	}

	return &image_folder_path_proto.GetResponse{
		Imagefolderpath: grpc_convertions.ImageFolderPathToProto(image_folder_path),
	}, nil
}
