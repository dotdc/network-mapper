package resolvers

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/labstack/echo/v4"
	"github.com/otterize/network-mapper/mapper/pkg/graph/generated"
	"github.com/otterize/network-mapper/mapper/pkg/kubefinder"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	kubeFinder    *kubefinder.KubeFinder
	intentsHolder *intentsHolder
}

func NewResolver(kubeFinder *kubefinder.KubeFinder) *Resolver {
	return &Resolver{
		kubeFinder:    kubeFinder,
		intentsHolder: NewIntentsHolder(),
	}
}

func (r *Resolver) Register(e *echo.Echo) {
	c := generated.Config{Resolvers: r}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(c))
	e.Any("/query", func(c echo.Context) error {
		srv.ServeHTTP(c.Response(), c.Request())
		return nil
	})
}