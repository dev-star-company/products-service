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
		CreatedAt:  product_has_image.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	if product_has_image.DeletedAt != nil {
		x := product_has_image.DeletedAt.Format("2006-01-02 15:04:05")
		cur.DeletedAt = &x
	}

	return cur
}
