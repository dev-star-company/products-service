package grpc_convertions

import (
	"products-service/generated_protos/promotion_has_product_proto"
	"products-service/internal/app/ent"
)

func PromotionHasProductToProto(promotion_has_product *ent.PromotionHasProduct) *promotion_has_product_proto.PromotionHasProduct {
	if promotion_has_product == nil {
		return nil
	}

	cur := &promotion_has_product_proto.PromotionHasProduct{
		Id:               uint32(promotion_has_product.ID),
		ProductsId:       uint32(*promotion_has_product.ProductID),
		PromotionId:      uint32(*promotion_has_product.PromotionID),
		PromocionalPrice: float64(*promotion_has_product.PromocionalPrice),
		CreatedBy:        uint32(promotion_has_product.CreatedBy),
		UpdatedBy:        uint32(promotion_has_product.UpdatedBy),
		CreatedAt:        promotion_has_product.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:        promotion_has_product.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	if promotion_has_product.DeletedAt != nil {
		x := promotion_has_product.DeletedAt.Format("2006-01-02 15:04:05")
		cur.DeletedAt = &x
	}

	if promotion_has_product.DeletedBy != nil {
		x := uint32(*promotion_has_product.DeletedBy)
		cur.DeletedBy = &x
	}

	return cur
}
