package features_unit_values_controller

import (
	"context"
	"products-service/internal/app/ent"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/features_unit_values_proto"
)

type Controller interface {
	features_unit_values_proto.FeaturesUnitValuesServiceServer

	Create(ctx context.Context, in *features_unit_values_proto.CreateRequest) (*features_unit_values_proto.CreateResponse, error)
	Get(ctx context.Context, in *features_unit_values_proto.GetRequest) (*features_unit_values_proto.GetResponse, error)
	Update(ctx context.Context, in *features_unit_values_proto.UpdateRequest) (*features_unit_values_proto.UpdateResponse, error)
	Delete(ctx context.Context, in *features_unit_values_proto.DeleteRequest) (*features_unit_values_proto.DeleteResponse, error)
	List(ctx context.Context, in *features_unit_values_proto.ListRequest) (*features_unit_values_proto.ListResponse, error)
}

type controller struct {
	features_unit_values_proto.UnimplementedFeaturesUnitValuesServiceServer

	Db *ent.Client
}

func New(Db *ent.Client) Controller {
	return &controller{
		Db: Db,
	}
}
