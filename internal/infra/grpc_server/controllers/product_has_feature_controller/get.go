package product_has_feature_controller

import (
	"context"
	"products-service/internal/adapters/grpc_convertions"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/producthasfeature"
	"products-service/internal/pkg/errs"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/product_has_feature_proto"
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
		Producthasfeature: grpc_convertions.ProductHasFeatureToProto(product_has_feature),
	}, nil
}
