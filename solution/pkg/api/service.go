package api

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"solution/internal/storage"
	"time"
)

var NoUserFoundError = status.Error(codes.NotFound, "no user found")

//1. step ako funkcia bude vyzerat / pass db []
func getUserWithNearestBirthday(db storage.DBAccess) (storage.User, error) {
	users, err := db.ListUsers(context.Background())
	now := time.Now()
	var birthdayBoy *storage.User
	for _, u := range users {
		if birthdayBoy == nil {
			birthdayBoy = &u
		}
		if birthdayBoy.BornAt.Sub(now).Seconds() > u.BornAt.Sub(now).Seconds() {
			birthdayBoy = &u
		}
	}
	if birthdayBoy == nil {
		//return storage.User{}, errors.New("no user found")
		return storage.User{}, status.Error(codes.NotFound, "no user found")
	}
	return *birthdayBoy, err
}
