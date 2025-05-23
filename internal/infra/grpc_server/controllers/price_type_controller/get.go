package price_type_controller

import (
	"context"
	"products-service/generated_protos/price_type_proto"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/pricetype"
	"products-service/internal/pkg/errs"
)

func (c *controller) Get(ctx context.Context, in *price_type_proto.GetRequest) (*price_type_proto.GetResponse, error) {
	price_type, err := c.Db.PriceType.
		Query().
		Where(pricetype.ID(int(in.Id))).
		Only(ctx)

	if ent.IsNotFound(err) {
		return nil, errs.PriceTypeNotFound(int(in.Id))
	}

	return &price_type_proto.GetResponse{
		RequesterId: uint32(price_type.CreatedBy),
		Name:        *price_type.Name,
	}, nil
}
