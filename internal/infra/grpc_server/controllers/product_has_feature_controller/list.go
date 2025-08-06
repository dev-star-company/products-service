package product_has_feature_controller

import (
	"context"
	"errors"
	grpc_convertions "products-service/internal/adapters/grpc"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/producthasfeature"
	"products-service/internal/app/ent/schema"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/product_has_feature_proto"
)

func (c *controller) List(ctx context.Context, in *product_has_feature_proto.ListRequest) (*product_has_feature_proto.ListResponse, error) {
	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	if in.IncludeDeleted != nil && *in.IncludeDeleted {
		ctx = schema.SkipSoftDelete(ctx)
	}

	query := tx.ProductHasFeature.Query()

	if in.FeaturesId != nil {
		query = query.Where(producthasfeature.FeatureID(int(*in.FeaturesId)))
	}

	if in.ProductsId != nil {
		query = query.Where(producthasfeature.ProductID(int(*in.ProductsId)))
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errs.ListingError("querying product_has_feature", err)
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
				query = query.Order(ent.Asc(producthasfeature.FieldID))
			case "DESC":
				query = query.Order(ent.Desc(producthasfeature.FieldID))
			default:
				return nil, errs.InvalidOrderByValue(errors.New(*in.Orderby.Id))
			}
		}
	}

	product_has_feature, err := query.All(ctx)
	if err != nil {
		return nil, errs.ListingError("querying product_has_feature", err)
	}

	responseProductHasFeature := make([]*product_has_feature_proto.ProductHasFeature, len(product_has_feature))
	for i, acc := range product_has_feature {
		responseProductHasFeature[i] = grpc_convertions.ProductHasFeatureToProto(acc)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &product_has_feature_proto.ListResponse{
		Rows:  responseProductHasFeature,
		Count: uint32(count),
	}, nil
}
