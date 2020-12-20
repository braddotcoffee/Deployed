package routes

import (
	"deployed/datastore"
	"deployed/utils"
	"log"
	"net/http"
)

// GetDeployments returns all deployments
func GetDeployments(w http.ResponseWriter, r *http.Request) {
	authToken := r.Header.Get("Authorization")
	firestoreClient, err := datastore.Connect(authToken)
	if err != nil {
		log.Printf("Failed to open firestore client: %s\n", err.Error())
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to connect to firestore")
		return
	}
	deployments, err := firestoreClient.GetAllDeployments()
	if err != nil {
		log.Fatalf("Failed to get deployments: %s\n", err.Error())
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to retrieve deployments")
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, deployments)
	firestoreClient.Close()
}
