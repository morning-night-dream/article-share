package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/morning-night-dream/platform/internal/adapter/controller"
	"github.com/morning-night-dream/platform/internal/driver/client"
	"github.com/morning-night-dream/platform/internal/driver/config"
	"github.com/morning-night-dream/platform/internal/driver/server"
	"github.com/morning-night-dream/platform/pkg/openapi"
)

func main() {
	c, err := client.New().Of(config.Gateway.AppCoreURL)
	if err != nil {
		panic(err)
	}

	ctr := controller.New(c)

	handler := openapi.HandlerWithOptions(ctr, openapi.ChiServerOptions{
		BaseURL:     "/api",
		BaseRouter:  chi.NewRouter(),
		Middlewares: []openapi.MiddlewareFunc{server.Middleware},
	})

	srv := server.NewHTTPServer(handler)

	srv.Run()
}
