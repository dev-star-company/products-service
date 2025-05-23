package tools_controller

import (
	"context"
	"products-service/generated_protos/tools_proto"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/tools"
	"products-service/internal/pkg/errs"
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
		RequesterId: uint32(tools.CreatedBy),
		Name:        *tools.Name,
	}, nil
}
