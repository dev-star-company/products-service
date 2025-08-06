package brand_controller

import (
	"context"
	"products-service/internal/app/ent"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/brand_proto"
)

type Controller interface {
	brand_proto.BrandServiceServer

	Create(ctx context.Context, in *brand_proto.CreateRequest) (*brand_proto.CreateResponse, error)
	Get(ctx context.Context, in *brand_proto.GetRequest) (*brand_proto.GetResponse, error)
	Update(ctx context.Context, in *brand_proto.UpdateRequest) (*brand_proto.UpdateResponse, error)
	Delete(ctx context.Context, in *brand_proto.DeleteRequest) (*brand_proto.DeleteResponse, error)
	List(ctx context.Context, in *brand_proto.ListRequest) (*brand_proto.ListResponse, error)
}

type controller struct {
	brand_proto.UnimplementedBrandServiceServer

	Db *ent.Client
}

func New(Db *ent.Client) Controller {
	return &controller{
		Db: Db,
	}
}
