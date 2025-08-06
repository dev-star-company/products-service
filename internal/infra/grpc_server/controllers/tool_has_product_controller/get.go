package tool_has_product_controller

import (
	"context"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/toolhasproduct"
	"products-service/internal/pkg/errs"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/tool_has_product_proto"
)

func (c *controller) Get(ctx context.Context, in *tool_has_product_proto.GetRequest) (*tool_has_product_proto.GetResponse, error) {
	tool_has_product, err := c.Db.ToolHasProduct.
		Query().
		Where(toolhasproduct.ID(int(in.Id))).
		WithTools().
		WithProducts().
		Only(ctx)

	if ent.IsNotFound(err) {
		return nil, errs.ToolHasProductNotFound(int(in.Id))
	}

	return &tool_has_product_proto.GetResponse{
		ToolId:     uint32(*tool_has_product.ToolsID),
		ProductsId: uint32(*tool_has_product.ProductsID),
	}, nil
}
