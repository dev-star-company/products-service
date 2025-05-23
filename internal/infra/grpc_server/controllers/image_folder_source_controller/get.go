package image_folder_source_controller

import (
	"context"
	"products-service/generated_protos/image_folder_source_proto"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/imagefoldersource"
	"products-service/internal/pkg/errs"
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
		RequesterId: uint32(image_folder_source.CreatedBy),
		Name:        *image_folder_source.Name,
		BaseUrl:     *image_folder_source.BaseURL,
		AcessKey:    *image_folder_source.AccessKey,
		SecretKey:   *image_folder_source.SecretKey,
	}, nil
}
