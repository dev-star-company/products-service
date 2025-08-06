package features_controller

import (
	"context"
	"products-service/internal/adapters/grpc_convertions"
	"products-service/internal/app/ent"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/features_proto"
)

func (c *controller) Update(ctx context.Context, in *features_proto.UpdateRequest) (*features_proto.UpdateResponse, error) {

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	var features *ent.Features

	featuresQ := tx.Features.UpdateOneID(int(in.Id))

	if in.FeatureValueId != nil && *in.FeatureValueId > 0 {
		featuresQ.SetFeatureValueID(int(*in.FeatureValueId))
	}

	if in.Name != nil && *in.Name != "" {
		featuresQ.SetName(string(*in.Name))
	}

	features, err = featuresQ.Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, utils.Rollback(tx, errs.ProductsNotFound(int(in.Id)))
		}
		if ent.IsConstraintError(err) {
			return nil, utils.Rollback(tx, errs.InvalidForeignKey(err))
		}
		return nil, errs.SavingError("features", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &features_proto.UpdateResponse{
		Features: grpc_convertions.FeaturesToProto(features),
	}, nil
}
