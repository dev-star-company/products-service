package grpc_convertions

import (
	"products-service/generated_protos/features_unit_values_proto"
	"products-service/internal/app/ent"
)

func FeaturesUnitValuesToProto(features_unit_values *ent.FeaturesUnitValues) *features_unit_values_proto.FeaturesUnitValues {
	if features_unit_values == nil {
		return nil
	}

	cur := &features_unit_values_proto.FeaturesUnitValues{
		Id:        uint32(features_unit_values.ID),
		Name:      *features_unit_values.Name,
		CreatedBy: uint32(features_unit_values.CreatedBy),
		UpdatedBy: uint32(features_unit_values.UpdatedBy),
		CreatedAt: features_unit_values.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: features_unit_values.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	if features_unit_values.Decimals != nil {
		x := float32(*features_unit_values.Decimals)
		cur.Decimals = &x
	}

	if features_unit_values.DeletedAt != nil {
		x := features_unit_values.DeletedAt.Format("2006-01-02 15:04:05")
		cur.DeletedAt = &x
	}

	if features_unit_values.DeletedBy != nil {
		x := uint32(*features_unit_values.DeletedBy)
		cur.DeletedBy = &x
	}

	return cur
}
