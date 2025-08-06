package references_sources_controller

import (
	"context"
	"products-service/internal/adapters/grpc_convertions"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/references_sources_proto"
)

func (c *controller) Create(ctx context.Context, in *references_sources_proto.CreateRequest) (*references_sources_proto.CreateResponse, error) {

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	create, err := c.Db.ReferenceSources.Create().
		SetName(in.Name).
		Save(ctx)

	if err != nil {
		return nil, errs.CreateError("reference sources type", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &references_sources_proto.CreateResponse{
		Referencessources: grpc_convertions.ReferencesSourcesToProto(create),
	}, nil
}
