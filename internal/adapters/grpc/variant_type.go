package grpc_convertions

import (
	"products-service/generated_protos/variant_type_proto"

	"products-service/internal/app/ent"
)

func VariantTypeToProto(variant_type *ent.VariantType) *variant_type_proto.VariantType {
	if variant_type == nil {
		return nil
	}

	cur := &variant_type_proto.VariantType{
		Id:        uint32(variant_type.ID),
		Name:      *variant_type.Name,
		CreatedBy: uint32(variant_type.CreatedBy),
		UpdatedBy: uint32(variant_type.UpdatedBy),
		CreatedAt: variant_type.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: variant_type.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	if variant_type.DeletedAt != nil {
		x := variant_type.DeletedAt.Format("2006-01-02 15:04:05")
		cur.DeletedAt = &x
	}

	if variant_type.DeletedBy != nil {
		x := uint32(*variant_type.DeletedBy)
		cur.DeletedBy = &x
	}

	return cur
}
