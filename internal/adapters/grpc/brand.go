package grpc_convertions

import (
	"products-service/generated_protos/brand_proto"

	"products-service/internal/app/ent"
)

func BrandToProto(brand *ent.Brand) *brand_proto.Brand {
	if brand == nil {
		return nil
	}

	cur := &brand_proto.Brand{
		Id:        uint32(brand.ID),
		Name:      *brand.Name,
		CreatedBy: uint32(brand.CreatedBy),
		UpdatedBy: uint32(brand.UpdatedBy),
		CreatedAt: brand.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: brand.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	if brand.DeletedAt != nil {
		x := brand.DeletedAt.Format("2006-01-02 15:04:05")
		cur.DeletedAt = &x
	}

	if brand.DeletedBy != nil {
		x := uint32(*brand.DeletedBy)
		cur.DeletedBy = &x
	}

	return cur
}
