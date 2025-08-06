package product_has_info_controller

import (
	"context"
	"products-service/internal/adapters/grpc_convertions"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/producthasinfo"
	"products-service/internal/pkg/errs"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/product_has_info_proto"
)

func (c *controller) Get(ctx context.Context, in *product_has_info_proto.GetRequest) (*product_has_info_proto.GetResponse, error) {
	product_has_info, err := c.Db.ProductHasInfo.
		Query().
		Where(producthasinfo.ID(int(in.Id))).
		WithProductInfo().
		WithProducts().
		Only(ctx)

	if ent.IsNotFound(err) {
		return nil, errs.ProductHasInfoNotFound(int(in.Id))
	}

	return &product_has_info_proto.GetResponse{
		Producthasinfo: grpc_convertions.ProductHasInfoToProto(product_has_info),
	}, nil
}
