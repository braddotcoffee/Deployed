package routes

import (
	"deployed/datastore"
	"deployed/utils"
	"log"
	"net/http"

	"github.com/gogo/protobuf/jsonpb"
)

// AddDomainConfig creates a new domain configuration from scratch
// and deploys the associated app for the first time
func AddDomainConfig(w http.ResponseWriter, r *http.Request) {
	authToken := r.Header.Get("Authorization")
	domainConfig := &datastore.DomainConfiguration{}
	if err := jsonpb.Unmarshal(r.Body, domainConfig); err != nil {
		log.Printf("Failed to parse domain configuration: %s\n", err.Error())
		utils.RespondWithError(w, http.StatusBadRequest, "Unable to parse domain configuration")
	}

	if domainConfig.GetDomain() == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Cannot add domain config without domain")
		return
	}

	if domainConfig.GetPort() == "" && domainConfig.GetForwardDirectory() == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Cannot add domain config without port or forward directory")
		return
	}

	keys, ok := r.URL.Query()["name"]
	if !ok || len(keys[0]) == 0 {
		log.Println("No application name specified")
		utils.RespondWithError(w, http.StatusBadRequest, "No application name specified")
		return
	}
	applicationName := keys[0]
	firestoreClient, err := datastore.Connect(authToken)
	if err != nil {
		log.Printf("Failed to open firestore client: %s\n", err.Error())
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to connect to firestore")
		return
	}
	err = firestoreClient.AddDomain(applicationName, domainConfig)
	if err != nil {
		log.Printf("Failed to store domain configuration: %s\n", err.Error())
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to store domain configuration")
	}

	go updateNetworkConfigsAndClose(firestoreClient)
	w.WriteHeader(http.StatusOK)
}
