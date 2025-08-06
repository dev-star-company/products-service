package tool_has_product_controller

import (
	"context"
	"errors"
	"products-service/internal/adapters/grpc_convertions"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/schema"
	"products-service/internal/app/ent/toolhasproduct"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/tool_has_product_proto"
)

func (c *controller) List(ctx context.Context, in *tool_has_product_proto.ListRequest) (*tool_has_product_proto.ListResponse, error) {
	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	if in.IncludeDeleted != nil && *in.IncludeDeleted {
		ctx = schema.SkipSoftDelete(ctx)
	}

	query := tx.ToolHasProduct.Query()

	if in.ToolId != nil {
		query = query.Where(toolhasproduct.ToolsID(int(*in.ToolId)))
	}

	if in.ProductsId != nil {
		query = query.Where(toolhasproduct.ProductsID(int(*in.ProductsId)))
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errs.ListingError("querying tool_has_product", err)
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
				query = query.Order(ent.Asc(toolhasproduct.FieldID))
			case "DESC":
				query = query.Order(ent.Desc(toolhasproduct.FieldID))
			default:
				return nil, errs.InvalidOrderByValue(errors.New(*in.Orderby.Id))
			}
		}
	}

	tool_has_product, err := query.All(ctx)
	if err != nil {
		return nil, errs.ListingError("querying tool_has_product", err)
	}

	responseToolHasProduct := make([]*tool_has_product_proto.ToolHasProduct, len(tool_has_product))
	for i, acc := range tool_has_product {
		responseToolHasProduct[i] = grpc_convertions.ToolHasProductToProto(acc)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &tool_has_product_proto.ListResponse{
		Rows:  responseToolHasProduct,
		Count: uint32(count),
	}, nil
}
