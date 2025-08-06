package product_prices_controller

import (
	"context"
	"products-service/internal/app/ent"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/product_prices_proto"
)

type Controller interface {
	product_prices_proto.ProductPricesServiceServer

	Create(ctx context.Context, in *product_prices_proto.CreateRequest) (*product_prices_proto.CreateResponse, error)
	Get(ctx context.Context, in *product_prices_proto.GetRequest) (*product_prices_proto.GetResponse, error)
	Update(ctx context.Context, in *product_prices_proto.UpdateRequest) (*product_prices_proto.UpdateResponse, error)
	Delete(ctx context.Context, in *product_prices_proto.DeleteRequest) (*product_prices_proto.DeleteResponse, error)
	List(ctx context.Context, in *product_prices_proto.ListRequest) (*product_prices_proto.ListResponse, error)
}

type controller struct {
	product_prices_proto.UnimplementedProductPricesServiceServer

	Db *ent.Client
}

func New(Db *ent.Client) Controller {
	return &controller{
		Db: Db,
	}
}
