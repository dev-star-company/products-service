package images_controller

import (
	"context"
	"products-service/internal/adapters/grpc_convertions"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/images"
	"products-service/internal/pkg/errs"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/images_proto"
)

func (c *controller) Get(ctx context.Context, in *images_proto.GetRequest) (*images_proto.GetResponse, error) {
	images, err := c.Db.Images.
		Query().
		Where(images.ID(int(in.Id))).
		Only(ctx)

	if ent.IsNotFound(err) {
		return nil, errs.ImagesNotFound(int(in.Id))
	}

	return &images_proto.GetResponse{
		Images: grpc_convertions.ImagesToProto(images),
	}, nil
}
