package features_controller

import (
	"context"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/features_proto"
)

func (c *controller) Create(ctx context.Context, in *features_proto.CreateRequest) (*features_proto.CreateResponse, error) {

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	create, err := c.Db.Features.Create().
		SetFeatureValueID(int(in.FeatureValueId)).
		SetName(in.Name).
		Save(ctx)

	if err != nil {
		return nil, errs.CreateError("product type", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &features_proto.CreateResponse{
		FeatureValueId: uint32(*create.FeatureValueID),
		Name:           string(*create.Name),
	}, nil
}
