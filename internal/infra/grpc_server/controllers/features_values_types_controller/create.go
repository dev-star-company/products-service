package features_values_types_controller

import (
	"context"
	"products-service/generated_protos/features_values_types_proto"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"
)

func (c *controller) Create(ctx context.Context, in *features_values_types_proto.CreateRequest) (*features_values_types_proto.CreateResponse, error) {

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	create, err := c.Db.FeaturesValuesTypes.Create().
		SetName(in.Name).
		Save(ctx)

	if err != nil {
		return nil, errs.CreateError("product type", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &features_values_types_proto.CreateResponse{
		Name: string(*create.Name),
	}, nil
}
