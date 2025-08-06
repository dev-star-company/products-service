package tools_controller

import (
	"context"
	"products-service/internal/app/ent"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/tools_proto"
)

type Controller interface {
	tools_proto.ToolsServiceServer

	Create(ctx context.Context, in *tools_proto.CreateRequest) (*tools_proto.CreateResponse, error)
	Get(ctx context.Context, in *tools_proto.GetRequest) (*tools_proto.GetResponse, error)
	Update(ctx context.Context, in *tools_proto.UpdateRequest) (*tools_proto.UpdateResponse, error)
	Delete(ctx context.Context, in *tools_proto.DeleteRequest) (*tools_proto.DeleteResponse, error)
	List(ctx context.Context, in *tools_proto.ListRequest) (*tools_proto.ListResponse, error)
}

type controller struct {
	tools_proto.UnimplementedToolsServiceServer

	Db *ent.Client
}

func New(Db *ent.Client) Controller {
	return &controller{
		Db: Db,
	}
}
