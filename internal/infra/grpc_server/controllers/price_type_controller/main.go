package price_type_controller

import (
	"context"
	"products-service/internal/app/ent"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/price_type_proto"
)

type Controller interface {
	price_type_proto.PriceTypeServiceServer

	Create(ctx context.Context, in *price_type_proto.CreateRequest) (*price_type_proto.CreateResponse, error)
	Get(ctx context.Context, in *price_type_proto.GetRequest) (*price_type_proto.GetResponse, error)
	Update(ctx context.Context, in *price_type_proto.UpdateRequest) (*price_type_proto.UpdateResponse, error)
	Delete(ctx context.Context, in *price_type_proto.DeleteRequest) (*price_type_proto.DeleteResponse, error)
	List(ctx context.Context, in *price_type_proto.ListRequest) (*price_type_proto.ListResponse, error)
}

type controller struct {
	price_type_proto.UnimplementedPriceTypeServiceServer

	Db *ent.Client
}

func New(Db *ent.Client) Controller {
	return &controller{
		Db: Db,
	}
}
