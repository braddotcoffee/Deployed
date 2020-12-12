package routes

import (
	"deployed/datastore"
	"deployed/docker"
	"deployed/git"
	"deployed/hostconfiguration"
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

func deployCommit(deployment *datastore.Deployment, commit string) error {
	deployment.Commit = commit
	err := datastore.UpdateDeploymentCommit(deployment)
	if err != nil {
		failDeployment("Failed to store current commit", err, deployment)
		return err
	}

	dockerfileLocation := git.GetRepoLocation(deployment.GetName()) + deployment.GetDockerfile()
	err = docker.BuildImage(dockerfileLocation, deployment.GetName(), commit)
	if err != nil {
		failDeployment("Failed to build image", err, deployment)
		return err
	}

	metadata, err := docker.StartContainer(deployment.GetName(), commit, commit)
	if err != nil {
		failDeployment("Failed to start container", err, deployment)
		return err
	}

	datastore.AddContainer(deployment.GetName(), metadata)
	if err != nil {
		failDeployment("Failed to store container", err, deployment)
		return err
	}

	if deployment.GetDomain() == "" || metadata.Port == nil {
		return err
	}

	domainConfig := &hostconfiguration.DomainConfiguration{
		Domain: deployment.GetDomain(),
		Port:   metadata.Port.HostPort,
	}
	err = datastore.AddDomain(deployment.GetName(), domainConfig)
	if err != nil {
		failDeployment("Failed to add domain", err, deployment)
		return err
	}

	err = updateNetworkConfigs()
	if err != nil {
		failDeployment("Failed to update network configs", err, deployment)
		return err
	}

	deployment.Status = datastore.Deployment_COMPLETE
	datastore.UpdateDeploymentStatus(deployment)
	return nil
}
