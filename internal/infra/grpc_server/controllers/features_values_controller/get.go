package features_values_controller

import (
	"context"
	"products-service/internal/adapters/grpc_convertions"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/featuresvalues"
	"products-service/internal/pkg/errs"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/features_values_proto"
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
		Featuresvalues: grpc_convertions.FeaturesValuesToProto(features_values),
	}, nil
}
