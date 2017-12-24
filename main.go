package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"./handlers"
	"./version"
)

var (
	port      int
	buildTime string
	commit    string
	release   string
)

func main() {
	version.BuildTime = buildTime
	version.Commit = commit
	version.Release = release

	log.Printf(
		"Starting the service...\ncommit: %s, build time: %s, release: %s\n",
		version.Commit, version.BuildTime, version.Release,
	)

	flag.IntVar(&port, "port", 8080, "server port.")
	flag.Parse()

	r := handlers.Router(version.BuildTime, version.Commit, version.Release)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: r,
	}
	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()
	log.Print("The service is ready to listen and serve.")
	log.Printf("port = %v\n", port)

	killSignal := <-interrupt
	switch killSignal {
	case os.Interrupt:
		log.Print("Got SIGINT...")
	case syscall.SIGTERM:
		log.Print("Got SIGTERM...")
	}

	log.Print("The service is shutting down...")
	srv.Shutdown(context.Background())
	log.Print("Done")
}
