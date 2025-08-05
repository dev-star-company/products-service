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
		CreatedAt:      features.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	if features.DeletedAt != nil {
		x := features.DeletedAt.Format("2006-01-02 15:04:05")
		cur.DeletedAt = &x
	}

	return cur
}
