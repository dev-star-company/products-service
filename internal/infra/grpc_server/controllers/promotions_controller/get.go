package promotions_controller

import (
	"context"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/promotions"
	"products-service/internal/pkg/errs"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/promotions_proto"
)

func (c *controller) Get(ctx context.Context, in *promotions_proto.GetRequest) (*promotions_proto.GetResponse, error) {
	promotions, err := c.Db.Promotions.
		Query().
		Where(promotions.ID(int(in.Id))).
		Only(ctx)

	if ent.IsNotFound(err) {
		return nil, errs.PromotionsNotFound(int(in.Id))
	}

	return &promotions_proto.GetResponse{

		Name: *promotions.Name,
	}, nil
}
