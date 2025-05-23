package grpc_convertions

import (
	"products-service/generated_protos/image_folder_source_proto"
	"products-service/internal/app/ent"
)

func ImageFolderSourceToProto(image_folder_source *ent.ImageFolderSource) *image_folder_source_proto.ImageFolderSource {
	if image_folder_source == nil {
		return nil
	}

	cur := &image_folder_source_proto.ImageFolderSource{
		Id:        uint32(image_folder_source.ID),
		Name:      *image_folder_source.Name,
		BaseUrl:   *image_folder_source.BaseURL,
		AcessKey:  image_folder_source.AccessKey,
		SecretKey: image_folder_source.SecretKey,
		CreatedBy: uint32(image_folder_source.CreatedBy),
		UpdatedBy: uint32(image_folder_source.UpdatedBy),
		CreatedAt: image_folder_source.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: image_folder_source.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	if image_folder_source.DeletedAt != nil {
		x := image_folder_source.DeletedAt.Format("2006-01-02 15:04:05")
		cur.DeletedAt = &x
	}

	if image_folder_source.DeletedBy != nil {
		x := uint32(*image_folder_source.DeletedBy)
		cur.DeletedBy = &x
	}

	return cur
}
