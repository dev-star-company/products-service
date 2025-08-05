package category_controller

import (
	"context"
	"products-service/generated_protos/category_proto"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"
)

func (c *controller) Create(ctx context.Context, in *category_proto.CreateRequest) (*category_proto.CreateResponse, error) {

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	create, err := c.Db.Category.Create().
		SetName(in.Name).
		Save(ctx)

	if err != nil {
		return nil, errs.CreateError("product type", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &category_proto.CreateResponse{

		Name: string(*create.Name),
	}, nil
}
