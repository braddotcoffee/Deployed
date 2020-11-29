package routes

import (
	"deployed/datastore"
	"log"
)

func failDeployment(msg string, err error, deployment *datastore.Deployment) {
	log.Printf("%s: %s\n", msg, err.Error())
	deployment.Status = datastore.Deployment_ERROR
	datastore.UpdateDeploymentStatus(deployment)
}
