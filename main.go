package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"

	"./handlers"
)

var (
	port int
)

func main() {
	log.Print("Starting the service...")

	flag.IntVar(&port, "port", 8080, "server port.")
	flag.Parse()

	router := handlers.Router()

	log.Println("The service is ready to listen and serve.")
	log.Printf("port = %v\n", port)

	err := http.ListenAndServe(":"+strconv.Itoa(port), router)
	if err != nil {
		log.Fatal(err)
	}
}
