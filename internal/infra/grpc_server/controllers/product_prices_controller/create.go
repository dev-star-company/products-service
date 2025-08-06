package product_prices_controller

import (
	"context"
	"products-service/internal/adapters/grpc_convertions"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/product_prices_proto"
)

func (c *controller) Create(ctx context.Context, in *product_prices_proto.CreateRequest) (*product_prices_proto.CreateResponse, error) {

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	create, err := c.Db.ProductPrices.Create().
		SetPriceTypeID(int(in.PriceTypeId)).
		SetProductID(int(in.ProductsId)).
		SetDefaultValue(float64(in.DefaultValue)).
		SetMinValue(float64(*in.MinValue)).
		SetProductID(int(in.ProductsId)).
		Save(ctx)

	if err != nil {
		return nil, errs.CreateError("product type", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &product_prices_proto.CreateResponse{
		Productprices: grpc_convertions.ProductPricesToProto(create),
	}, nil
}
