package grpc_convertions

import (
	"products-service/generated_protos/category_proto"
	"products-service/internal/app/ent"
)

func CategoryToProto(category *ent.Category) *category_proto.Category {
	if category == nil {
		return nil
	}

	cur := &category_proto.Category{
		Id:        uint32(category.ID),
		Name:      *category.Name,
		CreatedAt: category.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	if category.CategoryID != nil {
		x := uint32(*category.CategoryID)
		cur.CategoryId = &x
	}

	if category.DeletedAt != nil {
		x := category.DeletedAt.Format("2006-01-02 15:04:05")
		cur.DeletedAt = &x
	}

	return cur
}
