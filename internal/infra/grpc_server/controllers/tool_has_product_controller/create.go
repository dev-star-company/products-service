package tool_has_product_controller

import (
	"context"
	"products-service/generated_protos/tool_has_product_proto"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"
)

func (c *controller) Create(ctx context.Context, in *tool_has_product_proto.CreateRequest) (*tool_has_product_proto.CreateResponse, error) {

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	create, err := c.Db.ToolHasProduct.Create().
		SetToolsID(int(in.ToolId)).
		SetProductsID(int(in.ProductsId)).
		Save(ctx)

	if err != nil {
		return nil, errs.CreateError("product type", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &tool_has_product_proto.CreateResponse{
		ToolId:     uint32(*create.ToolsID),
		ProductsId: uint32(*create.ProductsID),
	}, nil
}
