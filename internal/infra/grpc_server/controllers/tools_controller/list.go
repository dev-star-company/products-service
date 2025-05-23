package tools_controller

import (
	"context"
	"errors"
	"products-service/generated_protos/tools_proto"
	grpc_convertions "products-service/internal/adapters/grpc"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/schema"
	"products-service/internal/app/ent/tools"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"
)

func (c *controller) List(ctx context.Context, in *tools_proto.ListRequest) (*tools_proto.ListResponse, error) {
	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	if in.IncludeDeleted != nil && *in.IncludeDeleted {
		ctx = schema.SkipSoftDelete(ctx)
	}

	query := tx.Tools.Query()

	if in.Name != nil {
		query = query.Where(tools.Name(string(*in.Name)))
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errs.ListingError("querying tools", err)
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
				query = query.Order(ent.Asc(tools.FieldID))
			case "DESC":
				query = query.Order(ent.Desc(tools.FieldID))
			default:
				return nil, errs.InvalidOrderByValue(errors.New(*in.Orderby.Id))
			}
		}
	}

	tools, err := query.All(ctx)
	if err != nil {
		return nil, errs.ListingError("querying tools", err)
	}

	responseTools := make([]*tools_proto.Tools, len(tools))
	for i, acc := range tools {
		responseTools[i] = grpc_convertions.ToolsToProto(acc)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &tools_proto.ListResponse{
		Rows:  responseTools,
		Count: uint32(count),
	}, nil
}
