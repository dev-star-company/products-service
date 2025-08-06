package info_types_controller

import (
	"context"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/infotypes"
	"products-service/internal/pkg/errs"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/info_types_proto"
)

func (c *controller) Get(ctx context.Context, in *info_types_proto.GetRequest) (*info_types_proto.GetResponse, error) {
	info_types, err := c.Db.InfoTypes.
		Query().
		Where(infotypes.ID(int(in.Id))).
		Only(ctx)

	if ent.IsNotFound(err) {
		return nil, errs.InfoTypesNotFound(int(in.Id))
	}

	return &info_types_proto.GetResponse{

		Name: *info_types.Name,
	}, nil
}
