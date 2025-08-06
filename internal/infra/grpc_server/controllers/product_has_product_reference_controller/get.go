package product_has_product_reference_controller

import (
	"context"
	"products-service/internal/adapters/grpc_convertions"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/producthasproductreference"
	"products-service/internal/pkg/errs"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/product_has_product_reference_proto"
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
		Producthasproductreference: grpc_convertions.ProductHasProductReferenceToProto(product_has_product_reference),
	}, nil
}
