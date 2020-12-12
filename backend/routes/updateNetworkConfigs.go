package routes

import (
	"deployed/utils"
	"log"
	"net/http"
)

// UpdateNetworkConfigs updates the Nginx and certificate configurations
func UpdateNetworkConfigs(w http.ResponseWriter, r *http.Request) {
	err := updateNetworkConfigs()
	if err != nil {
		log.Printf("Failed to update network configs")
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}
	w.WriteHeader(http.StatusOK)
}
