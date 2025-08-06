package grpc_convertions

import (
	"products-service/internal/app/ent"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/product_prices_proto"
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
		CreatedAt:    product_prices.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	if product_prices.DeletedAt != nil {
		x := product_prices.DeletedAt.Format("2006-01-02 15:04:05")
		cur.DeletedAt = &x
	}

	return cur
}
