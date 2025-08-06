package info_types_controller

import (
	"context"
	"products-service/internal/app/ent"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/info_types_proto"
)

type Controller interface {
	info_types_proto.InfoTypesServiceServer

	Create(ctx context.Context, in *info_types_proto.CreateRequest) (*info_types_proto.CreateResponse, error)
	Get(ctx context.Context, in *info_types_proto.GetRequest) (*info_types_proto.GetResponse, error)
	Update(ctx context.Context, in *info_types_proto.UpdateRequest) (*info_types_proto.UpdateResponse, error)
	Delete(ctx context.Context, in *info_types_proto.DeleteRequest) (*info_types_proto.DeleteResponse, error)
	List(ctx context.Context, in *info_types_proto.ListRequest) (*info_types_proto.ListResponse, error)
}

type controller struct {
	info_types_proto.UnimplementedInfoTypesServiceServer

	Db *ent.Client
}

func New(Db *ent.Client) Controller {
	return &controller{
		Db: Db,
	}
}
