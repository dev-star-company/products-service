package product_references_controller

import (
	"context"
	"errors"
	"products-service/generated_protos/product_references_proto"
	grpc_convertions "products-service/internal/adapters/grpc"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/productreferences"
	"products-service/internal/app/ent/schema"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"
)

func (c *controller) List(ctx context.Context, in *product_references_proto.ListRequest) (*product_references_proto.ListResponse, error) {
	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	if in.IncludeDeleted != nil && *in.IncludeDeleted {
		ctx = schema.SkipSoftDelete(ctx)
	}

	query := tx.ProductReferences.Query()

	if in.ReferenceSourceId != nil {
		query = query.Where(productreferences.ReferenceSourceID(int(*in.ReferenceSourceId)))
	}

	if in.Value != nil {
		query = query.Where(productreferences.Value(*in.Value))
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errs.ListingError("querying product_references", err)
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
				query = query.Order(ent.Asc(productreferences.FieldID))
			case "DESC":
				query = query.Order(ent.Desc(productreferences.FieldID))
			default:
				return nil, errs.InvalidOrderByValue(errors.New(*in.Orderby.Id))
			}
		}
	}

	product_references, err := query.All(ctx)
	if err != nil {
		return nil, errs.ListingError("querying product_references", err)
	}

	responseProductReferences := make([]*product_references_proto.ProductReferences, len(product_references))
	for i, acc := range product_references {
		responseProductReferences[i] = grpc_convertions.ProductReferencesToProto(acc)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &product_references_proto.ListResponse{
		Rows:  responseProductReferences,
		Count: uint32(count),
	}, nil
}
