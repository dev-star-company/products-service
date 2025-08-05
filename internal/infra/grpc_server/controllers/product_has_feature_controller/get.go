package product_has_feature_controller

import (
	"context"
	"products-service/generated_protos/product_has_feature_proto"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/producthasfeature"
	"products-service/internal/pkg/errs"
)

func (c *controller) Get(ctx context.Context, in *product_has_feature_proto.GetRequest) (*product_has_feature_proto.GetResponse, error) {
	product_has_feature, err := c.Db.ProductHasFeature.
		Query().
		Where(producthasfeature.ID(int(in.Id))).
		WithFeatures().
		WithProducts().
		Only(ctx)

	if ent.IsNotFound(err) {
		return nil, errs.ProductHasFeatureNotFound(int(in.Id))
	}

	return &product_has_feature_proto.GetResponse{

		FeaturesId: uint32(*product_has_feature.FeatureID),
		ProductsId: uint32(*product_has_feature.ProductID),
	}, nil
}
