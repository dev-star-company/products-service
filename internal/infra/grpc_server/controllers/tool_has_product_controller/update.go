package tool_has_product_controller

import (
	"context"
	"products-service/internal/app/ent"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/tool_has_product_proto"
)

func (c *controller) Update(ctx context.Context, in *tool_has_product_proto.UpdateRequest) (*tool_has_product_proto.UpdateResponse, error) {

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	var tool_has_product *ent.ToolHasProduct

	tool_has_productQ := tx.ToolHasProduct.UpdateOneID(int(in.Id))

	if in.ToolId != nil && *in.ToolId > 0 {
		tool_has_productQ.SetToolsID(int(*in.ToolId))
	}

	if in.ProductsId != nil && *in.ProductsId > 0 {
		tool_has_productQ.SetProductsID(int(*in.ProductsId))
	}

	tool_has_product, err = tool_has_productQ.Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, utils.Rollback(tx, errs.ProductsNotFound(int(in.Id)))
		}
		if ent.IsConstraintError(err) {
			return nil, utils.Rollback(tx, errs.InvalidForeignKey(err))
		}
		return nil, errs.SavingError("tool_has_product", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &tool_has_product_proto.UpdateResponse{
		ToolId:     uint32(*tool_has_product.ToolsID),
		ProductsId: uint32(*tool_has_product.ProductsID),
	}, nil
}
