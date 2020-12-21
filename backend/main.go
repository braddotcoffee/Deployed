package main

import (
	"deployed/docker"
	"deployed/routes"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()
	docker.Connect()

	r.HandleFunc("/add-deployment", routes.AddDeployment)
	r.HandleFunc("/add-domain-config", routes.AddDomainConfig)
	r.HandleFunc("/get-deployments", routes.GetDeployments)
	r.HandleFunc("/get-deployment", routes.GetDeployment)
	r.HandleFunc("/deploy-new-version", routes.DeployNewVersion)
	r.HandleFunc("/update-network-config", routes.UpdateNetworkConfigs)

	frontendURL := ""

	if prod := os.Getenv("PRODUCTION"); prod != "" {
		frontendURL = "https://deployed.brad.coffee"
	} else {
		frontendURL = "https://app.brad.coffee"
	}

	c := cors.New(cors.Options{
		AllowedOrigins: []string{frontendURL},
		AllowedHeaders: []string{"Authorization", "Content-Type"},
	})
	handler := c.Handler(r)

	srv := &http.Server{
		Handler: handler,
		Addr:    ":" + os.Getenv("PORT"),
	}

	srv.ListenAndServe()
}
