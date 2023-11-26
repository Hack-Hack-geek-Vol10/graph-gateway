package graph

import "github.com/Hack-Hack-geek-Vol10/graph-gateway/src/services"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	user services.UserService
	project services.ProjectService
}
