package image_folder_path_controller

import (
	"context"
	"products-service/generated_protos/image_folder_path_proto"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/imagefolderpath"
	"products-service/internal/pkg/errs"
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
		RequesterId:         uint32(image_folder_path.CreatedBy),
		ImageFolderSourceId: uint32(*image_folder_path.ImageFolderSourceID),
	}, nil
}
