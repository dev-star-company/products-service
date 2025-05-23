package image_folder_path_controller

import (
	"context"
	"errors"
	"products-service/generated_protos/image_folder_path_proto"
	grpc_convertions "products-service/internal/adapters/grpc"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/imagefolderpath"
	"products-service/internal/app/ent/schema"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"
)

func (c *controller) List(ctx context.Context, in *image_folder_path_proto.ListRequest) (*image_folder_path_proto.ListResponse, error) {
	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	if in.IncludeDeleted != nil && *in.IncludeDeleted {
		ctx = schema.SkipSoftDelete(ctx)
	}

	query := tx.ImageFolderPath.Query()

	if in.ImageFolderSourceId != nil {
		query = query.Where(imagefolderpath.ImageFolderSourceID(int(*in.ImageFolderSourceId)))
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errs.ListingError("querying image_folder_path", err)
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
				query = query.Order(ent.Asc(imagefolderpath.FieldID))
			case "DESC":
				query = query.Order(ent.Desc(imagefolderpath.FieldID))
			default:
				return nil, errs.InvalidOrderByValue(errors.New(*in.Orderby.Id))
			}
		}
	}

	image_folder_path, err := query.All(ctx)
	if err != nil {
		return nil, errs.ListingError("querying image_folder_path", err)
	}

	responseImageFolderPath := make([]*image_folder_path_proto.ImageFolderPath, len(image_folder_path))
	for i, acc := range image_folder_path {
		responseImageFolderPath[i] = grpc_convertions.ImageFolderPathToProto(acc)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &image_folder_path_proto.ListResponse{
		Rows:  responseImageFolderPath,
		Count: uint32(count),
	}, nil
}
