package features_values_types_controller

import (
	"context"
	"products-service/generated_protos/features_values_types_proto"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/featuresvaluestypes"
	"products-service/internal/pkg/errs"
)

func (c *controller) Get(ctx context.Context, in *features_values_types_proto.GetRequest) (*features_values_types_proto.GetResponse, error) {
	features_values_types, err := c.Db.FeaturesValuesTypes.
		Query().
		Where(featuresvaluestypes.ID(int(in.Id))).
		Only(ctx)

	if ent.IsNotFound(err) {
		return nil, errs.FeaturesValuesTypesNotFound(int(in.Id))
	}

	return &features_values_types_proto.GetResponse{
		RequesterId: uint32(features_values_types.CreatedBy),
		Name:        *features_values_types.Name,
	}, nil
}
