package price_type_controller

import (
	"context"
	"errors"
	grpc_convertions "products-service/internal/adapters/grpc"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/pricetype"
	"products-service/internal/app/ent/schema"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/price_type_proto"
)

func (c *controller) List(ctx context.Context, in *price_type_proto.ListRequest) (*price_type_proto.ListResponse, error) {
	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	if in.IncludeDeleted != nil && *in.IncludeDeleted {
		ctx = schema.SkipSoftDelete(ctx)
	}

	query := tx.PriceType.Query()

	if in.Name != nil {
		query = query.Where(pricetype.Name(string(*in.Name)))
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errs.ListingError("querying price_type", err)
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
				query = query.Order(ent.Asc(pricetype.FieldID))
			case "DESC":
				query = query.Order(ent.Desc(pricetype.FieldID))
			default:
				return nil, errs.InvalidOrderByValue(errors.New(*in.Orderby.Id))
			}
		}
	}

	price_type, err := query.All(ctx)
	if err != nil {
		return nil, errs.ListingError("querying price_type", err)
	}

	responsePriceType := make([]*price_type_proto.PriceType, len(price_type))
	for i, acc := range price_type {
		responsePriceType[i] = grpc_convertions.PriceTypeToProto(acc)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &price_type_proto.ListResponse{
		Rows:  responsePriceType,
		Count: uint32(count),
	}, nil
}
