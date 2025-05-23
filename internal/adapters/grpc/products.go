package grpc_convertions

import (
	"products-service/generated_protos/products_proto"
	"products-service/internal/app/ent"
)

func ProductsToProto(products *ent.Products) *products_proto.Products {
	if products == nil {
		return nil
	}

	cur := &products_proto.Products{
		Id:        uint32(products.ID),
		Name:      *products.Name,
		Stock:     uint32(products.Stock),
		CreatedBy: uint32(products.CreatedBy),
		UpdatedBy: uint32(products.UpdatedBy),
		CreatedAt: products.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: products.UpdatedAt.Format("2006-01-02 15:04:05"),
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

	if products.ImageID != nil {
		x := uint32(*products.ImageID)
		cur.ImageId = &x
	}

	if products.DeletedAt != nil {
		x := products.DeletedAt.Format("2006-01-02 15:04:05")
		cur.DeletedAt = &x
	}

	if products.DeletedBy != nil {
		x := uint32(*products.DeletedBy)
		cur.DeletedBy = &x
	}

	return cur
}
