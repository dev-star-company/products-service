package image_folder_source_controller

import (
	"context"
	"products-service/internal/adapters/grpc_convertions"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/imagefoldersource"
	"products-service/internal/pkg/errs"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/image_folder_source_proto"
)

func (c *controller) Get(ctx context.Context, in *image_folder_source_proto.GetRequest) (*image_folder_source_proto.GetResponse, error) {
	image_folder_source, err := c.Db.ImageFolderSource.
		Query().
		Where(imagefoldersource.ID(int(in.Id))).
		Only(ctx)

	if ent.IsNotFound(err) {
		return nil, errs.ImageFolderSourceNotFound(int(in.Id))
	}

	return &image_folder_source_proto.GetResponse{
		Imagefoldersource: grpc_convertions.ImageFolderSourceToProto(image_folder_source),
	}, nil
}
