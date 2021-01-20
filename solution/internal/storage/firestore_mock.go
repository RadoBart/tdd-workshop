package storage

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
)

type FirebaseMock struct {
	Users []User
}

func (f FirebaseMock) GetUserCollection() *firestore.CollectionRef {
	return nil
}

func (f FirebaseMock) ListUsers(ctx context.Context) ([]User, error) {
	return f.Users, nil
}

// FirebaseTesting contains tools necessary for testing firebase functionality
type FirebaseTesting struct {
	FirebaseAccess
}

// NewFirebaseTesting factory method for FirebaseTesting creation
func NewFirebaseTesting(ctx context.Context) (*FirebaseTesting, error) {
	// init firebase firestore
	firebaseClient, err := firebase.NewApp(ctx, &firebase.Config{
		ProjectID: "test",
	})
	if err != nil {
		return nil, err
	}

	// init firestore firestore from firebase firestore
	firestoreClient, err := firebaseClient.Firestore(ctx)
	if err != nil {
		return nil, err
	}

	return &FirebaseTesting{FirebaseAccess{Firestore: firestoreClient}}, nil
}

// region <<<Firestore Data Generators>>>

// GenerateUser generates user for testing
func (ft *FirebaseTesting) GenerateUser(ctx context.Context, user User) (*firestore.DocumentRef, error) {
	mockedItem := ft.Firestore.
		Collection(UserCollection).Doc(user.ID)
	if _, err := mockedItem.Set(ctx, user); err != nil {
		return nil, err
	}
	return mockedItem, nil
}

// endregion  <<<Firestore Data Generators>>>
