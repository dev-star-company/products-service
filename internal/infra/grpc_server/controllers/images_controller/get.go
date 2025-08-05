package images_controller

import (
	"context"
	"products-service/generated_protos/images_proto"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/images"
	"products-service/internal/pkg/errs"
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

		Content:           string(*images.Content),
		Path:              string(*images.Path),
	}, nil
}
