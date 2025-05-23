package product_has_info_controller

import (
	"context"
	"products-service/generated_protos/product_has_info_proto"
	"products-service/internal/app/ent"
)

type Controller interface {
	product_has_info_proto.ProductHasInfoServiceServer

	Create(ctx context.Context, in *product_has_info_proto.CreateRequest) (*product_has_info_proto.CreateResponse, error)
	Get(ctx context.Context, in *product_has_info_proto.GetRequest) (*product_has_info_proto.GetResponse, error)
	Update(ctx context.Context, in *product_has_info_proto.UpdateRequest) (*product_has_info_proto.UpdateResponse, error)
	Delete(ctx context.Context, in *product_has_info_proto.DeleteRequest) (*product_has_info_proto.DeleteResponse, error)
	List(ctx context.Context, in *product_has_info_proto.ListRequest) (*product_has_info_proto.ListResponse, error)
}

type controller struct {
	product_has_info_proto.UnimplementedProductHasInfoServiceServer

	Db *ent.Client
}

func New(Db *ent.Client) Controller {
	return &controller{
		Db: Db,
	}
}
