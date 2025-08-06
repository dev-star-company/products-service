package tools_controller

import (
	"context"
	"products-service/internal/adapters/grpc_convertions"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/tools_proto"
)

func (c *controller) Create(ctx context.Context, in *tools_proto.CreateRequest) (*tools_proto.CreateResponse, error) {

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	create, err := c.Db.Tools.Create().
		SetName(in.Name).
		Save(ctx)

	if err != nil {
		return nil, errs.CreateError("product type", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &tools_proto.CreateResponse{
		Tools: grpc_convertions.ToolsToProto(create),
	}, nil
}
