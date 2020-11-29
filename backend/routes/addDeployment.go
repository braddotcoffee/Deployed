package routes

import (
	"deployed/datastore"
	"deployed/docker"
	"deployed/git"
	"deployed/hostconfiguration"
	"deployed/hostconfiguration/certbot"
	"deployed/hostconfiguration/nginx"
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
	hash, err := git.CloneRepoToLocation(deployment.GetRepository(), deployment.GetName())
	if err != nil {
		failDeployment("Failed to clone repo", err, deployment)
		return
	}

	dockerfileLocation := git.GetRepoLocation(deployment.GetName()) + deployment.GetDockerfile()
	err = docker.BuildImage(dockerfileLocation, deployment.GetName(), hash)
	if err != nil {
		failDeployment("Failed to build image", err, deployment)
		return
	}

	metadata, err := docker.StartContainer(deployment.GetName(), hash, hash)
	if err != nil {
		failDeployment("Failed to start container", err, deployment)
		return
	}

	datastore.AddContainer(deployment.GetName(), metadata)
	if err != nil {
		failDeployment("Failed to store container", err, deployment)
		return
	}

	if deployment.GetDomain() == "" || metadata.Port == nil {
		return
	}

	domainConfig := &hostconfiguration.DomainConfiguration{
		Domain: deployment.GetDomain(),
		Port:   metadata.Port.HostPort,
	}
	err = datastore.AddDomain(deployment.GetName(), domainConfig)
	if err != nil {
		failDeployment("Failed to add domain", err, deployment)
		return
	}

	deployedDomains, err := datastore.GetAllDomains()
	if err != nil {
		failDeployment("Failed to get all domains", err, deployment)
		return
	}

	allDomains := append(hostconfiguration.DefaultDomains, deployedDomains...)
	sitesEnabled := nginx.BuildSitesEnabled(allDomains)
	err = utils.WriteExistingFile("/etc/nginx/sites-enabled/brad.coffee", sitesEnabled)
	if err != nil {
		failDeployment("Failed to write nginx conf", err, deployment)
		return
	}

	err = certbot.UpdateCertificates(allDomains)
	if err != nil {
		failDeployment("Failed to update certificate", err, deployment)
		return
	}
}
