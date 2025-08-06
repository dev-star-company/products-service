package grpc_convertions

import (
	"github.com/dev-star-company/protos-go/products_service/generated_protos/variant_type_proto"

	"products-service/internal/app/ent"
)

func VariantTypeToProto(variant_type *ent.VariantType) *variant_type_proto.VariantType {
	if variant_type == nil {
		return nil
	}

	cur := &variant_type_proto.VariantType{
		Id:        uint32(variant_type.ID),
		Name:      *variant_type.Name,
		CreatedAt: variant_type.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	if variant_type.DeletedAt != nil {
		x := variant_type.DeletedAt.Format("2006-01-02 15:04:05")
		cur.DeletedAt = &x
	}

	return cur
}
