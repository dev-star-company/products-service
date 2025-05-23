package app

import (
	"products-service/generated_protos/brand_proto"
	"products-service/generated_protos/category_proto"
	"products-service/generated_protos/features_proto"
	"products-service/generated_protos/features_unit_values_proto"
	"products-service/generated_protos/features_values_proto"
	"products-service/generated_protos/features_values_types_proto"
	"products-service/generated_protos/image_folder_path_proto"
	"products-service/generated_protos/image_folder_source_proto"
	"products-service/generated_protos/images_proto"
	"products-service/generated_protos/info_types_proto"
	"products-service/generated_protos/price_type_proto"
	"products-service/generated_protos/product_has_feature_proto"
	"products-service/generated_protos/product_has_image_proto"
	"products-service/generated_protos/product_has_info_proto"
	"products-service/generated_protos/product_has_product_reference_proto"
	"products-service/generated_protos/product_info_proto"
	"products-service/generated_protos/product_prices_proto"
	"products-service/generated_protos/product_references_proto"
	"products-service/generated_protos/products_proto"
	"products-service/generated_protos/promotion_has_product_proto"
	"products-service/generated_protos/promotions_proto"
	"products-service/generated_protos/references_sources_proto"
	"products-service/generated_protos/tool_has_product_proto"
	"products-service/generated_protos/tools_proto"
	"products-service/generated_protos/variant_type_proto"

	"products-service/internal/app/ent"
	"products-service/internal/infra/grpc_server/controllers/brand_controller"
	"products-service/internal/infra/grpc_server/controllers/category_controller"
	"products-service/internal/infra/grpc_server/controllers/features_controller"
	"products-service/internal/infra/grpc_server/controllers/features_unit_values_controller"
	"products-service/internal/infra/grpc_server/controllers/features_values_controller"
	"products-service/internal/infra/grpc_server/controllers/features_values_types_controller"
	"products-service/internal/infra/grpc_server/controllers/image_folder_path_controller"
	"products-service/internal/infra/grpc_server/controllers/image_folder_source_controller"
	"products-service/internal/infra/grpc_server/controllers/images_controller"
	"products-service/internal/infra/grpc_server/controllers/info_types_controller"
	"products-service/internal/infra/grpc_server/controllers/price_type_controller"
	"products-service/internal/infra/grpc_server/controllers/product_has_feature_controller"
	"products-service/internal/infra/grpc_server/controllers/product_has_image_controller"
	"products-service/internal/infra/grpc_server/controllers/product_has_info_controller"
	"products-service/internal/infra/grpc_server/controllers/product_has_product_reference_controller"
	"products-service/internal/infra/grpc_server/controllers/product_info_controller"
	"products-service/internal/infra/grpc_server/controllers/product_prices_controller"
	"products-service/internal/infra/grpc_server/controllers/product_references_controller"
	"products-service/internal/infra/grpc_server/controllers/products_controller"
	"products-service/internal/infra/grpc_server/controllers/promotion_has_product_controller"
	"products-service/internal/infra/grpc_server/controllers/promotions_controller"
	"products-service/internal/infra/grpc_server/controllers/references_sources_controller"
	"products-service/internal/infra/grpc_server/controllers/tool_has_product_controller"
	"products-service/internal/infra/grpc_server/controllers/tools_controller"
	"products-service/internal/infra/grpc_server/controllers/variant_type_controller"

	"google.golang.org/grpc"
)

func RegisterControllers(grpcServer *grpc.Server, client *ent.Client) {
	brand_proto.RegisterBrandServiceServer(grpcServer, brand_controller.New(client))
	category_proto.RegisterCategoryServiceServer(grpcServer, category_controller.New(client))
	features_proto.RegisterFeaturesServiceServer(grpcServer, features_controller.New(client))
	features_unit_values_proto.RegisterFeaturesUnitValuesServiceServer(grpcServer, features_unit_values_controller.New(client))
	features_values_proto.RegisterFeaturesValuesServiceServer(grpcServer, features_values_controller.New(client))
	features_values_types_proto.RegisterFeaturesValuesTypesServiceServer(grpcServer, features_values_types_controller.New(client))
	image_folder_path_proto.RegisterImageFolderPathServiceServer(grpcServer, image_folder_path_controller.New(client))
	image_folder_source_proto.RegisterImageFolderSourceServiceServer(grpcServer, image_folder_source_controller.New(client))
	images_proto.RegisterImagesServiceServer(grpcServer, images_controller.New(client))
	info_types_proto.RegisterInfoTypesServiceServer(grpcServer, info_types_controller.New(client))
	price_type_proto.RegisterPriceTypeServiceServer(grpcServer, price_type_controller.New(client))
	product_has_feature_proto.RegisterProductHasFeatureServiceServer(grpcServer, product_has_feature_controller.New(client))
	product_has_image_proto.RegisterProductHasImageServiceServer(grpcServer, product_has_image_controller.New(client))
	product_has_info_proto.RegisterProductHasInfoServiceServer(grpcServer, product_has_info_controller.New(client))
	product_has_product_reference_proto.RegisterProductHasProductReferenceServiceServer(grpcServer, product_has_product_reference_controller.New(client))
	product_info_proto.RegisterProductInfoServiceServer(grpcServer, product_info_controller.New(client))
	product_prices_proto.RegisterProductPricesServiceServer(grpcServer, product_prices_controller.New(client))
	product_references_proto.RegisterProductReferencesServiceServer(grpcServer, product_references_controller.New(client))
	products_proto.RegisterProductsServiceServer(grpcServer, products_controller.New(client))
	promotion_has_product_proto.RegisterPromotionHasProductServiceServer(grpcServer, promotion_has_product_controller.New(client))
	promotions_proto.RegisterPromotionsServiceServer(grpcServer, promotions_controller.New(client))
	references_sources_proto.RegisterReferencesSourcesServiceServer(grpcServer, references_sources_controller.New(client))
	tool_has_product_proto.RegisterToolHasProductServiceServer(grpcServer, tool_has_product_controller.New(client))
	tools_proto.RegisterToolsServiceServer(grpcServer, tools_controller.New(client))
	variant_type_proto.RegisterVariantTypeServiceServer(grpcServer, variant_type_controller.New(client))
}
