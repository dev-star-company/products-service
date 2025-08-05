package grpc_convertions

import (
	"products-service/generated_protos/product_has_feature_proto"
	"products-service/internal/app/ent"
)

func ProductHasFeatureToProto(product_has_feature *ent.ProductHasFeature) *product_has_feature_proto.ProductHasFeature {
	if product_has_feature == nil {
		return nil
	}

	cur := &product_has_feature_proto.ProductHasFeature{
		Id:         uint32(product_has_feature.ID),
		FeaturesId: uint32(*product_has_feature.FeatureID),
		ProductsId: uint32(*product_has_feature.ProductID),
		CreatedAt:  product_has_feature.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	if product_has_feature.DeletedAt != nil {
		x := product_has_feature.DeletedAt.Format("2006-01-02 15:04:05")
		cur.DeletedAt = &x
	}

	return cur
}
