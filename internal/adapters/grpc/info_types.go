package grpc_convertions

import (
	"products-service/generated_protos/info_types_proto"
	"products-service/internal/app/ent"
)

func InfoTypesToProto(info_types *ent.InfoTypes) *info_types_proto.InfoTypes {
	if info_types == nil {
		return nil
	}

	cur := &info_types_proto.InfoTypes{
		Id:        uint32(info_types.ID),
		Name:      *info_types.Name,
		CreatedBy: uint32(info_types.CreatedBy),
		UpdatedBy: uint32(info_types.UpdatedBy),
		CreatedAt: info_types.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: info_types.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	if info_types.DeletedAt != nil {
		x := info_types.DeletedAt.Format("2006-01-02 15:04:05")
		cur.DeletedAt = &x
	}

	if info_types.DeletedBy != nil {
		x := uint32(*info_types.DeletedBy)
		cur.DeletedBy = &x
	}

	return cur
}
