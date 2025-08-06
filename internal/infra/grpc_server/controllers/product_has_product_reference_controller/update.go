package product_has_product_reference_controller

import (
	"context"
	"products-service/internal/adapters/grpc_convertions"
	"products-service/internal/app/ent"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/product_has_product_reference_proto"
)

func (c *controller) Update(ctx context.Context, in *product_has_product_reference_proto.UpdateRequest) (*product_has_product_reference_proto.UpdateResponse, error) {

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
		Producthasproductreference: grpc_convertions.ProductHasProductReferenceToProto(product_has_product_reference),
	}, nil
}
