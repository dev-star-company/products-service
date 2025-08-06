package promotion_has_product_controller

import (
	"context"
	"products-service/internal/adapters/grpc_convertions"
	"products-service/internal/app/ent"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/promotion_has_product_proto"
)

func (c *controller) Update(ctx context.Context, in *promotion_has_product_proto.UpdateRequest) (*promotion_has_product_proto.UpdateResponse, error) {

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	var promotion_has_product *ent.PromotionHasProduct

	promotion_has_productQ := tx.PromotionHasProduct.UpdateOneID(int(in.Id))

	if in.PromotionId != nil && *in.PromotionId > 0 {
		promotion_has_productQ.SetPromotionsID(int(*in.PromotionId))
	}

	if in.ProductsId != nil && *in.ProductsId > 0 {
		promotion_has_productQ.SetProductsID(int(*in.ProductsId))
	}

	promotion_has_product, err = promotion_has_productQ.Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, utils.Rollback(tx, errs.ProductsNotFound(int(in.Id)))
		}
		if ent.IsConstraintError(err) {
			return nil, utils.Rollback(tx, errs.InvalidForeignKey(err))
		}
		return nil, errs.SavingError("promotion_has_product", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &promotion_has_product_proto.UpdateResponse{
		Promotionhasproduct: grpc_convertions.PromotionHasProductToProto(promotion_has_product),
	}, nil
}
