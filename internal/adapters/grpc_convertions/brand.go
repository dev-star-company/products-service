package grpc_convertions

import (
	"github.com/dev-star-company/protos-go/products_service/generated_protos/brand_proto"

	"products-service/internal/app/ent"
)

func BrandToProto(brand *ent.Brand) *brand_proto.Brand {
	if brand == nil {
		return nil
	}

	cur := &brand_proto.Brand{
		Id:        uint32(brand.ID),
		Name:      *brand.Name,
		CreatedAt: brand.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	if brand.DeletedAt != nil {
		x := brand.DeletedAt.Format("2006-01-02 15:04:05")
		cur.DeletedAt = &x
	}

	return cur
}
