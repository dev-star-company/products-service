package product_has_image_controller

import (
	"context"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/producthasimage"
	"products-service/internal/pkg/errs"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/product_has_image_proto"
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

		ProductsId: uint32(*product_has_image.ProductID),
		ImagesId:   uint32(*product_has_image.ImageID),
		Priority:   uint32(product_has_image.Priority),
	}, nil
}
