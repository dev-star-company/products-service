package product_references_controller

import (
	"context"
	"products-service/internal/adapters/grpc_convertions"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/productreferences"
	"products-service/internal/pkg/errs"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/product_references_proto"
)

func (c *controller) Get(ctx context.Context, in *product_references_proto.GetRequest) (*product_references_proto.GetResponse, error) {
	product_references, err := c.Db.ProductReferences.
		Query().
		Where(productreferences.ID(int(in.Id))).
		Only(ctx)

	if ent.IsNotFound(err) {
		return nil, errs.ProductReferencesNotFound(int(in.Id))
	}

	return &product_references_proto.GetResponse{
		Productreferences: grpc_convertions.ProductReferencesToProto(product_references),
	}, nil
}
