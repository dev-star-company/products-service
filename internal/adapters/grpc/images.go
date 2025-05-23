package grpc_convertions

import (
	"products-service/generated_protos/images_proto"
	"products-service/internal/app/ent"
)

func ImagesToProto(images *ent.Images) *images_proto.Images {
	if images == nil {
		return nil
	}

	cur := &images_proto.Images{
		Id:                uint32(images.ID),
		ImageFolderPathId: uint32(*images.ImageFolderPathID),
		Content:           *images.Content,
		Path:              *images.Path,
		CreatedBy:         uint32(images.CreatedBy),
		UpdatedBy:         uint32(images.UpdatedBy),
		CreatedAt:         images.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:         images.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	if images.DeletedAt != nil {
		x := images.DeletedAt.Format("2006-01-02 15:04:05")
		cur.DeletedAt = &x
	}

	if images.DeletedBy != nil {
		x := uint32(*images.DeletedBy)
		cur.DeletedBy = &x
	}

	return cur
}
