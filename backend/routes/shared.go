package routes

import (
	"deployed/datastore"
	"deployed/hostconfiguration/certbot"
	"deployed/hostconfiguration/nginx"
	"deployed/utils"
	"log"
)

func failDeployment(msg string, err error, deployment *datastore.Deployment) {
	log.Printf("%s: %s\n", msg, err.Error())
	deployment.Status = datastore.Deployment_ERROR
	datastore.UpdateDeploymentStatus(deployment)
}

func updateNetworkConfigs() error {
	deployedDomains, err := datastore.GetAllDomains()
	if err != nil {
		return err
	}

	sitesEnabled := nginx.BuildSitesEnabled(deployedDomains)
	err = utils.WriteExistingFile("/etc/nginx/sites-enabled/brad.coffee", sitesEnabled)
	if err != nil {
		return err
	}

	err = certbot.UpdateCertificates(deployedDomains)
	if err != nil {
		return err
	}

	return nil
}
