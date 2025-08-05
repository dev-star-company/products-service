package brand_controller

import (
	"context"
	"products-service/generated_protos/brand_proto"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/brand"
	"products-service/internal/pkg/errs"
)

func (c *controller) Get(ctx context.Context, in *brand_proto.GetRequest) (*brand_proto.GetResponse, error) {
	brand, err := c.Db.Brand.
		Query().
		Where(brand.ID(int(in.Id))).
		Only(ctx)

	if ent.IsNotFound(err) {
		return nil, errs.BrandNotFound(int(in.Id))
	}

	return &brand_proto.GetResponse{

		Name: *brand.Name,
	}, nil
}
