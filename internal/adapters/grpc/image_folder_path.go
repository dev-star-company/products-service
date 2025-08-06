package grpc_convertions

import (
	"products-service/internal/app/ent"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/image_folder_path_proto"
)

func ImageFolderPathToProto(image_folder_path *ent.ImageFolderPath) *image_folder_path_proto.ImageFolderPath {
	if image_folder_path == nil {
		return nil
	}

	cur := &image_folder_path_proto.ImageFolderPath{
		Id:                  uint32(image_folder_path.ID),
		ImageFolderSourceId: uint32(*image_folder_path.ImageFolderSourceID),
		CreatedAt:           image_folder_path.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	if image_folder_path.DeletedAt != nil {
		x := image_folder_path.DeletedAt.Format("2006-01-02 15:04:05")
		cur.DeletedAt = &x
	}

	return cur
}
