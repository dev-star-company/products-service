package grpc_convertions

import (
	"products-service/internal/app/ent"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/image_folder_source_proto"
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
		CreatedAt: image_folder_source.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	if image_folder_source.DeletedAt != nil {
		x := image_folder_source.DeletedAt.Format("2006-01-02 15:04:05")
		cur.DeletedAt = &x
	}

	return cur
}
