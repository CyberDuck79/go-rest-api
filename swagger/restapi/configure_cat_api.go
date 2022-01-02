// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"
	"os"
	"rest_api/clients/database"
	"rest_api/modules/cats"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"rest_api/swagger/restapi/operations"
)

//go:generate swagger generate server --target ../../swagger --name CatAPI --spec ../swagger.yml --principal interface{}

func configureFlags(api *operations.CatAPIAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.CatAPIAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	api.UseSwaggerUI()

	// TODO config with flags / files / env

	// DB client config
	dbClient := database.Init()
	if dbClient == nil {
		os.Exit(1)
	}

	catsModule := cats.Module{}
	catsModule.Init(api, dbClient)

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {
		// add all clients closing here
		_ = dbClient.Close()
	}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
