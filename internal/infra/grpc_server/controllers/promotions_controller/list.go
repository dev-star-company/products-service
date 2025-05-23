package promotions_controller

import (
	"context"
	"errors"
	"products-service/generated_protos/promotions_proto"
	grpc_convertions "products-service/internal/adapters/grpc"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/promotions"
	"products-service/internal/app/ent/schema"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"
	"time"
)

func (c *controller) List(ctx context.Context, in *promotions_proto.ListRequest) (*promotions_proto.ListResponse, error) {
	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	if in.IncludeDeleted != nil && *in.IncludeDeleted {
		ctx = schema.SkipSoftDelete(ctx)
	}

	query := tx.Promotions.Query()

	if in.Name != nil {
		query = query.Where(promotions.Name(string(*in.Name)))
	}

	if in.StartingDatetime != nil && *in.StartingDatetime != "" {
		parsedTime, err := time.Parse(time.RFC3339, *in.StartingDatetime)
		if err != nil {
			return nil, errs.ListingError("querying promotions", err)
		}
		query = query.Where(promotions.StartingDatetimeEQ(parsedTime))
	}

	if in.EndingDatetime != nil && *in.EndingDatetime != "" {
		parsedTime, err := time.Parse(time.RFC3339, *in.EndingDatetime)
		if err != nil {
			return nil, errs.ListingError("querying promotions", err)
		}
		query = query.Where(promotions.EndingDatetimeEQ(parsedTime))
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errs.ListingError("querying promotions", err)
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
				query = query.Order(ent.Asc(promotions.FieldID))
			case "DESC":
				query = query.Order(ent.Desc(promotions.FieldID))
			default:
				return nil, errs.InvalidOrderByValue(errors.New(*in.Orderby.Id))
			}
		}
	}

	promotions, err := query.All(ctx)
	if err != nil {
		return nil, errs.ListingError("querying promotions", err)
	}

	responsePromotions := make([]*promotions_proto.Promotions, len(promotions))
	for i, acc := range promotions {
		responsePromotions[i] = grpc_convertions.PromotionsToProto(acc)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &promotions_proto.ListResponse{
		Rows:  responsePromotions,
		Count: uint32(count),
	}, nil
}
