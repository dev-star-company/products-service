package features_values_types_controller

import (
	"context"
	"products-service/generated_protos/features_values_types_proto"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"
	"time"
)

func (c *controller) Delete(ctx context.Context, in *features_values_types_proto.DeleteRequest) (*features_values_types_proto.DeleteResponse, error) {

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	err = tx.FeaturesValuesTypes.UpdateOneID(int(in.Id)).
		SetDeletedAt(time.Now()).
		Exec(ctx)

	if err != nil {
		return nil, utils.Rollback(tx, errs.DeleteError("features_values_types", err))
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &features_values_types_proto.DeleteResponse{}, nil
}
