package products_controller

import (
	"context"
	"errors"
	"products-service/generated_protos/products_proto"
	grpc_convertions "products-service/internal/adapters/grpc"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/products"
	"products-service/internal/app/ent/schema"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"
)

func (c *controller) List(ctx context.Context, in *products_proto.ListRequest) (*products_proto.ListResponse, error) {
	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	if in.IncludeDeleted != nil && *in.IncludeDeleted {
		ctx = schema.SkipSoftDelete(ctx)
	}

	query := tx.Products.Query()

	if in.CategoryId != nil {
		query = query.Where(products.CategoryID(int(*in.CategoryId)))
	}

	if in.BrandId != nil {
		query = query.Where(products.BrandID(int(*in.BrandId)))
	}

	if in.VariantTypeId != nil {
		query = query.Where(products.VariantTypeID(int(*in.VariantTypeId)))
	}

	if in.ProductReferencesId != nil {
		query = query.Where(products.ProductReferencesID(int(*in.ProductReferencesId)))
	}

	if in.ImageId != nil {
		query = query.Where(products.ImageID(int(*in.ImageId)))
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errs.ListingError("querying products", err)
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
				query = query.Order(ent.Asc(products.FieldID))
			case "DESC":
				query = query.Order(ent.Desc(products.FieldID))
			default:
				return nil, errs.InvalidOrderByValue(errors.New(*in.Orderby.Id))
			}
		}
	}

	products, err := query.All(ctx)
	if err != nil {
		return nil, errs.ListingError("querying products", err)
	}

	responseProducts := make([]*products_proto.Products, len(products))
	for i, acc := range products {
		responseProducts[i] = grpc_convertions.ProductsToProto(acc)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &products_proto.ListResponse{
		Rows:  responseProducts,
		Count: uint32(count),
	}, nil
}
