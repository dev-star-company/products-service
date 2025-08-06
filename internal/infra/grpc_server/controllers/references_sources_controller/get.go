package references_sources_controller

import (
	"context"
	"products-service/internal/app/ent"
	"products-service/internal/app/ent/referencesources"
	"products-service/internal/pkg/errs"

	"github.com/dev-star-company/protos-go/products_service/generated_protos/references_sources_proto"
)

func (c *controller) Get(ctx context.Context, in *references_sources_proto.GetRequest) (*references_sources_proto.GetResponse, error) {
	references_sources, err := c.Db.ReferenceSources.
		Query().
		Where(referencesources.ID(int(in.Id))).
		Only(ctx)

	if ent.IsNotFound(err) {
		return nil, errs.ReferenceSourcesNotFound(int(in.Id))
	}

	return &references_sources_proto.GetResponse{
		Name: *references_sources.Name,
	}, nil
}
