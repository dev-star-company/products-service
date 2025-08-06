package features_unit_values_controller

import (
	"context"
	"products-service/internal/app/ent"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/features_unit_values_proto"
)

func (c *controller) Update(ctx context.Context, in *features_unit_values_proto.UpdateRequest) (*features_unit_values_proto.UpdateResponse, error) {

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	var features_unit_values *ent.FeaturesUnitValues

	features_unit_valuesQ := tx.FeaturesUnitValues.UpdateOneID(int(in.Id))

	if in.Decimals != nil && *in.Decimals > 0 {
		features_unit_valuesQ.SetDecimals(int(*in.Decimals))
	}

	if in.Name != nil && *in.Name != "" {
		features_unit_valuesQ.SetName(string(*in.Name))
	}

	features_unit_values, err = features_unit_valuesQ.Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, utils.Rollback(tx, errs.ProductsNotFound(int(in.Id)))
		}
		if ent.IsConstraintError(err) {
			return nil, utils.Rollback(tx, errs.InvalidForeignKey(err))
		}
		return nil, errs.SavingError("features_unit_values", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &features_unit_values_proto.UpdateResponse{

		Decimals: float32(*features_unit_values.Decimals),
		Name:     string(*features_unit_values.Name),
	}, nil
}
