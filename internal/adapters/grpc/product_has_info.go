package grpc_convertions

import (
	"products-service/generated_protos/product_has_info_proto"
	"products-service/internal/app/ent"
)

func ProductHasInfoToProto(product_has_info *ent.ProductHasInfo) *product_has_info_proto.ProductHasInfo {
	if product_has_info == nil {
		return nil
	}

	cur := &product_has_info_proto.ProductHasInfo{
		Id:            uint32(product_has_info.ID),
		ProductsId:    uint32(*product_has_info.ProductID),
		ProductInfoId: uint32(*product_has_info.ProductInfoID),
		CreatedAt:     product_has_info.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	if product_has_info.DeletedAt != nil {
		x := product_has_info.DeletedAt.Format("2006-01-02 15:04:05")
		cur.DeletedAt = &x
	}

	return cur
}
