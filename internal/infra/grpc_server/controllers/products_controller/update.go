package products_controller

import (
	"context"
	"products-service/internal/adapters/grpc_convertions"
	"products-service/internal/app/ent"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/products_proto"
)

func (c *controller) Update(ctx context.Context, in *products_proto.UpdateRequest) (*products_proto.UpdateResponse, error) {

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	var products *ent.Products

	productsQ := tx.Products.UpdateOneID(int(in.Id))

	if in.CategoryId != nil && *in.CategoryId > 0 {
		productsQ.SetCategoryID(int(*in.CategoryId))
	}

	if in.BrandId != nil && *in.BrandId > 0 {
		productsQ.SetBrandID(int(*in.BrandId))
	}

	if in.VariantTypeId != nil && *in.VariantTypeId > 0 {
		productsQ.SetVariantTypeID(int(*in.VariantTypeId))
	}

	if in.ProductReferencesId != nil && *in.ProductReferencesId > 0 {
		productsQ.SetProductReferencesID(int(*in.ProductReferencesId))
	}

	if in.Name != nil && *in.Name != "" {
		productsQ.SetName(string(*in.Name))
	}

	if in.Stock != nil && *in.Stock > 0 {
		productsQ.SetStock(int(*in.Stock))
	}

	products, err = productsQ.Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, utils.Rollback(tx, errs.ProductsNotFound(int(in.Id)))
		}
		if ent.IsConstraintError(err) {
			return nil, utils.Rollback(tx, errs.InvalidForeignKey(err))
		}
		return nil, errs.SavingError("products", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &products_proto.UpdateResponse{
		Products: grpc_convertions.ProductsToProto(products),
	}, nil
}
