package product_has_feature_controller

import (
	"context"
	"products-service/generated_protos/product_has_feature_proto"
	"products-service/internal/app/ent"
)

type Controller interface {
	product_has_feature_proto.ProductHasFeatureServiceServer

	Create(ctx context.Context, in *product_has_feature_proto.CreateRequest) (*product_has_feature_proto.CreateResponse, error)
	Get(ctx context.Context, in *product_has_feature_proto.GetRequest) (*product_has_feature_proto.GetResponse, error)
	Update(ctx context.Context, in *product_has_feature_proto.UpdateRequest) (*product_has_feature_proto.UpdateResponse, error)
	Delete(ctx context.Context, in *product_has_feature_proto.DeleteRequest) (*product_has_feature_proto.DeleteResponse, error)
	List(ctx context.Context, in *product_has_feature_proto.ListRequest) (*product_has_feature_proto.ListResponse, error)
}

type controller struct {
	product_has_feature_proto.UnimplementedProductHasFeatureServiceServer

	Db *ent.Client
}

func New(Db *ent.Client) Controller {
	return &controller{
		Db: Db,
	}
}
