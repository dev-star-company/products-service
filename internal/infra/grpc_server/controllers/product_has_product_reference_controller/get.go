package product_has_product_reference_controller

import (
	"context"
	"products-service/generated_protos/product_has_product_reference_proto"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/producthasproductreference"
	"products-service/internal/pkg/errs"
)

func (c *controller) Get(ctx context.Context, in *product_has_product_reference_proto.GetRequest) (*product_has_product_reference_proto.GetResponse, error) {
	product_has_product_reference, err := c.Db.ProductHasProductReference.
		Query().
		Where(producthasproductreference.ID(int(in.Id))).
		WithProductReference().
		WithProduct().
		Only(ctx)

	if ent.IsNotFound(err) {
		return nil, errs.ProductHasProductReferenceNotFound(int(in.Id))
	}

	return &product_has_product_reference_proto.GetResponse{

		ProductReferenceId: uint32(*product_has_product_reference.ProductReferenceID),
		ProductsId:         uint32(*product_has_product_reference.ProductID),
	}, nil
}
