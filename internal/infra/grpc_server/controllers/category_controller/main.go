package category_controller

import (
	"context"
	"products-service/generated_protos/category_proto"
	"products-service/internal/app/ent"
)

type Controller interface {
	category_proto.CategoryServiceServer

	Create(ctx context.Context, in *category_proto.CreateRequest) (*category_proto.CreateResponse, error)
	Get(ctx context.Context, in *category_proto.GetRequest) (*category_proto.GetResponse, error)
	Update(ctx context.Context, in *category_proto.UpdateRequest) (*category_proto.UpdateResponse, error)
	Delete(ctx context.Context, in *category_proto.DeleteRequest) (*category_proto.DeleteResponse, error)
	List(ctx context.Context, in *category_proto.ListRequest) (*category_proto.ListResponse, error)
}

type controller struct {
	category_proto.UnimplementedCategoryServiceServer

	Db *ent.Client
}

func New(Db *ent.Client) Controller {
	return &controller{
		Db: Db,
	}
}
