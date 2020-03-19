package main

import (
	"log"
	"net/http"

	"github.com/clarke94/weigh-me/srv/router"
	"github.com/go-chi/chi"
)

func main() {
	router := router.Routes()

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)
		return nil
	}

	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Logging error: %s\n", err.Error())
	}

	log.Print("Starting server on port 3000...")
	log.Fatal(http.ListenAndServe(":3000", router))
}
