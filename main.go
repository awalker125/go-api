package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/awalker125/go-api/handlers/cars"
	"github.com/awalker125/go-api/handlers/home"
	"github.com/awalker125/go-api/helpers"
	"github.com/awalker125/go-api/middleware"
	"github.com/awalker125/go-api/store"

	_ "github.com/awalker125/go-api/docs" // docs is generated by Swag CLI, you have to import it.

	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

//	@title			Example Cars API
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/v1

//	@securityDefinitions.basic	BasicAuth

//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/

func main() {

	mux := http.NewServeMux()

	publicUrl := helpers.GetEnv("GO_API_PUBLIC_URL", "http://localhost:8080/v1")

	uri, err := url.Parse(publicUrl)
	if err != nil {
		panic(err)
	}

	mux.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),                        // The url pointing to API definition
		httpSwagger.DefaultModelsExpandDepth(httpSwagger.HideModel), // Models will not be expanded
		httpSwagger.BeforeScript(helpers.SwaggerBeforeScriptJs),
		httpSwagger.Plugins(helpers.UrlMutatorPlugin()),
		httpSwagger.UIConfig(helpers.SwaggerUIConfig(uri)),
	))

	mux.HandleFunc("/", middleware.Chain(home.HomeHandler, middleware.Method("GET"), middleware.Logging()))

	s := store.NewMemStore()

	h := cars.NewCarsHandler(s)

	mux.HandleFunc("/v1/cars", middleware.Chain(h.ServeHTTP, middleware.Logging()))
	mux.HandleFunc("/v1/cars/", middleware.Chain(h.ServeHTTP, middleware.Logging()))

	log.Println("Starting server on 8080...")
	http.ListenAndServe(":8080", mux)
}
