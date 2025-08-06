package features_values_controller

import (
	"context"
	"products-service/internal/app/ent"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/features_values_proto"
)

type Controller interface {
	features_values_proto.FeaturesValuesServiceServer

	Create(ctx context.Context, in *features_values_proto.CreateRequest) (*features_values_proto.CreateResponse, error)
	Get(ctx context.Context, in *features_values_proto.GetRequest) (*features_values_proto.GetResponse, error)
	Update(ctx context.Context, in *features_values_proto.UpdateRequest) (*features_values_proto.UpdateResponse, error)
	Delete(ctx context.Context, in *features_values_proto.DeleteRequest) (*features_values_proto.DeleteResponse, error)
	List(ctx context.Context, in *features_values_proto.ListRequest) (*features_values_proto.ListResponse, error)
}

type controller struct {
	features_values_proto.UnimplementedFeaturesValuesServiceServer

	Db *ent.Client
}

func New(Db *ent.Client) Controller {
	return &controller{
		Db: Db,
	}
}
