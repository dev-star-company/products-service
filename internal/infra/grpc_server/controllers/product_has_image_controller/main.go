package product_has_image_controller

import (
	"context"
	"products-service/generated_protos/product_has_image_proto"
	"products-service/internal/app/ent"
)

type Controller interface {
	product_has_image_proto.ProductHasImageServiceServer

	Create(ctx context.Context, in *product_has_image_proto.CreateRequest) (*product_has_image_proto.CreateResponse, error)
	Get(ctx context.Context, in *product_has_image_proto.GetRequest) (*product_has_image_proto.GetResponse, error)
	Update(ctx context.Context, in *product_has_image_proto.UpdateRequest) (*product_has_image_proto.UpdateResponse, error)
	Delete(ctx context.Context, in *product_has_image_proto.DeleteRequest) (*product_has_image_proto.DeleteResponse, error)
	List(ctx context.Context, in *product_has_image_proto.ListRequest) (*product_has_image_proto.ListResponse, error)
}

type controller struct {
	product_has_image_proto.UnimplementedProductHasImageServiceServer

	Db *ent.Client
}

func New(Db *ent.Client) Controller {
	return &controller{
		Db: Db,
	}
}
