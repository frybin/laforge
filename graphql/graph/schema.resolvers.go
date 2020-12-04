package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"laforge/graphql/graph/generated"
	"laforge/graphql/graph/model"
)

func (r *queryResolver) Hosts(ctx context.Context) ([]*model.Host, error) {
	testHost := model.Host{ID: "TestHost", Hostname: "test.host.me", Os: "test", LastOctect: 1, AllowMacChanges: true}
	r.hosts = append(r.hosts, &testHost)
	return r.hosts, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
