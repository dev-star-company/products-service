package grpc_convertions

import (
	"products-service/generated_protos/tool_has_product_proto"
	"products-service/internal/app/ent"
)

func ToolHasProductToProto(tool_has_product *ent.ToolHasProduct) *tool_has_product_proto.ToolHasProduct {
	if tool_has_product == nil {
		return nil
	}

	cur := &tool_has_product_proto.ToolHasProduct{
		Id:         uint32(tool_has_product.ID),
		ToolId:     uint32(*tool_has_product.ToolID),
		ProductsId: uint32(*tool_has_product.ProductID),
		CreatedBy:  uint32(tool_has_product.CreatedBy),
		UpdatedBy:  uint32(tool_has_product.UpdatedBy),
		CreatedAt:  tool_has_product.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:  tool_has_product.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	if tool_has_product.DeletedAt != nil {
		x := tool_has_product.DeletedAt.Format("2006-01-02 15:04:05")
		cur.DeletedAt = &x
	}

	if tool_has_product.DeletedBy != nil {
		x := uint32(*tool_has_product.DeletedBy)
		cur.DeletedBy = &x
	}

	return cur
}
