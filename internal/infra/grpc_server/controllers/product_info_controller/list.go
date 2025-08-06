package product_info_controller

import (
	"context"
	"errors"
	grpc_convertions "products-service/internal/adapters/grpc"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/productinfo"
	"products-service/internal/app/ent/schema"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/product_info_proto"
)

func (c *controller) List(ctx context.Context, in *product_info_proto.ListRequest) (*product_info_proto.ListResponse, error) {
	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	if in.IncludeDeleted != nil && *in.IncludeDeleted {
		ctx = schema.SkipSoftDelete(ctx)
	}

	query := tx.ProductInfo.Query()

	if in.InfoTypeId != nil {
		query = query.Where(productinfo.InfoTypesID(int(*in.InfoTypeId)))
	}

	if in.Value != nil {
		query = query.Where(productinfo.Value(string(*in.Value)))
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errs.ListingError("querying product_info", err)
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
				query = query.Order(ent.Asc(productinfo.FieldID))
			case "DESC":
				query = query.Order(ent.Desc(productinfo.FieldID))
			default:
				return nil, errs.InvalidOrderByValue(errors.New(*in.Orderby.Id))
			}
		}
	}

	product_info, err := query.All(ctx)
	if err != nil {
		return nil, errs.ListingError("querying product_info", err)
	}

	responseProductInfo := make([]*product_info_proto.ProductInfo, len(product_info))
	for i, acc := range product_info {
		responseProductInfo[i] = grpc_convertions.ProductInfoToProto(acc)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &product_info_proto.ListResponse{
		Rows:  responseProductInfo,
		Count: uint32(count),
	}, nil
}
