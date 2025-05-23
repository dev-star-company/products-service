package promotion_has_product_controller

import (
	"context"
	"products-service/generated_protos/promotion_has_product_proto"
	"products-service/internal/app/ent"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"
)

func (c *controller) Update(ctx context.Context, in *promotion_has_product_proto.UpdateRequest) (*promotion_has_product_proto.UpdateResponse, error) {
	if in.RequesterId == 0 {
		return nil, errs.ProductsNotFound(int(in.Id))
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	var promotion_has_product *ent.PromotionHasProduct

	promotion_has_productQ := tx.PromotionHasProduct.UpdateOneID(int(in.Id))

	if in.PromotionId != nil && *in.PromotionId > 0 {
		promotion_has_productQ.SetPromotionID(int(*in.PromotionId))
	}

	if in.ProductsId != nil && *in.ProductsId > 0 {
		promotion_has_productQ.SetProductID(int(*in.ProductsId))
	}

	promotion_has_productQ.SetUpdatedBy(int(in.RequesterId))

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
		RequesterId: uint32(promotion_has_product.CreatedBy),
		PromotionId: uint32(*promotion_has_product.PromotionID),
		ProductsId:  uint32(*promotion_has_product.ProductID),
	}, nil
}
