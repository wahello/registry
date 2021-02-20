// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package http

import (
	endpoint "github.com/dataleodev/registry/pkg/endpoint"
	http "github.com/go-kit/kit/transport/http"
	mux "github.com/gorilla/mux"
	http1 "net/http"
)

// NewHTTPHandler returns a handler that makes a set of endpoints available on
// predefined paths.
func NewHTTPHandler(endpoints endpoint.Endpoints, options map[string][]http.ServerOption) http1.Handler {
	m := mux.NewRouter()
	makeRegisterHandler(m, endpoints, options["Register"])
	makeLoginHandler(m, endpoints, options["Login"])
	makeViewUserHandler(m, endpoints, options["ViewUser"])
	makeListUsersHandler(m, endpoints, options["ListUsers"])
	makeUpdateUserHandler(m, endpoints, options["UpdateUser"])
	makeChangePasswordHandler(m, endpoints, options["ChangePassword"])
	makeAddNodeHandler(m, endpoints, options["AddNode"])
	makeGetNodeHandler(m, endpoints, options["GetNode"])
	makeListNodesHandler(m, endpoints, options["ListNodes"])
	makeDeleteNodeHandler(m, endpoints, options["DeleteNode"])
	makeUpdateNodeHandler(m, endpoints, options["UpdateNode"])
	makeAddRegionHandler(m, endpoints, options["AddRegion"])
	makeListRegionsHandler(m, endpoints, options["ListRegions"])
	return m
}