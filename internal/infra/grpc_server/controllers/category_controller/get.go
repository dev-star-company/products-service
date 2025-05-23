package category_controller

import (
	"context"
	"products-service/generated_protos/category_proto"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/category"
	"products-service/internal/pkg/errs"
)

func (c *controller) Get(ctx context.Context, in *category_proto.GetRequest) (*category_proto.GetResponse, error) {
	category, err := c.Db.Category.
		Query().
		Where(category.ID(int(in.Id))).
		Only(ctx)

	if ent.IsNotFound(err) {
		return nil, errs.CategoryNotFound(int(in.Id))
	}

	return &category_proto.GetResponse{
		RequesterId: uint32(category.CreatedBy),
		Name:        *category.Name,
	}, nil
}
