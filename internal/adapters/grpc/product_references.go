package grpc_convertions

import (
	"products-service/generated_protos/product_references_proto"
	"products-service/internal/app/ent"
)

func ProductReferencesToProto(product_references *ent.ProductReferences) *product_references_proto.ProductReferences {
	if product_references == nil {
		return nil
	}

	cur := &product_references_proto.ProductReferences{
		Id:                uint32(product_references.ID),
		ReferenceSourceId: uint32(*product_references.ReferenceSourceID),
		Value:             *product_references.Value,
		CreatedAt:         product_references.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	if product_references.DeletedAt != nil {
		x := product_references.DeletedAt.Format("2006-01-02 15:04:05")
		cur.DeletedAt = &x
	}

	return cur
}
