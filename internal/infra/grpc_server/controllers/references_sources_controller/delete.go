package references_sources_controller

import (
	"context"
	"products-service/generated_protos/references_sources_proto"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"
	"time"
)

func (c *controller) Delete(ctx context.Context, in *references_sources_proto.DeleteRequest) (*references_sources_proto.DeleteResponse, error) {
	if in.RequesterId == 0 {
		return nil, errs.ReferenceSourcesNotFound(int(in.Id))
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	err = tx.ReferenceSources.UpdateOneID(int(in.Id)).
		SetDeletedAt(time.Now()).
		SetDeletedBy(int(in.RequesterId)).
		Exec(ctx)

	if err != nil {
		return nil, utils.Rollback(tx, errs.DeleteError("references_sources", err))
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &references_sources_proto.DeleteResponse{}, nil
}
