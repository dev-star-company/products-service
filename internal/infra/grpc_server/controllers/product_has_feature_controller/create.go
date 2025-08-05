package product_has_feature_controller

import (
	"context"
	"products-service/generated_protos/product_has_feature_proto"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"
)

func (c *controller) Create(ctx context.Context, in *product_has_feature_proto.CreateRequest) (*product_has_feature_proto.CreateResponse, error) {

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	create, err := c.Db.ProductHasFeature.Create().
		SetFeatureID(int(in.FeaturesId)).
		SetProductID(int(in.ProductsId)).
		Save(ctx)

	if err != nil {
		return nil, errs.CreateError("product type", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &product_has_feature_proto.CreateResponse{
		FeaturesId: uint32(*create.FeatureID),
		ProductsId: uint32(*create.ProductID),
	}, nil
}
