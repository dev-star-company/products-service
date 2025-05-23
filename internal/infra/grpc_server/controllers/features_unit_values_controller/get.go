package features_unit_values_controller

import (
	"context"
	"products-service/generated_protos/features_unit_values_proto"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/featuresunitvalues"
	"products-service/internal/pkg/errs"
)

func (c *controller) Get(ctx context.Context, in *features_unit_values_proto.GetRequest) (*features_unit_values_proto.GetResponse, error) {
	features_unit_values, err := c.Db.FeaturesUnitValues.
		Query().
		Where(featuresunitvalues.ID(int(in.Id))).
		Only(ctx)

	if ent.IsNotFound(err) {
		return nil, errs.FeaturesUnitValuesNotFound(int(in.Id))
	}

	return &features_unit_values_proto.GetResponse{
		RequesterId: uint32(features_unit_values.CreatedBy),
		Decimals:    float32(*features_unit_values.Decimals),
		Name:        *features_unit_values.Name,
	}, nil
}
