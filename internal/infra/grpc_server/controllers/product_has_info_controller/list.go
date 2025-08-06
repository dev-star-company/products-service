package product_has_info_controller

import (
	"context"
	"errors"
	grpc_convertions "products-service/internal/adapters/grpc"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/producthasinfo"
	"products-service/internal/app/ent/schema"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/product_has_info_proto"
)

func (c *controller) List(ctx context.Context, in *product_has_info_proto.ListRequest) (*product_has_info_proto.ListResponse, error) {
	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	if in.IncludeDeleted != nil && *in.IncludeDeleted {
		ctx = schema.SkipSoftDelete(ctx)
	}

	query := tx.ProductHasInfo.Query()

	if in.ProductInfoId != nil {
		query = query.Where(producthasinfo.ProductInfoID(int(*in.ProductInfoId)))
	}

	if in.ProductsId != nil {
		query = query.Where(producthasinfo.ProductID(int(*in.ProductsId)))
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errs.ListingError("querying product_has_info", err)
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
				query = query.Order(ent.Asc(producthasinfo.FieldID))
			case "DESC":
				query = query.Order(ent.Desc(producthasinfo.FieldID))
			default:
				return nil, errs.InvalidOrderByValue(errors.New(*in.Orderby.Id))
			}
		}
	}

	product_has_info, err := query.All(ctx)
	if err != nil {
		return nil, errs.ListingError("querying product_has_info", err)
	}

	responseProductHasInfo := make([]*product_has_info_proto.ProductHasInfo, len(product_has_info))
	for i, acc := range product_has_info {
		responseProductHasInfo[i] = grpc_convertions.ProductHasInfoToProto(acc)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &product_has_info_proto.ListResponse{
		Rows:  responseProductHasInfo,
		Count: uint32(count),
	}, nil
}
