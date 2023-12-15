package echo

import (
	"encoding/json"
	"fmt"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/newrelic/go-agent/v3/integrations/nrecho-v4"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/schema-creator/graph-gateway/src/graph"
	"github.com/schema-creator/graph-gateway/src/infra/auth"
	"github.com/schema-creator/graph-gateway/src/internal"
)

type Rotuer struct {
	e   *echo.Echo
	app *newrelic.Application
}

func NewRouter(app *newrelic.Application) *echo.Echo {
	router := &Rotuer{
		e:   echo.New(),
		app: app,
	}

	router.setMiddleware()

	resolver := graph.NewResolver()

	router.e.GET("/query", func(c echo.Context) error {
		handler.NewDefaultServer(
			internal.NewExecutableSchema(
				internal.Config{
					Resolvers: resolver,
				},
			),
		).ServeHTTP(c.Response(), c.Request())
		return nil
	})

	router.e.POST("/query", func(c echo.Context) error {
		handler.NewDefaultServer(
			internal.NewExecutableSchema(
				internal.Config{
					Resolvers: resolver,
				},
			),
		).ServeHTTP(c.Response(), c.Request())
		return nil
	})

	router.e.GET("/playground", func(c echo.Context) error {
		playground.Handler("GraphQL", "/query").ServeHTTP(c.Response(), c.Request())
		return nil
	})

	return router.e
}

func (r *Rotuer) setMiddleware() {
	r.e.Use(
		middleware.CORSWithConfig(
			middleware.CORSConfig{
				AllowOrigins:     []string{"*", "http://localhost:3000"},
				AllowMethods:     []string{"POST", "GET", "OPTIONS"},
				AllowHeaders:     []string{auth.TokenKey, "Content-Type"},
				AllowCredentials: true,
			},
		),
		nrecho.Middleware(r.app),
		middleware.RequestLoggerWithConfig(
			middleware.RequestLoggerConfig{
				// LogStatus instructs logger to extract response status code. If handler chain returns an echo.HTTPError,
				// the status code is extracted from the echo.HTTPError returned
				LogStatus: true,
				// LogURI instructs logger to extract request URI (i.e. `/list?lang=en&page=1`)
				LogURI:      true,
				HandleError: true,

				LogLatency: true,
				// LogProtocol instructs logger to extract request protocol (i.e. `HTTP/1.1` or `HTTP/2`)
				LogProtocol: true,
				// LogRemoteIP instructs logger to extract request remote IP. See `echo.Context.RealIP()` for implementation details.
				LogRemoteIP: true,
				// LogHost instructs logger to extract request host value (i.e. `example.com`)
				LogHost: true,
				// LogMethod instructs logger to extract request method value (i.e. `GET` etc)
				LogMethod: true,
				// LogURIPath instructs logger to extract request URI path part (i.e. `/list`)
				LogURIPath: true,
				// LogRoutePath instructs logger to extract route path part to which request was matched to (i.e. `/user/:id`)
				LogRoutePath: true,
				// LogRequestID instructs logger to extract request ID from request `X-Request-ID` header or response if request did not have value.
				LogRequestID: true,
				// LogReferer instructs logger to extract request referer values.
				LogReferer: true,
				// LogUserAgent instructs logger to extract request user agent values.
				LogUserAgent: true,
				// LogError instructs logger to extract error returned from executed handler chain.
				LogError: true,
				// LogContentLength instructs logger to extract content length header value. Note: this value could be different from
				// actual request body size as it could be spoofed etc.
				LogContentLength: true,
				// LogResponseSize instructs logger to extract response content length value. Note: when used with Gzip middleware
				// this value may not be always correct.
				LogResponseSize: true,
				LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
					data, err := json.Marshal(v)
					if err != nil {
						return err
					}
					fmt.Println(string(data))
					return nil
				},
			},
		),
		auth.FirebaseAuth(),
		middleware.Recover(),
	)
}
