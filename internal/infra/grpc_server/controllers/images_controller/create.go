package images_controller

import (
	"context"
	"products-service/internal/adapters/grpc_convertions"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/images_proto"
)

func (c *controller) Create(ctx context.Context, in *images_proto.CreateRequest) (*images_proto.CreateResponse, error) {

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	create, err := c.Db.Images.Create().
		SetImageFolderPathID(int(in.ImageFolderPathId)).
		SetContent(in.Content).
		SetPath(in.Path).
		Save(ctx)

	if err != nil {
		return nil, errs.CreateError("images", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &images_proto.CreateResponse{
		Images: grpc_convertions.ImagesToProto(create),
	}, nil
}
