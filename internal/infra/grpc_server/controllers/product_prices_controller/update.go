package product_prices_controller

import (
	"context"
	"products-service/internal/adapters/grpc_convertions"
	"products-service/internal/app/ent"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/product_prices_proto"
)

func (c *controller) Update(ctx context.Context, in *product_prices_proto.UpdateRequest) (*product_prices_proto.UpdateResponse, error) {

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	var product_prices *ent.ProductPrices

	product_pricesQ := tx.ProductPrices.UpdateOneID(int(in.Id))

	if in.ProductsId != nil && *in.ProductsId > 0 {
		product_pricesQ.SetProductID(int(*in.ProductsId))
	}

	if in.PriceTypeId != nil && *in.PriceTypeId > 0 {
		product_pricesQ.SetPriceTypeID(int(*in.PriceTypeId))
	}

	if in.DefaultValue != nil && *in.DefaultValue != 0 {
		product_pricesQ.SetDefaultValue(float64(*in.DefaultValue))
	}

	if in.MinValue != nil && *in.MinValue != 0 {
		product_pricesQ.SetMinValue(float64(*in.MinValue))
	}

	product_prices, err = product_pricesQ.Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, utils.Rollback(tx, errs.ProductsNotFound(int(in.Id)))
		}
		if ent.IsConstraintError(err) {
			return nil, utils.Rollback(tx, errs.InvalidForeignKey(err))
		}
		return nil, errs.SavingError("product_prices", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &product_prices_proto.UpdateResponse{
		Productprices: grpc_convertions.ProductPricesToProto(product_prices),
	}, nil
}
