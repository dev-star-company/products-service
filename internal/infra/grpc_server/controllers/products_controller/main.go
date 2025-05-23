package products_controller

import (
	"context"
	"products-service/generated_protos/products_proto"
	"products-service/internal/app/ent"
)

type Controller interface {
	products_proto.ProductsServiceServer

	Create(ctx context.Context, in *products_proto.CreateRequest) (*products_proto.CreateResponse, error)
	Get(ctx context.Context, in *products_proto.GetRequest) (*products_proto.GetResponse, error)
	Update(ctx context.Context, in *products_proto.UpdateRequest) (*products_proto.UpdateResponse, error)
	Delete(ctx context.Context, in *products_proto.DeleteRequest) (*products_proto.DeleteResponse, error)
	List(ctx context.Context, in *products_proto.ListRequest) (*products_proto.ListResponse, error)
}

type controller struct {
	products_proto.UnimplementedProductsServiceServer

	Db *ent.Client
}

func New(Db *ent.Client) Controller {
	return &controller{
		Db: Db,
	}
}
