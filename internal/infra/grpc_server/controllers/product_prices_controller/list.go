package product_prices_controller

import (
	"context"
	"errors"
	"products-service/internal/adapters/grpc_convertions"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/productprices"
	"products-service/internal/app/ent/schema"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/product_prices_proto"
)

func (c *controller) List(ctx context.Context, in *product_prices_proto.ListRequest) (*product_prices_proto.ListResponse, error) {
	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	if in.IncludeDeleted != nil && *in.IncludeDeleted {
		ctx = schema.SkipSoftDelete(ctx)
	}

	query := tx.ProductPrices.Query()

	if in.PriceTypeId != nil {
		query = query.Where(productprices.PriceTypeID(int(*in.PriceTypeId)))
	}

	if in.ProductsId != nil {
		query = query.Where(productprices.ProductID(int(*in.ProductsId)))
	}

	if in.DefaultValue != nil {
		query = query.Where(productprices.DefaultValue(float64(*in.DefaultValue)))
	}

	if in.MinValue != nil {
		query = query.Where(productprices.MinValue(float64(*in.MinValue)))
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errs.ListingError("querying product_prices", err)
	}

	if in.Limit != nil && *in.Limit > 0 {
		query = query.Limit(int(*in.Limit))
	}

	if in.Offset != nil && *in.Offset > 0 {
		query = query.Offset(int(*in.Offset))
	}

	if in.Orderby != nil {
		if in.Orderby.Id != nil {
			switch *in.Orderby.Id {
			case "ASC":
				query = query.Order(ent.Asc(productprices.FieldID))
			case "DESC":
				query = query.Order(ent.Desc(productprices.FieldID))
			default:
				return nil, errs.InvalidOrderByValue(errors.New(*in.Orderby.Id))
			}
		}
	}

	product_prices, err := query.All(ctx)
	if err != nil {
		return nil, errs.ListingError("querying product_prices", err)
	}

	responseProductPrices := make([]*product_prices_proto.ProductPrices, len(product_prices))
	for i, acc := range product_prices {
		responseProductPrices[i] = grpc_convertions.ProductPricesToProto(acc)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &product_prices_proto.ListResponse{
		Rows:  responseProductPrices,
		Count: uint32(count),
	}, nil
}
