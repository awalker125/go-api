# go-api

Hello World REST API in golang

## TL;DR

If already setup

```bash
air
http://localhost:8080/swagger/
```

## Repo setup

```bash
go mod init github.com/awalker125/go-api
```

### Middleware

<https://gowebexamples.com/advanced-middleware/>

<https://medium.com/@matryer/writing-middleware-in-golang-and-how-go-makes-it-so-much-fun-4375c1246e81>

<https://www.alexedwards.net/blog/an-introduction-to-handlers-and-servemuxes-in-go>

<https://www.jetbrains.com/guide/go/tutorials/rest_api_series/stdlib/>

### Testing

<https://blog.logrocket.com/a-deep-dive-into-unit-testing-in-go/>

### Live reload

<https://github.com/air-verse/air> has lots of stars.

```bash
go install github.com/air-verse/air@latest
air init
```

Then

```
air
```

### API Tests

```bash
curl -X 'POST'   'http://localhost:8080/v1/cars/'   -H 'accept: application/json'   -H 'Content-Type: application/json'   -d '{
  "model": "string",
  "parts": [
    {
      "name": "string"
    }
  ]
}'
```

### Swagger Docs

Full docs [here](https://github.com/swaggo/swag)

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

Add general info to main.go e.g

```golang
// @title           Example Cars API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
```

Then

```bash

swag init
```

This will create the docs folder.

Then run

```bash
go get -u github.com/swaggo/http-swagger/v2
```

Then see <https://github.com/swaggo/http-swagger/pull/116/files>

Full example [here](https://github.com/swaggo/http-swagger)

## Authentication/Authoiazion

Plan is to use <https://github.com/auth0/go-jwt-middleware/blob/master/examples/http-jwks-example/main.go>

See <https://github.com/auth0/go-jwt-middleware/tree/master>

Seems reasonably well maintained
