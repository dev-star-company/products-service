package product_has_product_reference_controller

import (
	"context"
	"errors"
	"products-service/internal/adapters/grpc_convertions"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/producthasproductreference"
	"products-service/internal/app/ent/schema"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/product_has_product_reference_proto"
)

func (c *controller) List(ctx context.Context, in *product_has_product_reference_proto.ListRequest) (*product_has_product_reference_proto.ListResponse, error) {
	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	if in.IncludeDeleted != nil && *in.IncludeDeleted {
		ctx = schema.SkipSoftDelete(ctx)
	}

	query := tx.ProductHasProductReference.Query()

	if in.ProductReferenceId != nil {
		query = query.Where(producthasproductreference.ProductReferenceID(int(*in.ProductReferenceId)))
	}

	if in.ProductsId != nil {
		query = query.Where(producthasproductreference.ProductID(int(*in.ProductsId)))
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errs.ListingError("querying product_has_product_reference", err)
	}

	if in.Limit != nil && *in.Limit > 0 {
		query = query.Limit(int(*in.Limit))
	}

	if in.Offset != nil && *in.Offset > 0 {
		query = query.Offset(int(*in.Offset))
	}

	if in.Orderby != nil {
		if in.Orderby.Id != nil {
			switch *in.Orderby.Id {
			case "ASC":
				query = query.Order(ent.Asc(producthasproductreference.FieldID))
			case "DESC":
				query = query.Order(ent.Desc(producthasproductreference.FieldID))
			default:
				return nil, errs.InvalidOrderByValue(errors.New(*in.Orderby.Id))
			}
		}
	}

	product_has_product_reference, err := query.All(ctx)
	if err != nil {
		return nil, errs.ListingError("querying product_has_product_reference", err)
	}

	responseProductHasProductReference := make([]*product_has_product_reference_proto.ProductHasProductReference, len(product_has_product_reference))
	for i, acc := range product_has_product_reference {
		responseProductHasProductReference[i] = grpc_convertions.ProductHasProductReferenceToProto(acc)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &product_has_product_reference_proto.ListResponse{
		Rows:  responseProductHasProductReference,
		Count: uint32(count),
	}, nil
}
