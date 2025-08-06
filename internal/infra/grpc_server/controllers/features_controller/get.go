package features_controller

import (
	"context"
	"products-service/internal/adapters/grpc_convertions"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/features"
	"products-service/internal/pkg/errs"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/features_proto"
)

func (c *controller) Get(ctx context.Context, in *features_proto.GetRequest) (*features_proto.GetResponse, error) {
	features, err := c.Db.Features.
		Query().
		Where(features.ID(int(in.Id))).
		Only(ctx)

	if ent.IsNotFound(err) {
		return nil, errs.FeaturesNotFound(int(in.Id))
	}

	return &features_proto.GetResponse{
		Features: grpc_convertions.FeaturesToProto(features),
	}, nil
}
