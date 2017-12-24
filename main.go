package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"

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

	router := handlers.Router(version.BuildTime, version.Commit, version.Release)

	log.Println("The service is ready to listen and serve.")
	log.Printf("port = %v\n", port)

	err := http.ListenAndServe(":"+strconv.Itoa(port), router)
	if err != nil {
		log.Fatal(err)
	}
}
