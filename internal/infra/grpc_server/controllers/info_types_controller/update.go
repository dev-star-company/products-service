package info_types_controller

import (
	"context"
	"products-service/generated_protos/info_types_proto"
	"products-service/internal/app/ent"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"
)

func (c *controller) Update(ctx context.Context, in *info_types_proto.UpdateRequest) (*info_types_proto.UpdateResponse, error) {
	if in.RequesterId == 0 {
		return nil, errs.ProductsNotFound(int(in.Id))
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	var info_types *ent.InfoTypes

	info_typesQ := tx.InfoTypes.UpdateOneID(int(in.Id))

	if in.Name != nil && *in.Name != "" {
		info_typesQ.SetName(string(*in.Name))
	}

	info_typesQ.SetUpdatedBy(int(in.RequesterId))

	info_types, err = info_typesQ.Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, utils.Rollback(tx, errs.ProductsNotFound(int(in.Id)))
		}
		if ent.IsConstraintError(err) {
			return nil, utils.Rollback(tx, errs.InvalidForeignKey(err))
		}
		return nil, errs.SavingError("info_types", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &info_types_proto.UpdateResponse{
		RequesterId: uint32(info_types.CreatedBy),
		Name:        string(*info_types.Name),
	}, nil
}
