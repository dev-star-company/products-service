package brand_controller

import (
	"context"
	"products-service/generated_protos/brand_proto"
	"products-service/internal/app/ent"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"
)

func (c *controller) Update(ctx context.Context, in *brand_proto.UpdateRequest) (*brand_proto.UpdateResponse, error) {
	if in.RequesterId == 0 {
		return nil, errs.ProductsNotFound(int(in.Id))
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	var brand *ent.Brand

	brandQ := tx.Brand.UpdateOneID(int(in.Id))

	if in.Name != nil && *in.Name != "" {
		brandQ.SetName(string(*in.Name))
	}

	brandQ.SetUpdatedBy(int(in.RequesterId))

	brand, err = brandQ.Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, utils.Rollback(tx, errs.ProductsNotFound(int(in.Id)))
		}
		if ent.IsConstraintError(err) {
			return nil, utils.Rollback(tx, errs.InvalidForeignKey(err))
		}
		return nil, errs.SavingError("brand", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &brand_proto.UpdateResponse{
		RequesterId: uint32(brand.CreatedBy),
		Name:        string(*brand.Name),
	}, nil
}
