package grpc_convertions

import (
	"products-service/generated_protos/features_proto"
	"products-service/internal/app/ent"
)

func FeaturesToProto(features *ent.Features) *features_proto.Features {
	if features == nil {
		return nil
	}

	cur := &features_proto.Features{
		Id:             uint32(features.ID),
		FeatureValueId: uint32(*features.FeatureValueID),
		Name:           *features.Name,
		CreatedBy:      uint32(features.CreatedBy),
		UpdatedBy:      uint32(features.UpdatedBy),
		CreatedAt:      features.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:      features.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	if features.DeletedAt != nil {
		x := features.DeletedAt.Format("2006-01-02 15:04:05")
		cur.DeletedAt = &x
	}

	if features.DeletedBy != nil {
		x := uint32(*features.DeletedBy)
		cur.DeletedBy = &x
	}

	return cur
}
