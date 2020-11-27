package main

import (
	"deployed/datastore"
	"deployed/docker"
	"deployed/git"
	"deployed/utils"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()
	datastore.Connect()

	r.HandleFunc("/add-deployment", addDeployment)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"https://app.brad.coffee"},
	})
	handler := c.Handler(r)

	srv := &http.Server{
		Handler: handler,
		Addr:    ":" + os.Getenv("PORT"),
	}

	git.CloneRepoToLocation("git@github.com:crscillitoe/DiscordBotsToCleanseYourSoul.git", "~/DiscordBots")
	docker.BuildImage("~/DiscordBots/MarkovBot/Dockerfile", "test:latest")

	log.Fatal(srv.ListenAndServe())
}

func addDeployment(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln("Failed to read body of request:", err)
		utils.RespondWithError(w, http.StatusBadRequest, "Unable to read body of request")
	}

	deployment := &datastore.Deployment{}
	if err := json.Unmarshal(body, deployment); err != nil {
		log.Fatalln("Failed to parse deployment:", err)
		utils.RespondWithError(w, http.StatusBadRequest, "Unable to parse deployment")
	}
	if err := datastore.AddDeployment(deployment); err != nil {
		log.Fatalln("Failed to store deployment:", err)
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}
	w.WriteHeader(http.StatusOK)
}
