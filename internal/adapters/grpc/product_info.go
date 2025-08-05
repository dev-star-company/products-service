package grpc_convertions

import (
	"products-service/generated_protos/product_info_proto"
	"products-service/internal/app/ent"
)

func ProductInfoToProto(product_info *ent.ProductInfo) *product_info_proto.ProductInfo {
	if product_info == nil {
		return nil
	}

	cur := &product_info_proto.ProductInfo{
		Id:         uint32(product_info.ID),
		InfoTypeId: uint32(*product_info.InfoTypesID),
		Value:      *product_info.Value,
		CreatedAt:  product_info.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	if product_info.DeletedAt != nil {
		x := product_info.DeletedAt.Format("2006-01-02 15:04:05")
		cur.DeletedAt = &x
	}

	return cur
}
