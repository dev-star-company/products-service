package grpc_convertions

import (
	"products-service/generated_protos/product_has_image_proto"
	"products-service/internal/app/ent"
)

func ProductHasImageToProto(product_has_image *ent.ProductHasImage) *product_has_image_proto.ProductHasImage {
	if product_has_image == nil {
		return nil
	}

	cur := &product_has_image_proto.ProductHasImage{
		Id:         uint32(product_has_image.ID),
		ImagesId:   uint32(*product_has_image.ImageID),
		ProductsId: uint32(*product_has_image.ProductID),
		Priority:   uint32(product_has_image.Priority),
		CreatedBy:  uint32(product_has_image.CreatedBy),
		UpdatedBy:  uint32(product_has_image.UpdatedBy),
		CreatedAt:  product_has_image.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:  product_has_image.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	if product_has_image.DeletedAt != nil {
		x := product_has_image.DeletedAt.Format("2006-01-02 15:04:05")
		cur.DeletedAt = &x
	}

	if product_has_image.DeletedBy != nil {
		x := uint32(*product_has_image.DeletedBy)
		cur.DeletedBy = &x
	}

	return cur
}
