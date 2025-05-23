package product_has_info_controller

import (
	"context"
	"products-service/generated_protos/product_has_info_proto"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/producthasinfo"
	"products-service/internal/pkg/errs"
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
		RequesterId:   uint32(product_has_info.CreatedBy),
		ProductInfoId: uint32(*product_has_info.ProductInfoID),
		ProductsId:    uint32(*product_has_info.ProductID),
	}, nil
}
