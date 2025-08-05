package images_controller

import (
	"context"
	"products-service/generated_protos/images_proto"
	"products-service/internal/app/ent"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"
)

func (c *controller) Update(ctx context.Context, in *images_proto.UpdateRequest) (*images_proto.UpdateResponse, error) {

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	var images *ent.Images

	imagesQ := tx.Images.UpdateOneID(int(in.Id))

	if in.ImageFolderPathId != nil && *in.ImageFolderPathId > 0 {
		imagesQ.SetImageFolderPathID(int(*in.ImageFolderPathId))
	}

	if in.Content != nil && *in.Content != "" {
		imagesQ.SetContent(string(*in.Content))
	}

	if in.Path != nil && *in.Path != "" {
		imagesQ.SetPath(string(*in.Path))
	}

	images, err = imagesQ.Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, utils.Rollback(tx, errs.ProductsNotFound(int(in.Id)))
		}
		if ent.IsConstraintError(err) {
			return nil, utils.Rollback(tx, errs.InvalidForeignKey(err))
		}
		return nil, errs.SavingError("images", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &images_proto.UpdateResponse{

		Content:           string(*images.Content),
		Path:              string(*images.Path),
	}, nil
}
