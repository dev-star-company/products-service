package features_values_controller

import (
	"context"
	"products-service/generated_protos/features_values_proto"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/featuresvalues"
	"products-service/internal/pkg/errs"
)

func (c *controller) Get(ctx context.Context, in *features_values_proto.GetRequest) (*features_values_proto.GetResponse, error) {
	features_values, err := c.Db.FeaturesValues.
		Query().
		Where(featuresvalues.ID(int(in.Id))).
		Only(ctx)

	if ent.IsNotFound(err) {
		return nil, errs.FeaturesValuesNotFound(int(in.Id))
	}

	return &features_values_proto.GetResponse{
		RequesterId:         uint32(features_values.CreatedBy),
		FeatureId:           uint32(*features_values.FeatureID),
		FeatureUnitValuesId: uint32(*features_values.FeatureUnitValuesID),
		Value:               *features_values.Value,
	}, nil
}
