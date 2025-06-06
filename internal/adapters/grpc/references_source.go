package grpc_convertions

import (
	"products-service/generated_protos/references_sources_proto"
	"products-service/internal/app/ent"
)

func ReferencesSourcesToProto(references_sources *ent.ReferenceSources) *references_sources_proto.ReferencesSources {
	if references_sources == nil {
		return nil
	}

	cur := &references_sources_proto.ReferencesSources{
		Id:        uint32(references_sources.ID),
		Name:      *references_sources.Name,
		CreatedBy: uint32(references_sources.CreatedBy),
		UpdatedBy: uint32(references_sources.UpdatedBy),
		CreatedAt: references_sources.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: references_sources.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	if references_sources.DeletedAt != nil {
		x := references_sources.DeletedAt.Format("2006-01-02 15:04:05")
		cur.DeletedAt = &x
	}

	if references_sources.DeletedBy != nil {
		x := uint32(*references_sources.DeletedBy)
		cur.DeletedBy = &x
	}

	return cur
}
