package grpc_convertions

import (
	"products-service/generated_protos/promotions_proto"
	"products-service/internal/app/ent"
)

func PromotionsToProto(promotions *ent.Promotions) *promotions_proto.Promotions {
	if promotions == nil {
		return nil
	}

	cur := &promotions_proto.Promotions{
		Id:               uint32(promotions.ID),
		Name:             *promotions.Name,
		StartingDatetime: promotions.StartingDatetime.Format("2006-01-02 15:04:05"),
		EndingDatetime:   promotions.EndingDatetime.Format("2006-01-02 15:04:05"),
		CreatedAt:        promotions.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	if promotions.DeletedAt != nil {
		x := promotions.DeletedAt.Format("2006-01-02 15:04:05")
		cur.DeletedAt = &x
	}

	return cur
}
