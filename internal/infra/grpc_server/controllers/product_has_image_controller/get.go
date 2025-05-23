package product_has_image_controller

import (
	"context"
	"products-service/generated_protos/product_has_image_proto"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/producthasimage"
	"products-service/internal/pkg/errs"
)

func (c *controller) Get(ctx context.Context, in *product_has_image_proto.GetRequest) (*product_has_image_proto.GetResponse, error) {
	product_has_image, err := c.Db.ProductHasImage.
		Query().
		Where(producthasimage.ID(int(in.Id))).
		Only(ctx)

	if ent.IsNotFound(err) {
		return nil, errs.ProductsNotFound(int(in.Id))
	}

	return &product_has_image_proto.GetResponse{
		RequesterId: uint32(product_has_image.CreatedBy),
		ProductsId:  uint32(*product_has_image.ProductID),
		ImagesId:    uint32(*product_has_image.ImageID),
		Priority:    uint32(product_has_image.Priority),
	}, nil
}
