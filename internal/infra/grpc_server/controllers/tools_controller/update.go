package tools_controller

import (
	"context"
	"products-service/generated_protos/tools_proto"
	"products-service/internal/app/ent"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"
)

func (c *controller) Update(ctx context.Context, in *tools_proto.UpdateRequest) (*tools_proto.UpdateResponse, error) {
	if in.RequesterId == 0 {
		return nil, errs.ProductsNotFound(int(in.Id))
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	var tools *ent.Brand

	toolsQ := tx.Brand.UpdateOneID(int(in.Id))

	if in.Name != nil && *in.Name != "" {
		toolsQ.SetName(string(*in.Name))
	}

	toolsQ.SetUpdatedBy(int(in.RequesterId))

	tools, err = toolsQ.Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, utils.Rollback(tx, errs.ProductsNotFound(int(in.Id)))
		}
		if ent.IsConstraintError(err) {
			return nil, utils.Rollback(tx, errs.InvalidForeignKey(err))
		}
		return nil, errs.SavingError("tools", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &tools_proto.UpdateResponse{
		RequesterId: uint32(tools.CreatedBy),
		Name:        string(*tools.Name),
	}, nil
}
