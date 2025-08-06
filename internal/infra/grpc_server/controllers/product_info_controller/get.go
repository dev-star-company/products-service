package product_info_controller

import (
	"context"
	"products-service/internal/adapters/grpc_convertions"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/productinfo"
	"products-service/internal/pkg/errs"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/product_info_proto"
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
		Productinfo: grpc_convertions.ProductInfoToProto(product_info),
	}, nil
}
