package routes

import (
	"deployed/datastore"
	"deployed/utils"
	"log"
	"net/http"
)

// GetDeployments returns all deployments
func GetDeployments(w http.ResponseWriter, r *http.Request) {
	deployments, err := datastore.GetAllDeployments()
	if err != nil {
		log.Fatalf("Failed to get deployments: %s\n", err.Error())
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to retrieve deployments")
	}
	utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"Deployments": deployments,
	})
}
