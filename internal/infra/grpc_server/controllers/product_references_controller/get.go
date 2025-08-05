package product_references_controller

import (
	"context"
	"products-service/generated_protos/product_references_proto"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/productreferences"
	"products-service/internal/pkg/errs"
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

		ReferenceSourceId: uint32(*product_references.ReferenceSourceID),
		Value:             *product_references.Value,
	}, nil
}
