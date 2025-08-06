package product_info_controller

import (
	"context"
	"products-service/internal/app/ent"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/product_info_proto"
)

type Controller interface {
	product_info_proto.ProductInfoServiceServer

	Create(ctx context.Context, in *product_info_proto.CreateRequest) (*product_info_proto.CreateResponse, error)
	Get(ctx context.Context, in *product_info_proto.GetRequest) (*product_info_proto.GetResponse, error)
	Update(ctx context.Context, in *product_info_proto.UpdateRequest) (*product_info_proto.UpdateResponse, error)
	Delete(ctx context.Context, in *product_info_proto.DeleteRequest) (*product_info_proto.DeleteResponse, error)
	List(ctx context.Context, in *product_info_proto.ListRequest) (*product_info_proto.ListResponse, error)
}

type controller struct {
	product_info_proto.UnimplementedProductInfoServiceServer

	Db *ent.Client
}

func New(Db *ent.Client) Controller {
	return &controller{
		Db: Db,
	}
}
