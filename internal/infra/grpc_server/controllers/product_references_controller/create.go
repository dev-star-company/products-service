package product_references_controller

import (
	"context"
	"products-service/generated_protos/product_references_proto"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"
)

func (c *controller) Create(ctx context.Context, in *product_references_proto.CreateRequest) (*product_references_proto.CreateResponse, error) {

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	create, err := c.Db.ProductReferences.Create().
		SetReferenceSourceID(int(in.ReferenceSourceId)).
		SetValue(in.Value).
		Save(ctx)

	if err != nil {
		return nil, errs.CreateError("product type", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &product_references_proto.CreateResponse{
		ReferenceSourceId: uint32(*create.ReferenceSourceID),
		Value:             string(*create.Value),
	}, nil
}
