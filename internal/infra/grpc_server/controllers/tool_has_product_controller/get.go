package tool_has_product_controller

import (
	"context"
	"products-service/generated_protos/tool_has_product_proto"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/toolhasproduct"
	"products-service/internal/pkg/errs"
)

func (c *controller) Get(ctx context.Context, in *tool_has_product_proto.GetRequest) (*tool_has_product_proto.GetResponse, error) {
	tool_has_product, err := c.Db.ToolHasProduct.
		Query().
		Where(toolhasproduct.ID(int(in.Id))).
		WithTool().
		WithProduct().
		Only(ctx)

	if ent.IsNotFound(err) {
		return nil, errs.ToolHasProductNotFound(int(in.Id))
	}

	return &tool_has_product_proto.GetResponse{
		RequesterId: uint32(tool_has_product.CreatedBy),
		ToolId:      uint32(*tool_has_product.ToolID),
		ProductsId:  uint32(*tool_has_product.ProductID),
	}, nil
}
