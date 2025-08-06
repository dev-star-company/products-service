package image_folder_source_controller

import (
	"context"
	"products-service/internal/app/ent"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/image_folder_source_proto"
)

type Controller interface {
	image_folder_source_proto.ImageFolderSourceServiceServer

	Create(ctx context.Context, in *image_folder_source_proto.CreateRequest) (*image_folder_source_proto.CreateResponse, error)
	Get(ctx context.Context, in *image_folder_source_proto.GetRequest) (*image_folder_source_proto.GetResponse, error)
	Update(ctx context.Context, in *image_folder_source_proto.UpdateRequest) (*image_folder_source_proto.UpdateResponse, error)
	Delete(ctx context.Context, in *image_folder_source_proto.DeleteRequest) (*image_folder_source_proto.DeleteResponse, error)
	List(ctx context.Context, in *image_folder_source_proto.ListRequest) (*image_folder_source_proto.ListResponse, error)
}

type controller struct {
	image_folder_source_proto.UnimplementedImageFolderSourceServiceServer

	Db *ent.Client
}

func New(Db *ent.Client) Controller {
	return &controller{
		Db: Db,
	}
}
