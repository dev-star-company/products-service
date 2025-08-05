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
		CreatedAt: tools.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	if tools.DeletedAt != nil {
		x := tools.DeletedAt.Format("2006-01-02 15:04:05")
		cur.DeletedAt = &x
	}

	return cur
}
