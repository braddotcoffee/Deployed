package routes

import (
	"deployed/datastore"
	"deployed/git"
	"deployed/utils"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/golang/protobuf/ptypes"
)

// AddDeployment creates a new deployment from scratch
// and deploys it for the first time
func AddDeployment(w http.ResponseWriter, r *http.Request) {
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
	if deployment.GetRepository() == "" || deployment.GetDockerfile() == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Cannot deploy without repository and Dockerfile")
	}
	deployment.LastDeploy = ptypes.TimestampNow()
	if err := datastore.AddDeployment(deployment); err != nil {
		log.Fatalln("Failed to store deployment:", err)
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}

	go initializeDeployment(deployment)

	w.WriteHeader(http.StatusOK)
}

func initializeDeployment(deployment *datastore.Deployment) {
	deployment.Status = datastore.Deployment_IN_PROGRESS
	datastore.UpdateDeploymentStatus(deployment)

	hash, err := git.CloneRepoToLocation(deployment.GetRepository(), deployment.GetName())
	if err != nil {
		failDeployment("Failed to clone repo", err, deployment)
		return
	}

	deployCommit(deployment, hash)
}
