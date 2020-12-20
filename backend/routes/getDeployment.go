package routes

import (
	"deployed/datastore"
	"deployed/utils"
	"log"
	"net/http"
)

// GetDeployment returns deployment with given name
func GetDeployment(w http.ResponseWriter, r *http.Request) {
	authToken := r.Header.Get("Authorization")
	keys, ok := r.URL.Query()["name"]
	if !ok || len(keys[0]) == 0 {
		log.Println("No deployment name specified")
		utils.RespondWithError(w, http.StatusBadRequest, "No deployment name specified")
		return
	}
	deploymentName := keys[0]

	firestoreClient, err := datastore.Connect(authToken)
	if err != nil {
		log.Printf("Failed to open firestore client: %s\n", err.Error())
	}
	deployment, err := firestoreClient.GetDeploymentByName(deploymentName)
	if err != nil {
		log.Print("Failed to get deployment")
		utils.RespondWithError(w, http.StatusNotFound, "Failed to retrieve deployment with name "+deploymentName)
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, deployment)
	firestoreClient.Close()
}
