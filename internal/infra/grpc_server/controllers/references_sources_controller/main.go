package references_sources_controller

import (
	"context"
	"products-service/internal/app/ent"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/references_sources_proto"
)

type Controller interface {
	references_sources_proto.ReferencesSourcesServiceServer

	Create(ctx context.Context, in *references_sources_proto.CreateRequest) (*references_sources_proto.CreateResponse, error)
	Get(ctx context.Context, in *references_sources_proto.GetRequest) (*references_sources_proto.GetResponse, error)
	Update(ctx context.Context, in *references_sources_proto.UpdateRequest) (*references_sources_proto.UpdateResponse, error)
	Delete(ctx context.Context, in *references_sources_proto.DeleteRequest) (*references_sources_proto.DeleteResponse, error)
	List(ctx context.Context, in *references_sources_proto.ListRequest) (*references_sources_proto.ListResponse, error)
}

type controller struct {
	references_sources_proto.UnimplementedReferencesSourcesServiceServer

	Db *ent.Client
}

func New(Db *ent.Client) Controller {
	return &controller{
		Db: Db,
	}
}
