package products_controller

import (
	"context"
	"products-service/internal/adapters/grpc_convertions"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/products_proto"
)

func (c *controller) Create(ctx context.Context, in *products_proto.CreateRequest) (*products_proto.CreateResponse, error) {

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	create, err := c.Db.Products.Create().
		SetCategoryID(int(*in.CategoryId)).
		SetBrandID(int(*in.BrandId)).
		SetVariantTypeID(int(*in.VariantTypeId)).
		SetProductReferencesID(int(*in.ProductReferencesId)).
		// SetImagesID(int(*in.ImageId)).
		SetName(in.Name).
		SetStock(int(in.Stock)).
		Save(ctx)

	if err != nil {
		return nil, errs.CreateError("product type", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &products_proto.CreateResponse{
		Products: grpc_convertions.ProductsToProto(create),
	}, nil
}
