package main

import (
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"canvas/storage"

	log "github.com/sirupsen/logrus"
)

var (
	port          = flag.String("port", ":8080", "http port")
	redisAddr     = flag.String("redis-addr", "redis:6379", "Redis address")
	redisTimeout  = flag.Duration("redis-timeout", time.Second, "Redis timeout")
	redisPoolSize = flag.Int("redis-pool-size", 10, "Redis connections pool size")
	canvasTTL     = flag.Duration("canvas-ttl", time.Hour*24, "Canvas ttl")
)

func main() {
	flag.Parse()
	log.SetFormatter(&log.TextFormatter{})

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)
	go func() {
		sig := <-sigCh
		log.Printf("interrupted with [%s] signal, bye", sig)
		os.Exit(0)
	}()

	redis, err := storage.New(*redisAddr, *redisTimeout, *redisPoolSize)
	if err != nil {
		log.Fatal(err)
	}

	router := newRouter(redis, *canvasTTL)
	log.Printf("ready to accept connections on %s port", *port)

	if err := http.ListenAndServe(*port, router); err != nil {
		log.Fatal(err)
	}
}
