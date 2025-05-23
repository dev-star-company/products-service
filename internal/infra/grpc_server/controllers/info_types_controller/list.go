package info_types_controller

import (
	"context"
	"errors"
	"products-service/generated_protos/info_types_proto"
	grpc_convertions "products-service/internal/adapters/grpc"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/infotypes"
	"products-service/internal/app/ent/schema"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"
)

func (c *controller) List(ctx context.Context, in *info_types_proto.ListRequest) (*info_types_proto.ListResponse, error) {
	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	if in.IncludeDeleted != nil && *in.IncludeDeleted {
		ctx = schema.SkipSoftDelete(ctx)
	}

	query := tx.InfoTypes.Query()

	if in.Name != nil {
		query = query.Where(infotypes.Name(string(*in.Name)))
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errs.ListingError("querying info_types", err)
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
				query = query.Order(ent.Asc(infotypes.FieldID))
			case "DESC":
				query = query.Order(ent.Desc(infotypes.FieldID))
			default:
				return nil, errs.InvalidOrderByValue(errors.New(*in.Orderby.Id))
			}
		}
	}

	info_types, err := query.All(ctx)
	if err != nil {
		return nil, errs.ListingError("querying info_types", err)
	}

	responseInfoTypes := make([]*info_types_proto.InfoTypes, len(info_types))
	for i, acc := range info_types {
		responseInfoTypes[i] = grpc_convertions.InfoTypesToProto(acc)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &info_types_proto.ListResponse{
		Rows:  responseInfoTypes,
		Count: uint32(count),
	}, nil
}
