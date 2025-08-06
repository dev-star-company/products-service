package features_values_controller

import (
	"context"
	"errors"
	"products-service/internal/adapters/grpc_convertions"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/featuresvalues"
	"products-service/internal/app/ent/schema"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/features_values_proto"
)

func (c *controller) List(ctx context.Context, in *features_values_proto.ListRequest) (*features_values_proto.ListResponse, error) {
	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	if in.IncludeDeleted != nil && *in.IncludeDeleted {
		ctx = schema.SkipSoftDelete(ctx)
	}

	query := tx.FeaturesValues.Query()

	if in.FeatureId != 0 {
		query = query.Where(featuresvalues.FeatureID(int(in.FeatureId)))
	}

	if in.FeatureUnitValuesId != 0 {
		query = query.Where(featuresvalues.FeatureUnitValuesID(int(in.FeatureUnitValuesId)))
	}

	if in.Value != nil {
		query = query.Where(featuresvalues.Value(string(*in.Value)))
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errs.ListingError("querying features_values", err)
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
				query = query.Order(ent.Asc(featuresvalues.FieldID))
			case "DESC":
				query = query.Order(ent.Desc(featuresvalues.FieldID))
			default:
				return nil, errs.InvalidOrderByValue(errors.New(*in.Orderby.Id))
			}
		}
	}

	features_values, err := query.All(ctx)
	if err != nil {
		return nil, errs.ListingError("querying features_values", err)
	}

	responseFeaturesValues := make([]*features_values_proto.FeaturesValues, len(features_values))
	for i, acc := range features_values {
		responseFeaturesValues[i] = grpc_convertions.FeaturesValuesToProto(acc)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &features_values_proto.ListResponse{
		Rows:  responseFeaturesValues,
		Count: uint32(count),
	}, nil
}
