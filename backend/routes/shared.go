package routes

import (
	"deployed/datastore"
	"deployed/docker"
	"deployed/git"
	"deployed/hostconfiguration/certbot"
	"deployed/hostconfiguration/nginx"
	"deployed/utils"
	"log"
	"os"
	"os/exec"

	"github.com/otiai10/copy"
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

func deployDockerImage(deployment *datastore.Deployment, commit string, firestoreClient *datastore.FirestoreClient) error {
	dockerfileLocation := git.GetRepoLocation(deployment.GetName()) + deployment.GetDockerfile()
	err := docker.BuildImage(dockerfileLocation, deployment.GetName(), commit)
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

	return nil
}

func deployManualBuild(deployment *datastore.Deployment, firestoreClient *datastore.FirestoreClient) error {
	baseDir := git.GetRepoLocation(deployment.GetName())
	cmd := exec.Command("/bin/sh", "-c", deployment.GetBuildCommand())
	cmd.Dir = baseDir
	cmd.Stdout = os.Stdout
	log.Println(deployment.GetBuildCommand())
	if err := cmd.Run(); err != nil {
		failDeployment("Failed to build application", err, deployment, firestoreClient)
		return err
	}

	// Assume that we will deploy a domain config manually for this project
	if deployment.GetOutputDirectory() == "" {
		return nil
	}

	if err := copy.Copy(baseDir+deployment.GetOutputDirectory(), "/var/www/"+deployment.GetName()); err != nil {
		failDeployment("Failed to copy application to nginx dir", err, deployment, firestoreClient)
		return err
	}
	domainConfig := &datastore.DomainConfiguration{
		Domain:           deployment.GetDomain(),
		Port:             "",
		ForwardDirectory: deployment.GetName(),
	}
	if err := firestoreClient.AddDomain(deployment.GetName(), domainConfig); err != nil {
		failDeployment("Failed to add domain", err, deployment, firestoreClient)
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

	if deployment.GetDockerfile() != "" {
		if err := deployDockerImage(deployment, commit, firestoreClient); err != nil {
			return err
		}
	}

	if deployment.GetBuildCommand() != "" {
		if err := deployManualBuild(deployment, firestoreClient); err != nil {
			return err
		}
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
