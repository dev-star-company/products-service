package promotions_controller

import (
	"context"
	"products-service/generated_protos/promotions_proto"
	"products-service/internal/app/ent"
)

type Controller interface {
	promotions_proto.PromotionsServiceServer

	Create(ctx context.Context, in *promotions_proto.CreateRequest) (*promotions_proto.CreateResponse, error)
	Get(ctx context.Context, in *promotions_proto.GetRequest) (*promotions_proto.GetResponse, error)
	Update(ctx context.Context, in *promotions_proto.UpdateRequest) (*promotions_proto.UpdateResponse, error)
	Delete(ctx context.Context, in *promotions_proto.DeleteRequest) (*promotions_proto.DeleteResponse, error)
	List(ctx context.Context, in *promotions_proto.ListRequest) (*promotions_proto.ListResponse, error)
}

type controller struct {
	promotions_proto.UnimplementedPromotionsServiceServer

	Db *ent.Client
}

func New(Db *ent.Client) Controller {
	return &controller{
		Db: Db,
	}
}
