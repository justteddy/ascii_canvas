package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"canvas/handlers/floodfill"
	"canvas/handlers/get"
	"canvas/handlers/post"
	"canvas/handlers/rectangle"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var (
	port = flag.String("port", ":8080", "http port")
)

func main() {
	flag.Parse()

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Post("/canvas/{canvasID}", post.New().Handle)
	router.Get("/canvas/{canvasID}", get.New().Handle)
	router.Put("/canvas/{canvasID}/floodFill", floodfill.New().Handle)
	router.Put("/canvas/{canvasID}/drawRectangle", rectangle.New().Handle)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)
	go func() {
		sig := <-sigCh
		log.Printf("interrupted with [%s] signal, bye", sig)
		os.Exit(0)
	}()

	log.Printf("ready to accept connections on %s port", *port)
	if err := http.ListenAndServe(*port, router); err != nil {
		log.Fatal(err)
	}
}
