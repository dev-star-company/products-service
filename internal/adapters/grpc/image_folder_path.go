package grpc_convertions

import (
	"products-service/generated_protos/image_folder_path_proto"
	"products-service/internal/app/ent"
)

func ImageFolderPathToProto(image_folder_path *ent.ImageFolderPath) *image_folder_path_proto.ImageFolderPath {
	if image_folder_path == nil {
		return nil
	}

	cur := &image_folder_path_proto.ImageFolderPath{
		Id:                  uint32(image_folder_path.ID),
		ImageFolderSourceId: uint32(*image_folder_path.ImageFolderSourceID),
		CreatedBy:           uint32(image_folder_path.CreatedBy),
		UpdatedBy:           uint32(image_folder_path.UpdatedBy),
		CreatedAt:           image_folder_path.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:           image_folder_path.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	if image_folder_path.DeletedAt != nil {
		x := image_folder_path.DeletedAt.Format("2006-01-02 15:04:05")
		cur.DeletedAt = &x
	}

	if image_folder_path.DeletedBy != nil {
		x := uint32(*image_folder_path.DeletedBy)
		cur.DeletedBy = &x
	}

	return cur
}
