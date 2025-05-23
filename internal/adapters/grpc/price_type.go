package grpc_convertions

import (
	"products-service/generated_protos/price_type_proto"
	"products-service/internal/app/ent"
)

func PriceTypeToProto(price_type *ent.PriceType) *price_type_proto.PriceType {
	if price_type == nil {
		return nil
	}

	cur := &price_type_proto.PriceType{
		Id:        uint32(price_type.ID),
		Name:      *price_type.Name,
		CreatedBy: uint32(price_type.CreatedBy),
		UpdatedBy: uint32(price_type.UpdatedBy),
		CreatedAt: price_type.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: price_type.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	if price_type.DeletedAt != nil {
		x := price_type.DeletedAt.Format("2006-01-02 15:04:05")
		cur.DeletedAt = &x
	}

	if price_type.DeletedBy != nil {
		x := uint32(*price_type.DeletedBy)
		cur.DeletedBy = &x
	}

	return cur
}
