package datastore

import (
	"context"
	"deployed/docker"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

// FirestoreClient wraps all interactions with the firestore
type FirestoreClient struct {
	client *firestore.Client
	ctx    context.Context
}

// NewFirestoreClient opens a new connection to the
// Firestore associated with Deployed
func newFirestoreClient(tokenSource *FirebaseTokenSource) (*firestore.Client, context.Context, error) {
	// Use a service account
	ctx := context.Background()
	creds := option.WithTokenSource(tokenSource)
	app, err := firebase.NewApp(ctx, &firebase.Config{
		ProjectID: "deployed-d4c32",
	}, creds)
	if err != nil {
		log.Printf("Failed to connect to firestore: %s\n", err.Error())
		return nil, nil, err
	}

	client, err := app.Firestore(ctx)
	return client, ctx, err
}

// Connect establishes a new firestore client
func Connect(token string) (*FirestoreClient, error) {
	tokenSource := &FirebaseTokenSource{
		token: token,
	}
	client, ctx, err := newFirestoreClient(tokenSource)
	if err != nil {
		log.Printf("Failed to connect to firestore: %s\n", err.Error())
		return nil, err
	}
	return &FirestoreClient{
		client: client,
		ctx:    ctx,
	}, nil
}

// Close all sessions to firebase for client
func (fc FirestoreClient) Close() {
	fc.client.Close()
}

// AddDeployment adds new deployment to the firestore
func (fc FirestoreClient) AddDeployment(deployment *Deployment) error {
	_, err := fc.client.Collection("deployments").Doc(deployment.GetName()).Set(fc.ctx, deployment)
	if err != nil {
		log.Printf("Failed adding new deployment")
	}
	return err
}

// UpdateDeploymentCommit updates the commit field on the given deployment
func (fc FirestoreClient) UpdateDeploymentCommit(deployment *Deployment) error {
	_, err := fc.client.Collection("deployments").Doc(deployment.GetName()).Update(fc.ctx, []firestore.Update{
		{
			Path:  "Commit",
			Value: deployment.GetCommit(),
		},
	})
	return err
}

// UpdateDeploymentStatus updates the status field on the given deployment
func (fc FirestoreClient) UpdateDeploymentStatus(deployment *Deployment) error {
	_, err := fc.client.Collection("deployments").Doc(deployment.GetName()).Update(fc.ctx, []firestore.Update{
		{
			Path:  "Status",
			Value: deployment.GetStatus(),
		},
	})
	return err
}

// GetAllDeployments returns an array of all deployments tracked by Deployed
func (fc FirestoreClient) GetAllDeployments() ([]*Deployment, error) {
	deployments := []*Deployment{}
	iter := fc.client.Collection("deployments").OrderBy("LastDeploy", firestore.Desc).Documents(fc.ctx)
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
func (fc FirestoreClient) GetDeploymentByName(name string) (*Deployment, error) {
	doc, err := fc.client.Collection("deployments").Doc(name).Get(fc.ctx)
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
func (fc FirestoreClient) AddContainer(application string, metadata *docker.ContainerMetadata) error {
	_, err := fc.client.Collection("containers").Doc(application).Set(fc.ctx, metadata)
	if err != nil {
		log.Printf("Failed adding container")
	}
	return err
}

// AddDomain adds new domain configuration to the firestore
func (fc FirestoreClient) AddDomain(application string, domainConfig *DomainConfiguration) error {
	_, err := fc.client.Collection("domains").Doc(application).Set(fc.ctx, domainConfig)
	if err != nil {
		log.Printf("Failed adding domain")
	}
	return err
}

// GetAllDomains gets all of the domains currently tracked by Deployed
func (fc FirestoreClient) GetAllDomains() ([]*DomainConfiguration, error) {
	domains := []*DomainConfiguration{}
	iter := fc.client.Collection("domains").Documents(fc.ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return domains, err
		}
		domainConfiguration := &DomainConfiguration{}
		err = doc.DataTo(domainConfiguration)
		if err != nil {
			log.Printf("Failed to parse document into deployment: %s\n", err.Error())
			return nil, err
		}
		domains = append(domains, domainConfiguration)
	}
	return domains, nil
}
