package features_values_types_controller

import (
	"context"
	"products-service/generated_protos/features_values_types_proto"
	"products-service/internal/app/ent"
)

type Controller interface {
	features_values_types_proto.FeaturesValuesTypesServiceServer

	Create(ctx context.Context, in *features_values_types_proto.CreateRequest) (*features_values_types_proto.CreateResponse, error)
	Get(ctx context.Context, in *features_values_types_proto.GetRequest) (*features_values_types_proto.GetResponse, error)
	Update(ctx context.Context, in *features_values_types_proto.UpdateRequest) (*features_values_types_proto.UpdateResponse, error)
	Delete(ctx context.Context, in *features_values_types_proto.DeleteRequest) (*features_values_types_proto.DeleteResponse, error)
	List(ctx context.Context, in *features_values_types_proto.ListRequest) (*features_values_types_proto.ListResponse, error)
}

type controller struct {
	features_values_types_proto.UnimplementedFeaturesValuesTypesServiceServer

	Db *ent.Client
}

func New(Db *ent.Client) Controller {
	return &controller{
		Db: Db,
	}
}
