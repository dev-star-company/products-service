package product_has_image_controller

import (
	"context"
	"products-service/internal/app/ent"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/product_has_image_proto"
)

func (c *controller) Update(ctx context.Context, in *product_has_image_proto.UpdateRequest) (*product_has_image_proto.UpdateResponse, error) {

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	var product_has_image *ent.ProductHasImage

	product_has_imageQ := tx.ProductHasImage.UpdateOneID(int(in.Id))

	if in.ProductsId != nil && *in.ProductsId > 0 {
		product_has_imageQ.SetProductID(int(*in.ProductsId))
	}

	if in.ImagesId != nil && *in.ImagesId > 0 {
		product_has_imageQ.SetImageID(int(*in.ImagesId))
	}

	if in.Priority != nil && *in.Priority > 0 {
		product_has_imageQ.SetPriority(int(*in.Priority))
	}

	product_has_image, err = product_has_imageQ.Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, utils.Rollback(tx, errs.ProductsNotFound(int(in.Id)))
		}
		if ent.IsConstraintError(err) {
			return nil, utils.Rollback(tx, errs.InvalidForeignKey(err))
		}
		return nil, errs.SavingError("product_has_image", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &product_has_image_proto.UpdateResponse{

		ProductsId: uint32(*product_has_image.ProductID),
		ImagesId:   uint32(*product_has_image.ImageID),
		Priority:   uint32(product_has_image.Priority),
	}, nil
}
