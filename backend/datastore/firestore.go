package datastore

import (
	"context"
	"deployed/docker"
	"deployed/hostconfiguration"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

var client *firestore.Client = nil
var ctx context.Context = nil

// NewFirestoreClient opens a new connection to the
// Firestore associated with Deployed
func newFirestoreClient() (*firestore.Client, context.Context, error) {
	// Use a service account
	ctx := context.Background()
	sa := option.WithCredentialsFile("secrets/server-token.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
		return nil, nil, err
	}

	client, err := app.Firestore(ctx)
	return client, ctx, err
}

// Connect establishes connection to firebase
func Connect() error {
	var err error
	client, ctx, err = newFirestoreClient()
	if err != nil {
		log.Fatalln(err)
		return err
	}
	return nil
}

// Disconnect closes connection to firebase
func Disconnect() {
	client.Close()
	client = nil
	ctx = nil
}

// AddDeployment adds new deployment to the firestore
func AddDeployment(deployment *Deployment) error {
	_, err := client.Collection("deployments").Doc(deployment.GetName()).Set(ctx, deployment)
	if err != nil {
		log.Fatalf("Failed adding new deployment")
	}
	return err
}

// UpdateDeploymentCommit updates the commit field on the given deployment
func UpdateDeploymentCommit(deployment *Deployment) error {
	_, err := client.Collection("deployments").Doc(deployment.GetName()).Update(ctx, []firestore.Update{
		{
			Path:  "Commit",
			Value: deployment.GetCommit(),
		},
	})
	return err
}

// UpdateDeploymentStatus updates the status field on the given deployment
func UpdateDeploymentStatus(deployment *Deployment) error {
	_, err := client.Collection("deployments").Doc(deployment.GetName()).Update(ctx, []firestore.Update{
		{
			Path:  "Status",
			Value: deployment.GetStatus(),
		},
	})
	return err
}

// GetAllDeployments returns an array of all deployments tracked by Deployed
func GetAllDeployments() ([]*Deployment, error) {
	deployments := []*Deployment{}
	iter := client.Collection("deployments").OrderBy("LastDeploy", firestore.Desc).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return deployments, err
		}
		deployment := &Deployment{}
		err = doc.DataTo(deployment)
		if err != nil {
			return deployments, err
		}
		deployments = append(deployments, deployment)
	}
	return deployments, nil
}

// GetDeploymentByName gets the deployment corresponding to the given name
func GetDeploymentByName(name string) (*Deployment, error) {
	doc, err := client.Collection("deployments").Doc(name).Get(ctx)
	if err != nil {
		log.Printf("Failed to get deployment with name %s: %s\n", name, err.Error())
		return nil, err
	}

	deployment := &Deployment{}
	err = doc.DataTo(deployment)
	if err != nil {
		log.Printf("Failed to parse document into deployment: %s\n", err.Error())
		return nil, err
	}
	return deployment, nil
}

// AddContainer adds new container to the firestore
func AddContainer(application string, metadata *docker.ContainerMetadata) error {
	_, err := client.Collection("containers").Doc(application).Set(ctx, metadata)
	if err != nil {
		log.Fatalf("Failed adding container")
	}
	return err
}

// AddDomain adds new domain configuration to the firestore
func AddDomain(application string, domainConfig *hostconfiguration.DomainConfiguration) error {
	_, err := client.Collection("domains").Doc(application).Set(ctx, domainConfig)
	if err != nil {
		log.Fatalf("Failed adding domain")
	}
	return err
}

// GetAllDomains gets all of the domains currently tracked by Deployed
func GetAllDomains() ([]hostconfiguration.DomainConfiguration, error) {
	domains := []hostconfiguration.DomainConfiguration{}
	iter := client.Collection("domains").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return domains, err
		}
		rawData := doc.Data()
		domains = append(domains, hostconfiguration.DomainConfiguration{
			Domain: rawData["Domain"].(string),
			Port:   rawData["Port"].(string),
		})
	}
	return domains, nil
}
