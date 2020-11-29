package main

import (
	"deployed/datastore"
	"deployed/docker"
	"deployed/routes"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()
	datastore.Connect()
	docker.Connect()

	r.HandleFunc("/add-deployment", routes.AddDeployment)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"https://app.brad.coffee"},
	})
	handler := c.Handler(r)

	srv := &http.Server{
		Handler: handler,
		Addr:    ":" + os.Getenv("PORT"),
	}

	log.Fatal(srv.ListenAndServe())
}
