package features_controller

import (
	"context"
	"products-service/internal/app/ent"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/features_proto"
)

type Controller interface {
	features_proto.FeaturesServiceServer

	Create(ctx context.Context, in *features_proto.CreateRequest) (*features_proto.CreateResponse, error)
	Get(ctx context.Context, in *features_proto.GetRequest) (*features_proto.GetResponse, error)
	Update(ctx context.Context, in *features_proto.UpdateRequest) (*features_proto.UpdateResponse, error)
	Delete(ctx context.Context, in *features_proto.DeleteRequest) (*features_proto.DeleteResponse, error)
	List(ctx context.Context, in *features_proto.ListRequest) (*features_proto.ListResponse, error)
}

type controller struct {
	features_proto.UnimplementedFeaturesServiceServer

	Db *ent.Client
}

func New(Db *ent.Client) Controller {
	return &controller{
		Db: Db,
	}
}
