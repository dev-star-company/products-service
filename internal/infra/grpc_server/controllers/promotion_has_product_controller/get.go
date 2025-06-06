package promotion_has_product_controller

import (
	"context"
	"products-service/generated_protos/promotion_has_product_proto"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/promotionhasproduct"
	"products-service/internal/pkg/errs"
)

func (c *controller) Get(ctx context.Context, in *promotion_has_product_proto.GetRequest) (*promotion_has_product_proto.GetResponse, error) {
	promotion_has_product, err := c.Db.PromotionHasProduct.
		Query().
		Where(promotionhasproduct.ID(int(in.Id))).
		WithPromotion().
		WithProduct().
		Only(ctx)

	if ent.IsNotFound(err) {
		return nil, errs.PromotionHasProductNotFound(int(in.Id))
	}

	return &promotion_has_product_proto.GetResponse{
		RequesterId: uint32(promotion_has_product.CreatedBy),
		PromotionId: uint32(*promotion_has_product.PromotionID),
		ProductsId:  uint32(*promotion_has_product.ProductID),
	}, nil
}
