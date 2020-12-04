package graph

import "laforge/graphql/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver Specify all the options that are able to be resolved here
type Resolver struct {
	hosts           []*model.Host
	povisionedHosts []*model.ProvisionedHost
}
