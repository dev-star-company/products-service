package features_values_types_controller

import (
	"context"
	"products-service/internal/adapters/grpc_convertions"
	"products-service/internal/app/ent"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/features_values_types_proto"
)

func (c *controller) Update(ctx context.Context, in *features_values_types_proto.UpdateRequest) (*features_values_types_proto.UpdateResponse, error) {

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	var features_values_types *ent.FeaturesValuesTypes

	features_values_typesQ := tx.FeaturesValuesTypes.UpdateOneID(int(in.Id))

	if in.Name != nil && *in.Name != "" {
		features_values_typesQ.SetName(string(*in.Name))
	}

	features_values_types, err = features_values_typesQ.Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, utils.Rollback(tx, errs.ProductsNotFound(int(in.Id)))
		}
		if ent.IsConstraintError(err) {
			return nil, utils.Rollback(tx, errs.InvalidForeignKey(err))
		}
		return nil, errs.SavingError("features_values_types", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &features_values_types_proto.UpdateResponse{
		Featuresvaluestypes: grpc_convertions.FeaturesValuesTypesToProto(features_values_types),
	}, nil
}
