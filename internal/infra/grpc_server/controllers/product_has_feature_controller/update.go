package product_has_feature_controller

import (
	"context"
	"products-service/internal/app/ent"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/product_has_feature_proto"
)

func (c *controller) Update(ctx context.Context, in *product_has_feature_proto.UpdateRequest) (*product_has_feature_proto.UpdateResponse, error) {

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	var product_has_feature *ent.ProductHasFeature

	product_has_featureQ := tx.ProductHasFeature.UpdateOneID(int(in.Id))

	if in.FeaturesId != nil && *in.FeaturesId > 0 {
		product_has_featureQ.SetFeaturesID(int(*in.FeaturesId))
	}

	if in.ProductsId != nil && *in.ProductsId > 0 {
		product_has_featureQ.SetProductID(int(*in.ProductsId))
	}

	product_has_feature, err = product_has_featureQ.Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, utils.Rollback(tx, errs.ProductsNotFound(int(in.Id)))
		}
		if ent.IsConstraintError(err) {
			return nil, utils.Rollback(tx, errs.InvalidForeignKey(err))
		}
		return nil, errs.SavingError("product_has_feature", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &product_has_feature_proto.UpdateResponse{

		FeaturesId: uint32(*product_has_feature.FeatureID),
		ProductsId: uint32(*product_has_feature.ProductID),
	}, nil
}
