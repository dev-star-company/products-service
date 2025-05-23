package grpc_convertions

import (
	"products-service/generated_protos/features_values_types_proto"
	"products-service/internal/app/ent"
)

func FeaturesValuesTypesToProto(features_values_types *ent.FeaturesValuesTypes) *features_values_types_proto.FeaturesValuesTypes {
	if features_values_types == nil {
		return nil
	}

	cur := &features_values_types_proto.FeaturesValuesTypes{
		Id:        uint32(features_values_types.ID),
		Name:      *features_values_types.Name,
		CreatedBy: uint32(features_values_types.CreatedBy),
		UpdatedBy: uint32(features_values_types.UpdatedBy),
		CreatedAt: features_values_types.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: features_values_types.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	if features_values_types.DeletedAt != nil {
		x := features_values_types.DeletedAt.Format("2006-01-02 15:04:05")
		cur.DeletedAt = &x
	}

	if features_values_types.DeletedBy != nil {
		x := uint32(*features_values_types.DeletedBy)
		cur.DeletedBy = &x
	}

	return cur
}
