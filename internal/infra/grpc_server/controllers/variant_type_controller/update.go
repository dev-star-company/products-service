package variant_type_controller

import (
	"context"
	"products-service/generated_protos/variant_type_proto"
	"products-service/internal/app/ent"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"
)

func (c *controller) Update(ctx context.Context, in *variant_type_proto.UpdateRequest) (*variant_type_proto.UpdateResponse, error) {
	if in.RequesterId == 0 {
		return nil, errs.ProductsNotFound(int(in.Id))
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	var variant_type *ent.VariantType

	variant_typeQ := tx.VariantType.UpdateOneID(int(in.Id))

	if in.Name != nil && *in.Name != "" {
		variant_typeQ.SetName(string(*in.Name))
	}

	variant_typeQ.SetUpdatedBy(int(in.RequesterId))

	variant_type, err = variant_typeQ.Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, utils.Rollback(tx, errs.ProductsNotFound(int(in.Id)))
		}
		if ent.IsConstraintError(err) {
			return nil, utils.Rollback(tx, errs.InvalidForeignKey(err))
		}
		return nil, errs.SavingError("variant_type", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &variant_type_proto.UpdateResponse{
		RequesterId: uint32(variant_type.CreatedBy),
		Name:        string(*variant_type.Name),
	}, nil
}
