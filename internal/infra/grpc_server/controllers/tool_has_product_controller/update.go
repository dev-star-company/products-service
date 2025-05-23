package tool_has_product_controller

import (
	"context"
	"products-service/generated_protos/tool_has_product_proto"
	"products-service/internal/app/ent"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"
)

func (c *controller) Update(ctx context.Context, in *tool_has_product_proto.UpdateRequest) (*tool_has_product_proto.UpdateResponse, error) {
	if in.RequesterId == 0 {
		return nil, errs.ProductsNotFound(int(in.Id))
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	var tool_has_product *ent.ToolHasProduct

	tool_has_productQ := tx.ToolHasProduct.UpdateOneID(int(in.Id))

	if in.ToolId != nil && *in.ToolId > 0 {
		tool_has_productQ.SetToolID(int(*in.ToolId))
	}

	if in.ProductsId != nil && *in.ProductsId > 0 {
		tool_has_productQ.SetProductID(int(*in.ProductsId))
	}

	tool_has_productQ.SetUpdatedBy(int(in.RequesterId))

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
		RequesterId: uint32(tool_has_product.CreatedBy),
		ToolId:      uint32(*tool_has_product.ToolID),
		ProductsId:  uint32(*tool_has_product.ProductID),
	}, nil
}
