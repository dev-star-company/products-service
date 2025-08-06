package product_references_controller

import (
	"context"
	"products-service/internal/adapters/grpc_convertions"
	"products-service/internal/app/ent"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/product_references_proto"
)

func (c *controller) Update(ctx context.Context, in *product_references_proto.UpdateRequest) (*product_references_proto.UpdateResponse, error) {

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	var product_references *ent.ProductReferences

	product_referencesQ := tx.ProductReferences.UpdateOneID(int(in.Id))

	if in.ReferenceSourceId != nil && *in.ReferenceSourceId > 0 {
		product_referencesQ.SetReferenceSourceID(int(*in.ReferenceSourceId))
	}

	if in.Value != nil && *in.Value != "" {
		product_referencesQ.SetValue(string(*in.Value))
	}

	product_references, err = product_referencesQ.Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, utils.Rollback(tx, errs.ProductsNotFound(int(in.Id)))
		}
		if ent.IsConstraintError(err) {
			return nil, utils.Rollback(tx, errs.InvalidForeignKey(err))
		}
		return nil, errs.SavingError("product_references", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &product_references_proto.UpdateResponse{
		Productreferences: grpc_convertions.ProductReferencesToProto(product_references),
	}, nil
}
