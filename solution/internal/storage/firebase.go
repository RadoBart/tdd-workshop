package storage

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
)

type DBAccess interface {
	GetUserCollection() *firestore.CollectionRef
	ListUsers(ctx context.Context) ([]User, error)
}

// FirebaseAccess represents a collection of Firebase tools used during the application runtime
type FirebaseAccess struct {
	Firestore *firestore.Client
	Firebase  *firebase.App
	Auth      *auth.Client
}

// NewFirebaseAccess creates new FirebaseAccess object using given service account.
// Accepts optionally serviceAccount path.
// Nil can be passed, if this factory function is used in the testing environment and emulator.
// Reference to the documentation of emulator: https://cloud.google.com/sdk/gcloud/reference/beta/emulators
func NewFirebaseAccess(ctx context.Context, projectID string) (*FirebaseAccess, error) {
	var (
		firebaseClient *firebase.App
		err            error
	)
	// Init firebase firestore
	firebaseClient, err = firebase.NewApp(ctx, &firebase.Config{
		ProjectID: projectID,
	})
	if err != nil {
		return nil, err
	}

	// Init firestore firestore from firebase firestore
	firestoreClient, err := firebaseClient.Firestore(ctx)
	if err != nil {
		return nil, err
	}

	// Init firebase auth firestore from firebase firestore
	firebaseAuth, err := firebaseClient.Auth(ctx)
	if err != nil {
		return nil, err
	}

	return &FirebaseAccess{firestoreClient, firebaseClient, firebaseAuth}, nil
}
