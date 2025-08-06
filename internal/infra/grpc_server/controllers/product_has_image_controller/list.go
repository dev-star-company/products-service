package product_has_image_controller

import (
	"context"
	"errors"
	grpc_convertions "products-service/internal/adapters/grpc"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/producthasimage"
	"products-service/internal/app/ent/schema"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/product_has_image_proto"
)

func (c *controller) List(ctx context.Context, in *product_has_image_proto.ListRequest) (*product_has_image_proto.ListResponse, error) {
	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	if in.IncludeDeleted != nil && *in.IncludeDeleted {
		ctx = schema.SkipSoftDelete(ctx)
	}

	query := tx.ProductHasImage.Query()

	if in.ProductsId != nil {
		query = query.Where(producthasimage.ProductID(int(*in.ProductsId)))
	}

	if in.ImagesId != nil {
		query = query.Where(producthasimage.ImageID(int(*in.ImagesId)))
	}

	if in.Priority != nil {
		query = query.Where(producthasimage.Priority(int(*in.Priority)))
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errs.ListingError("querying product_has_image", err)
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
				query = query.Order(ent.Asc(producthasimage.FieldID))
			case "DESC":
				query = query.Order(ent.Desc(producthasimage.FieldID))
			default:
				return nil, errs.InvalidOrderByValue(errors.New(*in.Orderby.Id))
			}
		}
	}

	product_has_image, err := query.All(ctx)
	if err != nil {
		return nil, errs.ListingError("querying product_has_image", err)
	}

	responseProductHasImage := make([]*product_has_image_proto.ProductHasImage, len(product_has_image))
	for i, acc := range product_has_image {
		responseProductHasImage[i] = grpc_convertions.ProductHasImageToProto(acc)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &product_has_image_proto.ListResponse{
		Rows:  responseProductHasImage,
		Count: uint32(count),
	}, nil
}
