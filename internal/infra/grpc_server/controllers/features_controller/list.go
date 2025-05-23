package features_controller

import (
	"context"
	"errors"
	"products-service/generated_protos/features_proto"
	grpc_convertions "products-service/internal/adapters/grpc"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/features"
	"products-service/internal/app/ent/schema"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"
)

func (c *controller) List(ctx context.Context, in *features_proto.ListRequest) (*features_proto.ListResponse, error) {
	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	if in.IncludeDeleted != nil && *in.IncludeDeleted {
		ctx = schema.SkipSoftDelete(ctx)
	}

	query := tx.Features.Query()

	if in.FeatureValueId != nil {
		query = query.Where(features.FeatureValueID(int(*in.FeatureValueId)))
	}

	if in.Name != nil {
		query = query.Where(features.Name(string(*in.Name)))
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errs.ListingError("querying features", err)
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
				query = query.Order(ent.Asc(features.FieldID))
			case "DESC":
				query = query.Order(ent.Desc(features.FieldID))
			default:
				return nil, errs.InvalidOrderByValue(errors.New(*in.Orderby.Id))
			}
		}
	}

	features, err := query.All(ctx)
	if err != nil {
		return nil, errs.ListingError("querying features", err)
	}

	responseFeatures := make([]*features_proto.Features, len(features))
	for i, acc := range features {
		responseFeatures[i] = grpc_convertions.FeaturesToProto(acc)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &features_proto.ListResponse{
		Rows:  responseFeatures,
		Count: uint32(count),
	}, nil
}
