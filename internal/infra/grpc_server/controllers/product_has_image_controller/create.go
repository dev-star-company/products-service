package product_has_image_controller

import (
	"context"
	"products-service/generated_protos/product_has_image_proto"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"
)

func (c *controller) Create(ctx context.Context, in *product_has_image_proto.CreateRequest) (*product_has_image_proto.CreateResponse, error) {

	if in.RequesterId == 0 {
		return nil, errs.RequesterIdRequired()
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	create, err := c.Db.ProductHasImage.Create().
		SetProductID(int(in.ProductsId)).
		SetImageID(int(in.ImagesId)).
		SetPriority(int(in.Priority)).
		SetCreatedBy(int(in.RequesterId)).
		SetUpdatedBy(int(in.RequesterId)).
		Save(ctx)

	if err != nil {
		return nil, errs.CreateError("product type", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &product_has_image_proto.CreateResponse{
		ProductsId: uint32(*create.ProductID),
		ImagesId:   uint32(*create.ImageID),
		Priority:   uint32(create.Priority),
	}, nil
}
