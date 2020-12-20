package routes

import (
	"deployed/datastore"
	"deployed/docker"
	"deployed/git"
	"deployed/hostconfiguration/certbot"
	"deployed/hostconfiguration/nginx"
	"deployed/utils"
	"log"
)

func failDeployment(msg string, err error, deployment *datastore.Deployment, firestoreClient *datastore.FirestoreClient) {
	log.Printf("%s: %s\n", msg, err.Error())
	deployment.Status = datastore.Deployment_ERROR
	firestoreClient.UpdateDeploymentStatus(deployment)
}

func updateNetworkConfigs(firestoreClient *datastore.FirestoreClient) error {
	deployedDomains, err := firestoreClient.GetAllDomains()
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

func deployCommit(deployment *datastore.Deployment, commit string, firestoreClient *datastore.FirestoreClient) error {
	deployment.Commit = commit
	err := firestoreClient.UpdateDeploymentCommit(deployment)
	if err != nil {
		failDeployment("Failed to store current commit", err, deployment, firestoreClient)
		return err
	}

	dockerfileLocation := git.GetRepoLocation(deployment.GetName()) + deployment.GetDockerfile()
	err = docker.BuildImage(dockerfileLocation, deployment.GetName(), commit)
	if err != nil {
		failDeployment("Failed to build image", err, deployment, firestoreClient)
		return err
	}

	metadata, err := docker.StartContainer(deployment.GetName(), commit, commit)
	if err != nil {
		failDeployment("Failed to start container", err, deployment, firestoreClient)
		return err
	}

	firestoreClient.AddContainer(deployment.GetName(), metadata)
	if err != nil {
		failDeployment("Failed to store container", err, deployment, firestoreClient)
		return err
	}

	if deployment.GetDomain() == "" || metadata.Port == nil {
		return err
	}

	domainConfig := &datastore.DomainConfiguration{
		Domain: deployment.GetDomain(),
		Port:   metadata.Port.HostPort,
	}
	err = firestoreClient.AddDomain(deployment.GetName(), domainConfig)
	if err != nil {
		failDeployment("Failed to add domain", err, deployment, firestoreClient)
		return err
	}

	err = updateNetworkConfigs(firestoreClient)
	if err != nil {
		failDeployment("Failed to update network configs", err, deployment, firestoreClient)
		return err
	}

	deployment.Status = datastore.Deployment_COMPLETE
	firestoreClient.UpdateDeploymentStatus(deployment)
	return nil
}

func updateNetworkConfigsAndClose(firestoreClient *datastore.FirestoreClient) {
	updateNetworkConfigs(firestoreClient)
	firestoreClient.Close()
}
