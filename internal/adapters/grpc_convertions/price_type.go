package grpc_convertions

import (
	"products-service/internal/app/ent"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/price_type_proto"
)

func PriceTypeToProto(price_type *ent.PriceType) *price_type_proto.PriceType {
	if price_type == nil {
		return nil
	}

	cur := &price_type_proto.PriceType{
		Id:        uint32(price_type.ID),
		Name:      *price_type.Name,
		CreatedAt: price_type.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	if price_type.DeletedAt != nil {
		x := price_type.DeletedAt.Format("2006-01-02 15:04:05")
		cur.DeletedAt = &x
	}

	return cur
}
