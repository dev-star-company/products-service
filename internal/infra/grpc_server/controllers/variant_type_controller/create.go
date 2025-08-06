package variant_type_controller

import (
	"context"
	"products-service/internal/adapters/grpc_convertions"
	"products-service/internal/pkg/errs"
	"products-service/internal/pkg/utils"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/variant_type_proto"
)

func (c *controller) Create(ctx context.Context, in *variant_type_proto.CreateRequest) (*variant_type_proto.CreateResponse, error) {

	tx, err := c.Db.Tx(ctx)
	if err != nil {
		return nil, errs.StartProductsError(err)
	}

	create, err := c.Db.VariantType.Create().
		SetName(in.Name).
		Save(ctx)

	if err != nil {
		return nil, errs.CreateError("variant type", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, utils.Rollback(tx, errs.CommitProductsError(err))
	}

	return &variant_type_proto.CreateResponse{
		Varianttype: grpc_convertions.VariantTypeToProto(create),
	}, nil
}
