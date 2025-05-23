package product_references_controller

import (
	"context"
	"products-service/generated_protos/product_references_proto"
	"products-service/internal/app/ent"
)

type Controller interface {
	product_references_proto.ProductReferencesServiceServer

	Create(ctx context.Context, in *product_references_proto.CreateRequest) (*product_references_proto.CreateResponse, error)
	Get(ctx context.Context, in *product_references_proto.GetRequest) (*product_references_proto.GetResponse, error)
	Update(ctx context.Context, in *product_references_proto.UpdateRequest) (*product_references_proto.UpdateResponse, error)
	Delete(ctx context.Context, in *product_references_proto.DeleteRequest) (*product_references_proto.DeleteResponse, error)
	List(ctx context.Context, in *product_references_proto.ListRequest) (*product_references_proto.ListResponse, error)
}

type controller struct {
	product_references_proto.UnimplementedProductReferencesServiceServer

	Db *ent.Client
}

func New(Db *ent.Client) Controller {
	return &controller{
		Db: Db,
	}
}
