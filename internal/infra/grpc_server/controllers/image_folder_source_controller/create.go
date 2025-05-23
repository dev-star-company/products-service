package image_folder_source_controller

import (
	"context"
	"products-service/generated_protos/image_folder_source_proto"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"
)

func (c *controller) Create(ctx context.Context, in *image_folder_source_proto.CreateRequest) (*image_folder_source_proto.CreateResponse, error) {

	if in.RequesterId == 0 {
		return nil, errs.RequesterIdRequired()
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	create, err := c.Db.ImageFolderSource.Create().
		SetName(in.Name).
		SetBaseURL(in.BaseUrl).
		SetAccessKey(*in.AcessKey).
		SetSecretKey(*in.SecretKey).
		SetCreatedBy(int(in.RequesterId)).
		SetUpdatedBy(int(in.RequesterId)).
		Save(ctx)

	if err != nil {
		return nil, errs.CreateError("product type", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &image_folder_source_proto.CreateResponse{
		Name:      *create.Name,
		BaseUrl:   *create.BaseURL,
		AcessKey:  *create.AccessKey,
		SecretKey: *create.SecretKey,
	}, nil
}
