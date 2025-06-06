package product_has_product_reference_controller

import (
	"context"
	"products-service/generated_protos/product_has_product_reference_proto"
	"products-service/internal/app/ent"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"
)

func (c *controller) Update(ctx context.Context, in *product_has_product_reference_proto.UpdateRequest) (*product_has_product_reference_proto.UpdateResponse, error) {
	if in.RequesterId == 0 {
		return nil, errs.ProductHasProductReferenceNotFound(int(in.Id))
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	var product_has_product_reference *ent.ProductHasProductReference

	product_has_product_referenceQ := tx.ProductHasProductReference.UpdateOneID(int(in.Id))

	if in.ProductReferenceId != nil && *in.ProductReferenceId > 0 {
		product_has_product_referenceQ.SetProductReferenceID(int(*in.ProductReferenceId))
	}

	if in.ProductsId != nil && *in.ProductsId > 0 {
		product_has_product_referenceQ.SetProductID(int(*in.ProductsId))
	}

	product_has_product_referenceQ.SetUpdatedBy(int(in.RequesterId))

	product_has_product_reference, err = product_has_product_referenceQ.Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, utils.Rollback(tx, errs.ProductsNotFound(int(in.Id)))
		}
		if ent.IsConstraintError(err) {
			return nil, utils.Rollback(tx, errs.InvalidForeignKey(err))
		}
		return nil, errs.SavingError("product_has_product_reference", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &product_has_product_reference_proto.UpdateResponse{
		RequesterId:        uint32(product_has_product_reference.CreatedBy),
		ProductReferenceId: uint32(*product_has_product_reference.ProductReferenceID),
		ProductsId:         uint32(*product_has_product_reference.ProductID),
	}, nil
}
