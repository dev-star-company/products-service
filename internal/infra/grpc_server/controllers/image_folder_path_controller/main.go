package image_folder_path_controller

import (
	"context"
	"products-service/generated_protos/image_folder_path_proto"
	"products-service/internal/app/ent"
)

type Controller interface {
	image_folder_path_proto.ImageFolderPathServiceServer

	Create(ctx context.Context, in *image_folder_path_proto.CreateRequest) (*image_folder_path_proto.CreateResponse, error)
	Get(ctx context.Context, in *image_folder_path_proto.GetRequest) (*image_folder_path_proto.GetResponse, error)
	Update(ctx context.Context, in *image_folder_path_proto.UpdateRequest) (*image_folder_path_proto.UpdateResponse, error)
	Delete(ctx context.Context, in *image_folder_path_proto.DeleteRequest) (*image_folder_path_proto.DeleteResponse, error)
	List(ctx context.Context, in *image_folder_path_proto.ListRequest) (*image_folder_path_proto.ListResponse, error)
}

type controller struct {
	image_folder_path_proto.UnimplementedImageFolderPathServiceServer

	Db *ent.Client
}

func New(Db *ent.Client) Controller {
	return &controller{
		Db: Db,
	}
}
