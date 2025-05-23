package grpc_convertions

import (
	"products-service/generated_protos/product_prices_proto"
	"products-service/internal/app/ent"
)

func ProductPricesToProto(product_prices *ent.ProductPrices) *product_prices_proto.ProductPrices {
	if product_prices == nil {
		return nil
	}

	cur := &product_prices_proto.ProductPrices{
		Id:           uint32(product_prices.ID),
		ProductsId:   uint32(*product_prices.ProductID),
		PriceTypeId:  uint32(*product_prices.PriceTypeID),
		DefaultValue: *product_prices.DefaultValue,
		MinValue:     product_prices.MinValue,
		CreatedBy:    uint32(product_prices.CreatedBy),
		UpdatedBy:    uint32(product_prices.UpdatedBy),
		CreatedAt:    product_prices.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:    product_prices.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	if product_prices.DeletedAt != nil {
		x := product_prices.DeletedAt.Format("2006-01-02 15:04:05")
		cur.DeletedAt = &x
	}

	if product_prices.DeletedBy != nil {
		x := uint32(*product_prices.DeletedBy)
		cur.DeletedBy = &x
	}

	return cur
}
