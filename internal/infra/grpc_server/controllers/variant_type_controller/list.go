package variant_type_controller

import (
	"context"
	"errors"
	"products-service/generated_protos/variant_type_proto"
	grpc_convertions "products-service/internal/adapters/grpc"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/schema"
	"products-service/internal/app/ent/varianttype"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"
)

func (c *controller) List(ctx context.Context, in *variant_type_proto.ListRequest) (*variant_type_proto.ListResponse, error) {
	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	if in.IncludeDeleted != nil && *in.IncludeDeleted {
		ctx = schema.SkipSoftDelete(ctx)
	}

	query := tx.VariantType.Query()

	if in.Name != nil {
		query = query.Where(varianttype.Name(string(*in.Name)))
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errs.ListingError("querying variant_type", err)
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
				query = query.Order(ent.Asc(varianttype.FieldID))
			case "DESC":
				query = query.Order(ent.Desc(varianttype.FieldID))
			default:
				return nil, errs.InvalidOrderByValue(errors.New(*in.Orderby.Id))
			}
		}
	}

	variant_type, err := query.All(ctx)
	if err != nil {
		return nil, errs.ListingError("querying variant_type", err)
	}

	responseVariantType := make([]*variant_type_proto.VariantType, len(variant_type))
	for i, acc := range variant_type {
		responseVariantType[i] = grpc_convertions.VariantTypeToProto(acc)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &variant_type_proto.ListResponse{
		Rows:  responseVariantType,
		Count: uint32(count),
	}, nil
}
