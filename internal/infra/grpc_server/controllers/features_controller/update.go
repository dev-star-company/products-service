package features_controller

import (
	"context"
	"products-service/generated_protos/features_proto"
	"products-service/internal/app/ent"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"
)

func (c *controller) Update(ctx context.Context, in *features_proto.UpdateRequest) (*features_proto.UpdateResponse, error) {
	if in.RequesterId == 0 {
		return nil, errs.ProductsNotFound(int(in.Id))
	}

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

	featuresQ.SetUpdatedBy(int(in.RequesterId))

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
		RequesterId:    uint32(features.CreatedBy),
		FeatureValueId: uint32(*features.FeatureValueID),
		Name:           string(*features.Name),
	}, nil
}
