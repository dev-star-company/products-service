package product_has_info_controller

import (
	"context"
	"products-service/generated_protos/product_has_info_proto"
	"products-service/internal/app/ent"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"
)

func (c *controller) Update(ctx context.Context, in *product_has_info_proto.UpdateRequest) (*product_has_info_proto.UpdateResponse, error) {
	if in.RequesterId == 0 {
		return nil, errs.ProductHasInfoNotFound(int(in.Id))
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	var product_has_info *ent.ProductHasInfo

	product_has_infoQ := tx.ProductHasInfo.UpdateOneID(int(in.Id))

	if in.ProductInfoId != nil && *in.ProductInfoId > 0 {
		product_has_infoQ.SetProductInfoID(int(*in.ProductInfoId))
	}

	if in.ProductsId != nil && *in.ProductsId > 0 {
		product_has_infoQ.SetProductID(int(*in.ProductsId))
	}

	product_has_infoQ.SetUpdatedBy(int(in.RequesterId))

	product_has_info, err = product_has_infoQ.Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, utils.Rollback(tx, errs.ProductsNotFound(int(in.Id)))
		}
		if ent.IsConstraintError(err) {
			return nil, utils.Rollback(tx, errs.InvalidForeignKey(err))
		}
		return nil, errs.SavingError("product_has_info", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &product_has_info_proto.UpdateResponse{
		RequesterId:   uint32(product_has_info.CreatedBy),
		ProductInfoId: uint32(*product_has_info.ProductInfoID),
		ProductsId:    uint32(*product_has_info.ProductID),
	}, nil
}
