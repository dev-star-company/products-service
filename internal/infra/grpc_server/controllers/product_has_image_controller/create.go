package product_has_image_controller

import (
	"context"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/product_has_image_proto"
)

func (c *controller) Create(ctx context.Context, in *product_has_image_proto.CreateRequest) (*product_has_image_proto.CreateResponse, error) {

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	create, err := c.Db.ProductHasImage.Create().
		SetProductID(int(in.ProductsId)).
		SetImageID(int(in.ImagesId)).
		SetPriority(int(in.Priority)).
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
