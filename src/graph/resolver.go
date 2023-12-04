package graph

import (
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/schema-creator/graph-gateway/src/services"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	App            *newrelic.Application
	UserService    services.UserService
	ProjectService services.ProjectService
	MemberService  services.MemberService
}
