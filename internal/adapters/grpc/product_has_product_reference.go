package grpc_convertions

import (
	"products-service/generated_protos/product_has_product_reference_proto"
	"products-service/internal/app/ent"
)

func ProductHasProductReferenceToProto(product_has_product_reference *ent.ProductHasProductReference) *product_has_product_reference_proto.ProductHasProductReference {
	if product_has_product_reference == nil {
		return nil
	}

	cur := &product_has_product_reference_proto.ProductHasProductReference{
		Id:                 uint32(product_has_product_reference.ID),
		ProductsId:         uint32(*product_has_product_reference.ProductID),
		ProductReferenceId: uint32(*product_has_product_reference.ProductReferenceID),
		CreatedAt:          product_has_product_reference.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	if product_has_product_reference.DeletedAt != nil {
		x := product_has_product_reference.DeletedAt.Format("2006-01-02 15:04:05")
		cur.DeletedAt = &x
	}
	return cur
}
