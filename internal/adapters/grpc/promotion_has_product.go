package grpc_convertions

import (
	"products-service/internal/app/ent"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/promotion_has_product_proto"
)

func PromotionHasProductToProto(promotion_has_product *ent.PromotionHasProduct) *promotion_has_product_proto.PromotionHasProduct {
	if promotion_has_product == nil {
		return nil
	}

	cur := &promotion_has_product_proto.PromotionHasProduct{
		Id:               uint32(promotion_has_product.ID),
		ProductsId:       uint32(*promotion_has_product.ProductsID),
		PromotionId:      uint32(*promotion_has_product.PromotionsID),
		PromocionalPrice: float64(*promotion_has_product.PromocionalPrice),
		CreatedAt:        promotion_has_product.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	if promotion_has_product.DeletedAt != nil {
		x := promotion_has_product.DeletedAt.Format("2006-01-02 15:04:05")
		cur.DeletedAt = &x
	}

	return cur
}
