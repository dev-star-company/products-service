package tools_controller

import (
	"context"
	"products-service/internal/adapters/grpc_convertions"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/tools"
	"products-service/internal/pkg/errs"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/tools_proto"
)

func (c *controller) Get(ctx context.Context, in *tools_proto.GetRequest) (*tools_proto.GetResponse, error) {
	tools, err := c.Db.Tools.
		Query().
		Where(tools.ID(int(in.Id))).
		Only(ctx)

	if ent.IsNotFound(err) {
		return nil, errs.ToolsNotFound(int(in.Id))
	}

	return &tools_proto.GetResponse{
		Tools: grpc_convertions.ToolsToProto(tools),
	}, nil
}
