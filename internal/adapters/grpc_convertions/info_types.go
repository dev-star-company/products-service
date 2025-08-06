package grpc_convertions

import (
	"products-service/internal/app/ent"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/info_types_proto"
)

func InfoTypesToProto(info_types *ent.InfoTypes) *info_types_proto.InfoTypes {
	if info_types == nil {
		return nil
	}

	cur := &info_types_proto.InfoTypes{
		Id:        uint32(info_types.ID),
		Name:      *info_types.Name,
		CreatedAt: info_types.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	if info_types.DeletedAt != nil {
		x := info_types.DeletedAt.Format("2006-01-02 15:04:05")
		cur.DeletedAt = &x
	}

	return cur
}
