package features_controller

import (
	"context"
	"products-service/generated_protos/features_proto"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/features"
	"products-service/internal/pkg/errs"
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

		FeatureValueId: uint32(*features.FeatureValueID),
		Name:           *features.Name,
	}, nil
}
