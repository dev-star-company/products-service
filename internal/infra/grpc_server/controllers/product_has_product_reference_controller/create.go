package product_has_product_reference_controller

import (
	"context"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/product_has_product_reference_proto"
)

func (c *controller) Create(ctx context.Context, in *product_has_product_reference_proto.CreateRequest) (*product_has_product_reference_proto.CreateResponse, error) {

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	create, err := c.Db.ProductHasProductReference.Create().
		SetProductReferenceID(int(in.ProductReferenceId)).
		SetProductID(int(in.ProductsId)).
		Save(ctx)

	if err != nil {
		return nil, errs.CreateError("product type", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &product_has_product_reference_proto.CreateResponse{
		ProductReferenceId: uint32(*create.ProductReferenceID),
		ProductsId:         uint32(*create.ProductID),
	}, nil
}
