package products_controller

import (
	"context"
	"products-service/internal/adapters/grpc_convertions"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/products"
	"products-service/internal/pkg/errs"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/products_proto"
)

func (c *controller) Get(ctx context.Context, in *products_proto.GetRequest) (*products_proto.GetResponse, error) {
	products, err := c.Db.Products.
		Query().
		Where(products.ID(int(in.Id))).
		WithCategory().
		WithBrand().
		WithVariantType().
		WithProductReferences().
		// WithImages().
		Only(ctx)

	if ent.IsNotFound(err) {
		return nil, errs.ProductsNotFound(int(in.Id))
	}

	return &products_proto.GetResponse{
		Products: grpc_convertions.ProductsToProto(products),
	}, nil
}
