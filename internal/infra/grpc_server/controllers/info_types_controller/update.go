package info_types_controller

import (
	"context"
	"products-service/internal/app/ent"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/info_types_proto"
)

func (c *controller) Update(ctx context.Context, in *info_types_proto.UpdateRequest) (*info_types_proto.UpdateResponse, error) {

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	var info_types *ent.InfoTypes

	info_typesQ := tx.InfoTypes.UpdateOneID(int(in.Id))

	if in.Name != nil && *in.Name != "" {
		info_typesQ.SetName(string(*in.Name))
	}

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

		Name: string(*info_types.Name),
	}, nil
}
