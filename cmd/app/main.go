package main

import (
	"log"
	"net/http"
	"randomgifsite/internal/transport/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RandomGifSiteRouter(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
