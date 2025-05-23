package images_controller

import (
	"context"
	"products-service/generated_protos/images_proto"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"
)

func (c *controller) Create(ctx context.Context, in *images_proto.CreateRequest) (*images_proto.CreateResponse, error) {

	if in.RequesterId == 0 {
		return nil, errs.RequesterIdRequired()
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	create, err := c.Db.Images.Create().
		SetImageFolderPathID(int(in.ImageFolderPathId)).
		SetContent(in.Content).
		SetPath(in.Path).
		SetCreatedBy(int(in.RequesterId)).
		SetUpdatedBy(int(in.RequesterId)).
		Save(ctx)

	if err != nil {
		return nil, errs.CreateError("images", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &images_proto.CreateResponse{
		ImageFolderPathId: uint32(*create.ImageFolderPathID),
		Content:           string(*create.Content),
		Path:              string(*create.Path),
	}, nil
}
