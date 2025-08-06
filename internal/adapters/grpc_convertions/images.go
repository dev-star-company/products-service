package grpc_convertions

import (
	"products-service/internal/app/ent"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/images_proto"
)

func ImagesToProto(images *ent.Images) *images_proto.Images {
	if images == nil {
		return nil
	}

	cur := &images_proto.Images{
		Id:        uint32(images.ID),
		Content:   *images.Content,
		Path:      *images.Path,
		CreatedAt: images.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	if images.DeletedAt != nil {
		x := images.DeletedAt.Format("2006-01-02 15:04:05")
		cur.DeletedAt = &x
	}

	return cur
}
