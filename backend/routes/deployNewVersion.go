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
	authToken := r.Header.Get("Authorization")
	keys, ok := r.URL.Query()["name"]
	if !ok || len(keys[0]) == 0 {
		log.Println("No deployment name specified")
		utils.RespondWithError(w, http.StatusBadRequest, "No deployment name specified")
		return
	}
	deploymentName := keys[0]
	log.Printf("Deploying new version: %s\n", deploymentName)
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
	go pullRepoAndDeploy(deployment, firestoreClient)
	w.WriteHeader(http.StatusOK)
}

func pullRepoAndDeploy(deployment *datastore.Deployment, firestoreClient *datastore.FirestoreClient) {
	deployment.Status = datastore.Deployment_IN_PROGRESS
	firestoreClient.UpdateDeploymentStatus(deployment)

	commit, err := git.PullRepoAtLocation(deployment.GetName())
	if err != nil {
		if err.Error() != "already up-to-date" {
			failDeployment("Failed to pull repo at location", err, deployment, firestoreClient)
			return
		}
		commit, err = git.GetCurrentCommit(deployment.GetName())
		if err != nil {
			failDeployment("Failed to get current commit", err, deployment, firestoreClient)
			return
		}
	}
	deployCommit(deployment, commit, firestoreClient)
	firestoreClient.Close()
}
