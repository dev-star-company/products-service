package brand_controller

import (
	"context"
	"errors"
	"products-service/generated_protos/brand_proto"
	grpc_convertions "products-service/internal/adapters/grpc"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/brand"
	"products-service/internal/app/ent/schema"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"
)

func (c *controller) List(ctx context.Context, in *brand_proto.ListRequest) (*brand_proto.ListResponse, error) {
	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	if in.IncludeDeleted != nil && *in.IncludeDeleted {
		ctx = schema.SkipSoftDelete(ctx)
	}

	query := tx.Brand.Query()

	if in.Name != nil {
		query = query.Where(brand.Name(string(*in.Name)))
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errs.ListingError("querying brand", err)
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
				query = query.Order(ent.Asc(brand.FieldID))
			case "DESC":
				query = query.Order(ent.Desc(brand.FieldID))
			default:
				return nil, errs.InvalidOrderByValue(errors.New(*in.Orderby.Id))
			}
		}
	}

	brand, err := query.All(ctx)
	if err != nil {
		return nil, errs.ListingError("querying brand", err)
	}

	responseBrand := make([]*brand_proto.Brand, len(brand))
	for i, acc := range brand {
		responseBrand[i] = grpc_convertions.BrandToProto(acc)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &brand_proto.ListResponse{
		Rows:  responseBrand,
		Count: uint32(count),
	}, nil
}
