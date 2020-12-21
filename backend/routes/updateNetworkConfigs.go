package routes

import (
	"deployed/datastore"
	"deployed/utils"
	"log"
	"net/http"
)

// UpdateNetworkConfigs updates the Nginx and certificate configurations
func UpdateNetworkConfigs(w http.ResponseWriter, r *http.Request) {
	authToken := r.Header.Get("Authorization")
	firestoreClient, err := datastore.Connect(authToken)
	if err != nil {
		log.Printf("Failed to open firestore client: %s\n", err.Error())
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to connect to firestore")
		return
	}
	log.Printf("Updating network config...")
	go updateNetworkConfigsAndClose(firestoreClient)
	w.WriteHeader(http.StatusOK)
}
