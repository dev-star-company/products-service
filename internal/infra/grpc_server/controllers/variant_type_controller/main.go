package variant_type_controller

import (
	"context"
	"products-service/internal/app/ent"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/variant_type_proto"
)

type Controller interface {
	variant_type_proto.VariantTypeServiceServer

	Create(ctx context.Context, in *variant_type_proto.CreateRequest) (*variant_type_proto.CreateResponse, error)
	Get(ctx context.Context, in *variant_type_proto.GetRequest) (*variant_type_proto.GetResponse, error)
	Update(ctx context.Context, in *variant_type_proto.UpdateRequest) (*variant_type_proto.UpdateResponse, error)
	Delete(ctx context.Context, in *variant_type_proto.DeleteRequest) (*variant_type_proto.DeleteResponse, error)
	List(ctx context.Context, in *variant_type_proto.ListRequest) (*variant_type_proto.ListResponse, error)
}

type controller struct {
	variant_type_proto.UnimplementedVariantTypeServiceServer

	Db *ent.Client
}

func New(Db *ent.Client) Controller {
	return &controller{
		Db: Db,
	}
}
