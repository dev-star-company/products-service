package price_type_controller

import (
	"context"
	"products-service/generated_protos/price_type_proto"
	"products-service/internal/app/ent"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"
)

func (c *controller) Update(ctx context.Context, in *price_type_proto.UpdateRequest) (*price_type_proto.UpdateResponse, error) {
	if in.RequesterId == 0 {
		return nil, errs.ProductsNotFound(int(in.Id))
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	var price_type *ent.PriceType

	price_typeQ := tx.PriceType.UpdateOneID(int(in.Id))

	if in.Name != nil && *in.Name != "" {
		price_typeQ.SetName(string(*in.Name))
	}

	price_typeQ.SetUpdatedBy(int(in.RequesterId))

	price_type, err = price_typeQ.Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, utils.Rollback(tx, errs.ProductsNotFound(int(in.Id)))
		}
		if ent.IsConstraintError(err) {
			return nil, utils.Rollback(tx, errs.InvalidForeignKey(err))
		}
		return nil, errs.SavingError("price_type", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &price_type_proto.UpdateResponse{
		RequesterId: uint32(price_type.CreatedBy),
		Name:        string(*price_type.Name),
	}, nil
}
