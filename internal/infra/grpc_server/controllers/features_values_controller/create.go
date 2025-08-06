package features_values_controller

import (
	"context"
	"products-service/internal/adapters/grpc_convertions"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/features_values_proto"
)

func (c *controller) Create(ctx context.Context, in *features_values_proto.CreateRequest) (*features_values_proto.CreateResponse, error) {

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	create, err := c.Db.FeaturesValues.Create().
		SetFeatureID(int(in.FeatureId)).
		SetFeatureUnitValuesID(int(in.FeatureUnitValuesId)).
		SetValue(in.Value).
		Save(ctx)

	if err != nil {
		return nil, errs.CreateError("product type", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &features_values_proto.CreateResponse{
		Featuresvalues: grpc_convertions.FeaturesValuesToProto(create),
	}, nil
}
