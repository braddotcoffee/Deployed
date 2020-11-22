package datastore

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
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
