package product_prices_controller

import (
	"context"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/productprices"
	"products-service/internal/pkg/errs"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/product_prices_proto"
)

func (c *controller) Get(ctx context.Context, in *product_prices_proto.GetRequest) (*product_prices_proto.GetResponse, error) {
	product_prices, err := c.Db.ProductPrices.
		Query().
		Where(productprices.ID(int(in.Id))).
		Only(ctx)

	if ent.IsNotFound(err) {
		return nil, errs.ProductPricesNotFound(int(in.Id))
	}

	return &product_prices_proto.GetResponse{

		PriceTypeId:  uint32(*product_prices.PriceTypeID),
		ProductsId:   uint32(*product_prices.ProductID),
		DefaultValue: float64(*product_prices.DefaultValue),
		MinValue:     float64(*product_prices.MinValue),
	}, nil
}
