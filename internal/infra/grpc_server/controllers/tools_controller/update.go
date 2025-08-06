package tools_controller

import (
	"context"
	"products-service/internal/app/ent"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/tools_proto"
)

func (c *controller) Update(ctx context.Context, in *tools_proto.UpdateRequest) (*tools_proto.UpdateResponse, error) {

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	var tools *ent.Brand

	toolsQ := tx.Brand.UpdateOneID(int(in.Id))

	if in.Name != nil && *in.Name != "" {
		toolsQ.SetName(string(*in.Name))
	}

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
		Name: string(*tools.Name),
	}, nil
}
