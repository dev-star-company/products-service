package image_folder_source_controller

import (
	"context"
	"errors"
	"products-service/generated_protos/image_folder_source_proto"
	grpc_convertions "products-service/internal/adapters/grpc"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/imagefoldersource"
	"products-service/internal/app/ent/schema"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"
)

func (c *controller) List(ctx context.Context, in *image_folder_source_proto.ListRequest) (*image_folder_source_proto.ListResponse, error) {
	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	if in.IncludeDeleted != nil && *in.IncludeDeleted {
		ctx = schema.SkipSoftDelete(ctx)
	}

	query := tx.ImageFolderSource.Query()

	if in.Name != nil {
		query = query.Where(imagefoldersource.Name(string(*in.Name)))
	}

	if in.BaseUrl != nil {
		query = query.Where(imagefoldersource.BaseURL(string(*in.BaseUrl)))
	}

	if in.AcessKey != nil {
		query = query.Where(imagefoldersource.AccessKey(string(*in.AcessKey)))
	}

	if in.SecretKey != nil {
		query = query.Where(imagefoldersource.SecretKey(string(*in.SecretKey)))
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errs.ListingError("querying image_folder_source", err)
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
				query = query.Order(ent.Asc(imagefoldersource.FieldID))
			case "DESC":
				query = query.Order(ent.Desc(imagefoldersource.FieldID))
			default:
				return nil, errs.InvalidOrderByValue(errors.New(*in.Orderby.Id))
			}
		}
	}

	image_folder_source, err := query.All(ctx)
	if err != nil {
		return nil, errs.ListingError("querying image_folder_source", err)
	}

	responseImageFolderSource := make([]*image_folder_source_proto.ImageFolderSource, len(image_folder_source))
	for i, acc := range image_folder_source {
		responseImageFolderSource[i] = grpc_convertions.ImageFolderSourceToProto(acc)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &image_folder_source_proto.ListResponse{
		Rows:  responseImageFolderSource,
		Count: uint32(count),
	}, nil
}
