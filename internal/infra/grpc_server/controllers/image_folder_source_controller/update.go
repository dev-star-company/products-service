package image_folder_source_controller

import (
	"context"
	"products-service/generated_protos/image_folder_source_proto"
	"products-service/internal/app/ent"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"
)

func (c *controller) Update(ctx context.Context, in *image_folder_source_proto.UpdateRequest) (*image_folder_source_proto.UpdateResponse, error) {
	if in.RequesterId == 0 {
		return nil, errs.ProductsNotFound(int(in.Id))
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	var image_folder_source *ent.ImageFolderSource

	image_folder_sourceQ := tx.ImageFolderSource.UpdateOneID(int(in.Id))

	if in.Name != nil && *in.Name != "" {
		image_folder_sourceQ.SetName(string(*in.Name))
	}

	image_folder_sourceQ.SetUpdatedBy(int(in.RequesterId))

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
		RequesterId: uint32(image_folder_source.CreatedBy),
		Name:        *image_folder_source.Name,
		BaseUrl:     *image_folder_source.BaseURL,
		AcessKey:    *image_folder_source.AccessKey,
		SecretKey:   *image_folder_source.SecretKey,
	}, nil
}
