package assetserver

import (
	"fmt"
	"log/slog"
	"net/http"
)

// Options defines the configuration of the AssetServer.
type Options struct {
	// Handler which serves all the content to the WebView.
	Handler http.Handler

	// Middleware is a HTTP Middleware which allows to hook into the AssetServer request chain. It allows to skip the default
	// request handler dynamically, e.g. implement specialized Routing etc.
	// The Middleware is called to build a new `http.Handler` used by the AssetSever and it also receives the default
	// handler used by the AssetServer as an argument.
	//
	// This middleware injects itself before any of Wails internal middlewares.
	//
	// If not defined, the default AssetServer request chain is executed.
	//
	// Multiple Middlewares can be chained together with:
	//   ChainMiddleware(middleware ...Middleware) Middleware
	Middleware Middleware

	// Logger is the logger used by the AssetServer. If not defined, no logging will be done.
	Logger *slog.Logger

	// RuntimeHandler is the handler used for the runtime calls.
	RuntimeHandler http.Handler

	// GetCapabilities returns the capabilities of the runtime
	GetCapabilities func() []byte

	// GetFlags returns the application flags
	GetFlags func() []byte
}

// Validate the options
func (o Options) Validate() error {
	if o.Handler == nil && o.Middleware == nil {
		return fmt.Errorf("AssetServer options invalid: either Handler or Middleware must be set")
	}

	return nil
}
