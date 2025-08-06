package references_sources_controller

import (
	"context"
	"errors"
	grpc_convertions "products-service/internal/adapters/grpc"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/referencesources"
	"products-service/internal/app/ent/schema"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/references_sources_proto"
)

func (c *controller) List(ctx context.Context, in *references_sources_proto.ListRequest) (*references_sources_proto.ListResponse, error) {
	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	if in.IncludeDeleted != nil && *in.IncludeDeleted {
		ctx = schema.SkipSoftDelete(ctx)
	}

	query := tx.ReferenceSources.Query()

	if in.Name != nil {
		query = query.Where(referencesources.Name(string(*in.Name)))
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errs.ListingError("querying references_sources", err)
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
				query = query.Order(ent.Asc(referencesources.FieldID))
			case "DESC":
				query = query.Order(ent.Desc(referencesources.FieldID))
			default:
				return nil, errs.InvalidOrderByValue(errors.New(*in.Orderby.Id))
			}
		}
	}

	references_sources, err := query.All(ctx)
	if err != nil {
		return nil, errs.ListingError("querying references_sources", err)
	}

	responseReferenceSources := make([]*references_sources_proto.ReferencesSources, len(references_sources))
	for i, acc := range references_sources {
		responseReferenceSources[i] = grpc_convertions.ReferencesSourcesToProto(acc)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &references_sources_proto.ListResponse{
		Rows:  responseReferenceSources,
		Count: uint32(count),
	}, nil
}
