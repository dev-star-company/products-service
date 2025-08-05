package product_info_controller

import (
	"context"
	"products-service/generated_protos/product_info_proto"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/productinfo"
	"products-service/internal/pkg/errs"
)

func (c *controller) Get(ctx context.Context, in *product_info_proto.GetRequest) (*product_info_proto.GetResponse, error) {
	product_info, err := c.Db.ProductInfo.
		Query().
		Where(productinfo.ID(int(in.Id))).
		Only(ctx)

	if ent.IsNotFound(err) {
		return nil, errs.ProductInfoNotFound(int(in.Id))
	}

	return &product_info_proto.GetResponse{

		InfoTypeId: uint32(*product_info.InfoTypesID),
		Value:      string(*product_info.Value),
	}, nil
}
