package promotion_has_product_controller

import (
	"context"
	"products-service/generated_protos/promotion_has_product_proto"
	"products-service/internal/app/ent"
)

type Controller interface {
	promotion_has_product_proto.PromotionHasProductServiceServer

	Create(ctx context.Context, in *promotion_has_product_proto.CreateRequest) (*promotion_has_product_proto.CreateResponse, error)
	Get(ctx context.Context, in *promotion_has_product_proto.GetRequest) (*promotion_has_product_proto.GetResponse, error)
	Update(ctx context.Context, in *promotion_has_product_proto.UpdateRequest) (*promotion_has_product_proto.UpdateResponse, error)
	Delete(ctx context.Context, in *promotion_has_product_proto.DeleteRequest) (*promotion_has_product_proto.DeleteResponse, error)
	List(ctx context.Context, in *promotion_has_product_proto.ListRequest) (*promotion_has_product_proto.ListResponse, error)
}

type controller struct {
	promotion_has_product_proto.UnimplementedPromotionHasProductServiceServer

	Db *ent.Client
}

func New(Db *ent.Client) Controller {
	return &controller{
		Db: Db,
	}
}
