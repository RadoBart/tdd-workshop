package storage

import (
	"cloud.google.com/go/firestore"
	"context"
	"errors"
	"google.golang.org/api/iterator"
	"time"
)

// UserCollection holds user documents
const UserCollection = "users"

// User defines model of user in database
type User struct {
	ID     string    `firestore:"id"`
	Name   string    `firestore:"name"`
	BornAt time.Time `firestore:"bornAt"`
}

// GetUserCollection creates CollectionRef for user collection
func (fa *FirebaseAccess) GetUserCollection() *firestore.CollectionRef {
	return fa.Firestore.Collection(UserCollection)
}

// ListUsers returns list of users with pagination and filter for userIDs
func (fa *FirebaseAccess) ListUsers(ctx context.Context) (users []User, err error) {
	var docs *firestore.DocumentIterator
	if docs = fa.GetUserCollection().Documents(ctx); docs == nil {
		return users, errors.New("iterator is nil")
	}
	defer docs.Stop()
	for {
		user := User{}
		doc, err := docs.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return users, err
		}
		if err := doc.DataTo(&user); err != nil {
			return []User{}, err
		}
		users = append(users, user)
	}
	return users, nil
}
