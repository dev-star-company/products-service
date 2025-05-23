package product_info_controller

import (
	"context"
	"products-service/generated_protos/product_info_proto"
	"products-service/internal/app/ent"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"
)

func (c *controller) Update(ctx context.Context, in *product_info_proto.UpdateRequest) (*product_info_proto.UpdateResponse, error) {
	if in.RequesterId == 0 {
		return nil, errs.ProductsNotFound(int(in.Id))
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	var product_info *ent.ProductInfo

	product_infoQ := tx.ProductInfo.UpdateOneID(int(in.Id))

	if in.InfoTypeId != nil && *in.InfoTypeId > 0 {
		product_infoQ.SetInfoTypeID(int(*in.InfoTypeId))
	}

	if in.Value != nil && *in.Value != "" {
		product_infoQ.SetValue(string(*in.Value))
	}

	product_infoQ.SetUpdatedBy(int(in.RequesterId))

	product_info, err = product_infoQ.Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, utils.Rollback(tx, errs.ProductsNotFound(int(in.Id)))
		}
		if ent.IsConstraintError(err) {
			return nil, utils.Rollback(tx, errs.InvalidForeignKey(err))
		}
		return nil, errs.SavingError("product_info", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &product_info_proto.UpdateResponse{
		RequesterId: uint32(product_info.CreatedBy),
		InfoTypeId:  uint32(*product_info.InfoTypesID),
		Value:       string(*product_info.Value),
	}, nil
}
