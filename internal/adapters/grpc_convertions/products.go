package grpc_convertions

import (
	"products-service/internal/app/ent"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/products_proto"
)

func ProductsToProto(products *ent.Products) *products_proto.Products {
	if products == nil {
		return nil
	}

	cur := &products_proto.Products{
		Id:        uint32(products.ID),
		Name:      *products.Name,
		Stock:     uint32(products.Stock),
		CreatedAt: products.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	if products.CategoryID != nil {
		x := uint32(*products.CategoryID)
		cur.CategoryId = &x
	}

	if products.BrandID != nil {
		x := uint32(*products.BrandID)
		cur.BrandId = &x
	}

	if products.VariantTypeID != nil {
		x := uint32(*products.VariantTypeID)
		cur.VariantTypeId = &x
	}

	if products.ProductReferencesID != nil {
		x := uint32(*products.ProductReferencesID)
		cur.ProductReferencesId = &x
	}

	if products.ImagesID != nil {
		x := uint32(*products.ImagesID)
		cur.ImageId = &x
	}

	if products.DeletedAt != nil {
		x := products.DeletedAt.Format("2006-01-02 15:04:05")
		cur.DeletedAt = &x
	}

	return cur
}
