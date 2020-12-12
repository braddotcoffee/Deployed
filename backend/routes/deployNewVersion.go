package routes

import (
	"deployed/datastore"
	"deployed/git"
	"deployed/utils"
	"log"
	"net/http"
)

// DeployNewVersion deploys the latest commit associated with the given deployment
func DeployNewVersion(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["name"]
	if !ok || len(keys[0]) == 0 {
		log.Println("No deployment name specified")
		utils.RespondWithError(w, http.StatusBadRequest, "No deployment name specified")
		return
	}
	deploymentName := keys[0]
	log.Printf("Deploying new version: %s\n", deploymentName)
	deployment, err := datastore.GetDeploymentByName(deploymentName)
	if err != nil {
		log.Print("Failed to get deployment")
		utils.RespondWithError(w, http.StatusNotFound, "Failed to retrieve deployment with name "+deploymentName)
		return
	}
	go pullRepoAndDeploy(deployment)
	w.WriteHeader(http.StatusOK)
}

func pullRepoAndDeploy(deployment *datastore.Deployment) {
	deployment.Status = datastore.Deployment_IN_PROGRESS
	datastore.UpdateDeploymentStatus(deployment)

	commit, err := git.PullRepoAtLocation(deployment.GetName())
	if err != nil {
		if err.Error() == "already up-to-date" {
			deployment.Status = datastore.Deployment_COMPLETE
			datastore.UpdateDeploymentStatus(deployment)
			log.Printf("No new version to deploy: %s\n", deployment.GetName())
			return
		}
		failDeployment("Failed to pull repo at location", err, deployment)
		return
	}
	deployCommit(deployment, commit)
}
