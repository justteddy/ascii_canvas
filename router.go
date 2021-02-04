package main

import (
	"net/http"
	"time"

	"canvas/handler"
	"canvas/storage"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func newRouter(redis *storage.Redis, canvasTTL time.Duration) http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.AllowContentType("application/json"))
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	h := handler.New(redis, canvasTTL)

	router.Get("/canvas/{canvasID}", h.HandleGet)
	router.Put("/canvas/{canvasID}/floodFill", h.HandleFloodFill)
	router.Put("/canvas/{canvasID}/drawRectangle", h.HandleDrawRectangle)

	return router
}
