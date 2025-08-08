package promotions_controller

import (
	"context"
	"fmt"
	"products-service/internal/adapters/grpc_convertions"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"
	"time"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/promotions_proto"
)

func (c *controller) Create(ctx context.Context, in *promotions_proto.CreateRequest) (*promotions_proto.CreateResponse, error) {

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	layout := time.RFC3339
	startTime, err := time.Parse(layout, in.StartingDatetime)
	if err != nil {
		return nil, fmt.Errorf("invalid starting_datetime: %w", err)
	}

	endTime, err := time.Parse(layout, in.EndingDatetime)
	if err != nil {
		return nil, fmt.Errorf("invalid ending_datetime: %w", err)
	}

	if endTime.Before(startTime) {
		return nil, fmt.Errorf("ending_datetime cannot be before starting_datetime")
	}

	create, err := c.Db.Promotions.Create().
		SetName(in.Name).
		SetStartingDatetime(startTime).
		SetEndingDatetime(endTime).
		Save(ctx)

	if err != nil {
		return nil, errs.CreateError("product type", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &promotions_proto.CreateResponse{
		Promotions: grpc_convertions.PromotionsToProto(create),
	}, nil
}
