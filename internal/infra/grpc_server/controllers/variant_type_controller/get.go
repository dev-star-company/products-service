package variant_type_controller

import (
	"context"
	"products-service/generated_protos/variant_type_proto"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/varianttype"
	"products-service/internal/pkg/errs"
)

func (c *controller) Get(ctx context.Context, in *variant_type_proto.GetRequest) (*variant_type_proto.GetResponse, error) {
	variant_type, err := c.Db.VariantType.
		Query().
		Where(varianttype.ID(int(in.Id))).
		Only(ctx)

	if ent.IsNotFound(err) {
		return nil, errs.VariantTypeNotFound(int(in.Id))
	}

	return &variant_type_proto.GetResponse{
		Name: *variant_type.Name,
	}, nil
}
