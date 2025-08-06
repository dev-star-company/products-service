package features_values_types_controller

import (
	"context"
	"products-service/internal/adapters/grpc_convertions"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/featuresvaluestypes"
	"products-service/internal/pkg/errs"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/features_values_types_proto"
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
		Featuresvaluestypes: grpc_convertions.FeaturesValuesTypesToProto(features_values_types),
	}, nil
}
