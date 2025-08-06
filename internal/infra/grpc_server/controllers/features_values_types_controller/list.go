package features_values_types_controller

import (
	"context"
	"errors"
	grpc_convertions "products-service/internal/adapters/grpc"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/featuresvaluestypes"
	"products-service/internal/app/ent/schema"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/features_values_types_proto"
)

func (c *controller) List(ctx context.Context, in *features_values_types_proto.ListRequest) (*features_values_types_proto.ListResponse, error) {
	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	if in.IncludeDeleted != nil && *in.IncludeDeleted {
		ctx = schema.SkipSoftDelete(ctx)
	}

	query := tx.FeaturesValuesTypes.Query()

	if in.Name != nil {
		query = query.Where(featuresvaluestypes.Name(string(*in.Name)))
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errs.ListingError("querying features_values_types", err)
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
				query = query.Order(ent.Asc(featuresvaluestypes.FieldID))
			case "DESC":
				query = query.Order(ent.Desc(featuresvaluestypes.FieldID))
			default:
				return nil, errs.InvalidOrderByValue(errors.New(*in.Orderby.Id))
			}
		}
	}

	features_values_types, err := query.All(ctx)
	if err != nil {
		return nil, errs.ListingError("querying features_values_types", err)
	}

	responseFeaturesValuesTypes := make([]*features_values_types_proto.FeaturesValuesTypes, len(features_values_types))
	for i, acc := range features_values_types {
		responseFeaturesValuesTypes[i] = grpc_convertions.FeaturesValuesTypesToProto(acc)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &features_values_types_proto.ListResponse{
		Rows:  responseFeaturesValuesTypes,
		Count: uint32(count),
	}, nil
}
