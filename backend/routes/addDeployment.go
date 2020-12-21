package routes

import (
	"deployed/datastore"
	"deployed/git"
	"deployed/utils"
	"log"
	"net/http"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/golang/protobuf/ptypes"
)

// AddDeployment creates a new deployment from scratch
// and deploys it for the first time
func AddDeployment(w http.ResponseWriter, r *http.Request) {
	authToken := r.Header.Get("Authorization")
	deployment := &datastore.Deployment{}
	if err := jsonpb.Unmarshal(r.Body, deployment); err != nil {
		log.Fatalln("Failed to parse deployment:", err)
		utils.RespondWithError(w, http.StatusBadRequest, "Unable to parse deployment")
		return
	}
	if deployment.GetRepository() == "" {
		log.Printf("Cannot deploy without repository")
		utils.RespondWithError(w, http.StatusBadRequest, "Cannot deploy without repository")
		return
	}

	if deployment.GetDockerfile() == "" && deployment.GetBuildCommand() == "" {
		log.Printf("Cannot deploy without dockerfile or build command")
		utils.RespondWithError(w, http.StatusBadRequest, "Cannot deploy without dockerfile or build command")
		return
	}
	deployment.LastDeploy = ptypes.TimestampNow()
	firestoreClient, err := datastore.Connect(authToken)
	if err != nil {
		log.Printf("Failed to open firestore client: %s\n", err.Error())
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to connect to firestore")
		return
	}
	if err := firestoreClient.AddDeployment(deployment); err != nil {
		log.Fatalln("Failed to store deployment:", err)
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	go initializeDeployment(deployment, firestoreClient)

	w.WriteHeader(http.StatusOK)
}

func initializeDeployment(deployment *datastore.Deployment, firestoreClient *datastore.FirestoreClient) {
	deployment.Status = datastore.Deployment_IN_PROGRESS
	firestoreClient.UpdateDeploymentStatus(deployment)

	hash, err := git.CloneRepoToLocation(deployment.GetRepository(), deployment.GetName())
	if err != nil {
		failDeployment("Failed to clone repo", err, deployment, firestoreClient)
		return
	}

	deployCommit(deployment, hash, firestoreClient)
	firestoreClient.Close()
}
