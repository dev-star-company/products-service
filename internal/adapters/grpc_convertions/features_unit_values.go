package grpc_convertions

import (
	"products-service/internal/app/ent"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/features_unit_values_proto"
)

func FeaturesUnitValuesToProto(features_unit_values *ent.FeaturesUnitValues) *features_unit_values_proto.FeaturesUnitValues {
	if features_unit_values == nil {
		return nil
	}

	cur := &features_unit_values_proto.FeaturesUnitValues{
		Id:        uint32(features_unit_values.ID),
		Name:      *features_unit_values.Name,
		CreatedAt: features_unit_values.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	if features_unit_values.Decimals != nil {
		x := float32(*features_unit_values.Decimals)
		cur.Decimals = &x
	}

	if features_unit_values.DeletedAt != nil {
		x := features_unit_values.DeletedAt.Format("2006-01-02 15:04:05")
		cur.DeletedAt = &x
	}

	return cur
}
