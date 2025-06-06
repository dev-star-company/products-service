package product_prices_controller

import (
	"context"
	"products-service/generated_protos/product_prices_proto"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"
)

func (c *controller) Create(ctx context.Context, in *product_prices_proto.CreateRequest) (*product_prices_proto.CreateResponse, error) {

	if in.RequesterId == 0 {
		return nil, errs.RequesterIdRequired()
	}

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
		SetCreatedBy(int(in.RequesterId)).
		SetUpdatedBy(int(in.RequesterId)).
		Save(ctx)

	if err != nil {
		return nil, errs.CreateError("product type", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &product_prices_proto.CreateResponse{
		PriceTypeId:  uint32(*create.PriceTypeID),
		ProductsId:   uint32(*create.ProductID),
		DefaultValue: float64(*create.DefaultValue),
		MinValue:     float64(*create.MinValue),
	}, nil
}
