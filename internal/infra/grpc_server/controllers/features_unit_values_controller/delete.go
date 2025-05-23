package features_unit_values_controller

import (
	"context"
	"products-service/generated_protos/features_unit_values_proto"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"
	"time"
)

func (c *controller) Delete(ctx context.Context, in *features_unit_values_proto.DeleteRequest) (*features_unit_values_proto.DeleteResponse, error) {
	if in.RequesterId == 0 {
		return nil, errs.FeaturesNotFound(int(in.Id))
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	err = tx.Features.UpdateOneID(int(in.Id)).
		SetDeletedAt(time.Now()).
		SetDeletedBy(int(in.RequesterId)).
		Exec(ctx)

	if err != nil {
		return nil, utils.Rollback(tx, errs.DeleteError("features_unit_values", err))
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &features_unit_values_proto.DeleteResponse{}, nil
}
