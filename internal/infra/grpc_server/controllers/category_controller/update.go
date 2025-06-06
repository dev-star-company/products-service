package category_controller

import (
	"context"
	"products-service/generated_protos/category_proto"
	"products-service/internal/app/ent"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"
)

func (c *controller) Update(ctx context.Context, in *category_proto.UpdateRequest) (*category_proto.UpdateResponse, error) {
	if in.RequesterId == 0 {
		return nil, errs.ProductsNotFound(int(in.Id))
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	var category *ent.Category

	categoryQ := tx.Category.UpdateOneID(int(in.Id))

	if in.Name != nil && *in.Name != "" {
		categoryQ.SetName(string(*in.Name))
	}

	categoryQ.SetUpdatedBy(int(in.RequesterId))

	category, err = categoryQ.Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, utils.Rollback(tx, errs.ProductsNotFound(int(in.Id)))
		}
		if ent.IsConstraintError(err) {
			return nil, utils.Rollback(tx, errs.InvalidForeignKey(err))
		}
		return nil, errs.SavingError("category", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &category_proto.UpdateResponse{
		RequesterId: uint32(category.CreatedBy),
		Name:        string(*category.Name),
	}, nil
}
