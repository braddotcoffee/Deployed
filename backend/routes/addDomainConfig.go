package routes

import (
	"deployed/datastore"
	"deployed/utils"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// AddDomainConfig creates a new domain configuration from scratch
// and deploys the associated app for the first time
func AddDomainConfig(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Failed to read body of request: %s\n", err.Error())
	}
	domainConfig := &datastore.DomainConfiguration{}
	if err := json.Unmarshal(body, domainConfig); err != nil {
		log.Printf("Failed to parse domain configuration: %s\n", err.Error())
		utils.RespondWithError(w, http.StatusBadRequest, "Unable to parse domain configuration")
	}

	if domainConfig.GetDomain() == "" || domainConfig.GetPort() == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Cannot add domain config without domain and port")
	}

	keys, ok := r.URL.Query()["name"]
	if !ok || len(keys[0]) == 0 {
		log.Println("No application name specified")
		utils.RespondWithError(w, http.StatusBadRequest, "No application name specified")
		return
	}
	applicationName := keys[0]
	err = datastore.AddDomain(applicationName, domainConfig)
	if err != nil {
		log.Printf("Failed to store domain configuration: %s\n", err.Error())
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to store domain configuration")
	}

	go updateNetworkConfigs()
	w.WriteHeader(http.StatusOK)
}
