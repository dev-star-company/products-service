package features_unit_values_controller

import (
	"context"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/featuresunitvalues"
	"products-service/internal/pkg/errs"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/features_unit_values_proto"
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

		Decimals: float32(*features_unit_values.Decimals),
		Name:     *features_unit_values.Name,
	}, nil
}
