package promotions_controller

import (
	"context"
	"products-service/generated_protos/promotions_proto"
	"products-service/internal/app/ent"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"
	"time"
)

func (c *controller) Update(ctx context.Context, in *promotions_proto.UpdateRequest) (*promotions_proto.UpdateResponse, error) {

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	var promotions *ent.Promotions

	promotionsQ := tx.Promotions.UpdateOneID(int(in.Id))

	if in.Name != nil && *in.Name != "" {
		promotionsQ.SetName(string(*in.Name))
	}

	if in.StartingDatetime != nil && *in.StartingDatetime != "" {
		parsedTime, err := time.Parse(time.RFC3339, *in.StartingDatetime)
		if err != nil {
			return nil, errs.ProductsNotFound(int(in.Id))
		}
		promotionsQ.SetStartingDatetime(parsedTime)
	}

	if in.EndingDatetime != nil && *in.EndingDatetime != "" {
		parsedTime, err := time.Parse(time.RFC3339, *in.EndingDatetime)
		if err != nil {
			return nil, errs.ProductsNotFound(int(in.Id))
		}
		promotionsQ.SetEndingDatetime(parsedTime)
	}

	promotions, err = promotionsQ.Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, utils.Rollback(tx, errs.ProductsNotFound(int(in.Id)))
		}
		if ent.IsConstraintError(err) {
			return nil, utils.Rollback(tx, errs.InvalidForeignKey(err))
		}
		return nil, errs.SavingError("promotions", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &promotions_proto.UpdateResponse{

		Name: string(*promotions.Name),
	}, nil
}
