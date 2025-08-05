package price_type_controller

import (
	"context"
	"products-service/generated_protos/price_type_proto"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"
)

func (c *controller) Create(ctx context.Context, in *price_type_proto.CreateRequest) (*price_type_proto.CreateResponse, error) {

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	create, err := c.Db.PriceType.Create().
		SetName(in.Name).
		Save(ctx)

	if err != nil {
		return nil, errs.CreateError("product type", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &price_type_proto.CreateResponse{
		Name: string(*create.Name),
	}, nil
}
