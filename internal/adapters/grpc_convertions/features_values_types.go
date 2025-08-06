package grpc_convertions

import (
	"products-service/internal/app/ent"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/features_values_types_proto"
)

func FeaturesValuesTypesToProto(features_values_types *ent.FeaturesValuesTypes) *features_values_types_proto.FeaturesValuesTypes {
	if features_values_types == nil {
		return nil
	}

	cur := &features_values_types_proto.FeaturesValuesTypes{
		Id:        uint32(features_values_types.ID),
		Name:      *features_values_types.Name,
		CreatedAt: features_values_types.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	if features_values_types.DeletedAt != nil {
		x := features_values_types.DeletedAt.Format("2006-01-02 15:04:05")
		cur.DeletedAt = &x
	}

	return cur
}
