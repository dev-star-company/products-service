package grpc_convertions

import (
	"products-service/generated_protos/tools_proto"
	"products-service/internal/app/ent"
)

func ToolsToProto(tools *ent.Tools) *tools_proto.Tools {
	if tools == nil {
		return nil
	}

	cur := &tools_proto.Tools{
		Id:        uint32(tools.ID),
		Name:      *tools.Name,
		CreatedBy: uint32(tools.CreatedBy),
		UpdatedBy: uint32(tools.UpdatedBy),
		CreatedAt: tools.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: tools.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	if tools.DeletedAt != nil {
		x := tools.DeletedAt.Format("2006-01-02 15:04:05")
		cur.DeletedAt = &x
	}

	if tools.DeletedBy != nil {
		x := uint32(*tools.DeletedBy)
		cur.DeletedBy = &x
	}

	return cur
}
