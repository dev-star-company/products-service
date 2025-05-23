package features_unit_values_controller

import (
	"context"
	"errors"
	"products-service/generated_protos/features_unit_values_proto"
	grpc_convertions "products-service/internal/adapters/grpc"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/featuresunitvalues"
	"products-service/internal/app/ent/schema"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"
)

func (c *controller) List(ctx context.Context, in *features_unit_values_proto.ListRequest) (*features_unit_values_proto.ListResponse, error) {
	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	if in.IncludeDeleted != nil && *in.IncludeDeleted {
		ctx = schema.SkipSoftDelete(ctx)
	}

	query := tx.FeaturesUnitValues.Query()

	if in.Decimals != nil {
		query = query.Where(featuresunitvalues.Decimals(int(*in.Decimals)))
	}

	if in.Name != nil {
		query = query.Where(featuresunitvalues.Name(string(*in.Name)))
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errs.ListingError("querying features_unit_values", err)
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
				query = query.Order(ent.Asc(featuresunitvalues.FieldID))
			case "DESC":
				query = query.Order(ent.Desc(featuresunitvalues.FieldID))
			default:
				return nil, errs.InvalidOrderByValue(errors.New(*in.Orderby.Id))
			}
		}
	}

	features_unit_values, err := query.All(ctx)
	if err != nil {
		return nil, errs.ListingError("querying features_unit_values", err)
	}

	responseFeaturesUnitValues := make([]*features_unit_values_proto.FeaturesUnitValues, len(features_unit_values))
	for i, acc := range features_unit_values {
		responseFeaturesUnitValues[i] = grpc_convertions.FeaturesUnitValuesToProto(acc)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &features_unit_values_proto.ListResponse{
		Rows:  responseFeaturesUnitValues,
		Count: uint32(count),
	}, nil
}
