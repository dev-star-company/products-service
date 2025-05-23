package references_sources_controller

import (
	"context"
	"products-service/generated_protos/references_sources_proto"
	"products-service/internal/app/ent"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"
)

func (c *controller) Update(ctx context.Context, in *references_sources_proto.UpdateRequest) (*references_sources_proto.UpdateResponse, error) {
	if in.RequesterId == 0 {
		return nil, errs.ProductsNotFound(int(in.Id))
	}

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	var references_sources *ent.ReferenceSources

	references_sourcesQ := tx.ReferenceSources.UpdateOneID(int(in.Id))

	if in.Name != nil && *in.Name != "" {
		references_sourcesQ.SetName(string(*in.Name))
	}

	references_sourcesQ.SetUpdatedBy(int(in.RequesterId))

	references_sources, err = references_sourcesQ.Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, utils.Rollback(tx, errs.ProductsNotFound(int(in.Id)))
		}
		if ent.IsConstraintError(err) {
			return nil, utils.Rollback(tx, errs.InvalidForeignKey(err))
		}
		return nil, errs.SavingError("references_sources", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &references_sources_proto.UpdateResponse{
		RequesterId: uint32(references_sources.CreatedBy),
		Name:        string(*references_sources.Name),
	}, nil
}
