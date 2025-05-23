package features_values_controller

import (
	"context"
	"products-service/generated_protos/features_values_proto"
	"products-service/internal/app/ent"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"
)

func (c *controller) Update(ctx context.Context, in *features_values_proto.UpdateRequest) (*features_values_proto.UpdateResponse, error) {
	if in.RequesterId == 0 {
		return nil, errs.ProductsNotFound(int(in.Id))
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	var features_values *ent.FeaturesValues

	features_valuesQ := tx.FeaturesValues.UpdateOneID(int(in.Id))

	if in.FeatureId != nil && *in.FeatureId > 0 {
		features_valuesQ.SetFeatureID(int(*in.FeatureId))
	}

	if in.FeatureUnitValuesId != nil && *in.FeatureUnitValuesId > 0 {
		features_valuesQ.SetFeatureUnitValuesID(int(*in.FeatureUnitValuesId))
	}

	if in.Value != nil && *in.Value != "" {
		features_valuesQ.SetValue(string(*in.Value))
	}

	features_valuesQ.SetUpdatedBy(int(in.RequesterId))

	features_values, err = features_valuesQ.Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, utils.Rollback(tx, errs.ProductsNotFound(int(in.Id)))
		}
		if ent.IsConstraintError(err) {
			return nil, utils.Rollback(tx, errs.InvalidForeignKey(err))
		}
		return nil, errs.SavingError("features_values", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &features_values_proto.UpdateResponse{
		RequesterId:         uint32(features_values.CreatedBy),
		FeatureId:           uint32(*features_values.FeatureID),
		FeatureUnitValuesId: uint32(*features_values.FeatureUnitValuesID),
		Value:               string(*features_values.Value),
	}, nil
}
