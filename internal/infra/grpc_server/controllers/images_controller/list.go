package images_controller

import (
	"context"
	"errors"
	grpc_convertions "products-service/internal/adapters/grpc"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/images"
	"products-service/internal/app/ent/schema"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/images_proto"
)

func (c *controller) List(ctx context.Context, in *images_proto.ListRequest) (*images_proto.ListResponse, error) {
	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	if in.IncludeDeleted != nil && *in.IncludeDeleted {
		ctx = schema.SkipSoftDelete(ctx)
	}

	query := tx.Images.Query()

	if in.Content != nil {
		query = query.Where(images.Content(string(*in.Content)))
	}

	if in.Path != nil {
		query = query.Where(images.Path(string(*in.Path)))
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errs.ListingError("querying images", err)
	}

	if in.Limit != nil && *in.Limit > 0 {
		query = query.Limit(int(*in.Limit))
	}

	if in.Offset != nil && *in.Offset > 0 {
		query = query.Offset(int(*in.Offset))
	}

	if in.Orderby != nil {
		if in.Orderby.Id != nil {
			switch *in.Orderby.Id {
			case "ASC":
				query = query.Order(ent.Asc(images.FieldID))
			case "DESC":
				query = query.Order(ent.Desc(images.FieldID))
			default:
				return nil, errs.InvalidOrderByValue(errors.New(*in.Orderby.Id))
			}
		}
	}

	images, err := query.All(ctx)
	if err != nil {
		return nil, errs.ListingError("querying images", err)
	}

	responseImages := make([]*images_proto.Images, len(images))
	for i, acc := range images {
		responseImages[i] = grpc_convertions.ImagesToProto(acc)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &images_proto.ListResponse{
		Rows:  responseImages,
		Count: uint32(count),
	}, nil
}
