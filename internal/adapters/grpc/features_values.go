package grpc_convertions

import (
	"products-service/generated_protos/features_values_proto"
	"products-service/internal/app/ent"
)

func FeaturesValuesToProto(features_values *ent.FeaturesValues) *features_values_proto.FeaturesValues {
	if features_values == nil {
		return nil
	}

	cur := &features_values_proto.FeaturesValues{
		Id:                  uint32(features_values.ID),
		FeatureId:           uint32(*features_values.FeatureID),
		FeatureUnitValuesId: uint32(*features_values.FeatureUnitValuesID),
		Value:               *features_values.Value,
		CreatedAt:           features_values.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	if features_values.FeatureValuesID != nil {
		x := uint32(*features_values.FeatureValuesID)
		cur.FeatureValuesId = &x
	}

	if features_values.DeletedAt != nil {
		x := features_values.DeletedAt.Format("2006-01-02 15:04:05")
		cur.DeletedAt = &x
	}

	return cur
}
