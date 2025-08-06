package promotion_has_product_controller

import (
	"context"
	"errors"
	grpc_convertions "products-service/internal/adapters/grpc"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/promotionhasproduct"
	"products-service/internal/app/ent/schema"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/promotion_has_product_proto"
)

func (c *controller) List(ctx context.Context, in *promotion_has_product_proto.ListRequest) (*promotion_has_product_proto.ListResponse, error) {
	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	if in.IncludeDeleted != nil && *in.IncludeDeleted {
		ctx = schema.SkipSoftDelete(ctx)
	}

	query := tx.PromotionHasProduct.Query()

	if in.PromotionId != nil {
		query = query.Where(promotionhasproduct.PromotionsID(int(*in.PromotionId)))
	}

	if in.ProductsId != nil {
		query = query.Where(promotionhasproduct.ProductsID(int(*in.ProductsId)))
	}

	count, err := query.Count(ctx)
	if err != nil {
		return nil, errs.ListingError("querying promotion_has_product", err)
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
				query = query.Order(ent.Asc(promotionhasproduct.FieldID))
			case "DESC":
				query = query.Order(ent.Desc(promotionhasproduct.FieldID))
			default:
				return nil, errs.InvalidOrderByValue(errors.New(*in.Orderby.Id))
			}
		}
	}

	promotion_has_product, err := query.All(ctx)
	if err != nil {
		return nil, errs.ListingError("querying promotion_has_product", err)
	}

	responsePromotionHasProduct := make([]*promotion_has_product_proto.PromotionHasProduct, len(promotion_has_product))
	for i, acc := range promotion_has_product {
		responsePromotionHasProduct[i] = grpc_convertions.PromotionHasProductToProto(acc)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &promotion_has_product_proto.ListResponse{
		Rows:  responsePromotionHasProduct,
		Count: uint32(count),
	}, nil
}
