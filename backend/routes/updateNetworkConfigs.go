package routes

import (
	"net/http"
)

// UpdateNetworkConfigs updates the Nginx and certificate configurations
func UpdateNetworkConfigs(w http.ResponseWriter, r *http.Request) {
	go updateNetworkConfigs()
	w.WriteHeader(http.StatusOK)
}
