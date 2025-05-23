package tool_has_product_controller

import (
	"context"
	"products-service/generated_protos/tool_has_product_proto"
	"products-service/internal/app/ent"
)

type Controller interface {
	tool_has_product_proto.ToolHasProductServiceServer

	Create(ctx context.Context, in *tool_has_product_proto.CreateRequest) (*tool_has_product_proto.CreateResponse, error)
	Get(ctx context.Context, in *tool_has_product_proto.GetRequest) (*tool_has_product_proto.GetResponse, error)
	Update(ctx context.Context, in *tool_has_product_proto.UpdateRequest) (*tool_has_product_proto.UpdateResponse, error)
	Delete(ctx context.Context, in *tool_has_product_proto.DeleteRequest) (*tool_has_product_proto.DeleteResponse, error)
	List(ctx context.Context, in *tool_has_product_proto.ListRequest) (*tool_has_product_proto.ListResponse, error)
}

type controller struct {
	tool_has_product_proto.UnimplementedToolHasProductServiceServer

	Db *ent.Client
}

func New(Db *ent.Client) Controller {
	return &controller{
		Db: Db,
	}
}
