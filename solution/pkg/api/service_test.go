package api

import (
	"context"
	"errors"
	"github.com/stretchr/testify/suite"
	"net/http"
	"solution/internal/storage"
	"testing"
	"time"
)

type ServiceSuite struct {
	suite.Suite
	//firebase storage.DBAccess
	firebase *storage.FirebaseTesting
	ctx      context.Context

	//testUUID  string
	//ctx       context.Context
}

// region <<ID generation>>

//const idForGeneratedObjects = "generated"

func (s *ServiceSuite) UserID(suffix string) string {
	//s.testUUID +
	return "-user-" + suffix
}

// endregion <<ID generation>>

func (s *ServiceSuite) SetupTest() {
	var err error
	//s.firebase = storage.FirebaseMock{
	//	[]storage.User{
	//		{
	//			ID:     "1",
	//			Name:   "",
	//			BornAt: time.Now().Add(5),
	//		},
	//		{
	//			ID:     "2",
	//			Name:   "",
	//			BornAt: time.Now(),
	//		},
	//	},
	//}
	s.ctx = context.Background()
	s.firebase, err = storage.NewFirebaseTesting(s.ctx)
	if err != nil {
		panic(err)
	}

}

func (s *ServiceSuite) TearDownTest() {
	// region <<Database Cleanup>>
	if err := s.clearEmulatorDb(); err != nil {
		s.Nil(err)
		return
	}
	// endregion
}

// region <<Database Initialization>>
func (s *ServiceSuite) generateUsers() {
	birthdayBoy := storage.User{
		ID:     s.UserID("2"),
		Name:   "Random boy",
		BornAt: time.Now(),
	}
	_, err := s.firebase.GenerateUser(s.ctx, birthdayBoy)
	s.Nil(err)
	notBirthdayBoy := storage.User{
		ID:     s.UserID("1"),
		Name:   "Random boy2",
		BornAt: time.Now().Add(5),
	}
	_, err = s.firebase.GenerateUser(s.ctx, notBirthdayBoy)
	s.Nil(err)
}

func (s *ServiceSuite) clearEmulatorDb() (err error) {
	emulatorClearDbRequest, err := http.NewRequest(
		http.MethodDelete, "http://localhost:42042/emulator/v1/projects/test/databases/(default)/documents", nil,
	)
	client := http.Client{}
	response, err := client.Do(emulatorClearDbRequest)
	if err != nil {
		s.Nil(err)
		return err
	} else if response == nil {
		s.NotNil(response)
		return errors.New("response is nil")
	}
	s.Equal(response.StatusCode, http.StatusOK)
	return nil
}

// endregion  <<Database Initialization>>

// region <<Token>>

func (s *ServiceSuite) TestBirthdayBoy() {
	s.generateUsers()
	birthdayBoy, err := getUserWithNearestBirthday(s.firebase)
	s.Nil(err)
	//s.Equal("2", birthdayBoy.ID)
	s.Equal(s.UserID("2"), birthdayBoy.ID)
}

func (s *ServiceSuite) TestNoBoysInDb() {
	//s.firebase = storage.FirebaseMock{
	//}
	birthdayBoy, err := getUserWithNearestBirthday(s.firebase)
	//s.Error(err, errors.New("no user found"))
	s.Error(err, NoUserFoundError)
	s.Equal(storage.User{}, birthdayBoy)
}

// endregion <<Configs>>
func TestReleaseServiceSuite(t *testing.T) {
	suite.Run(t, &ServiceSuite{})
}
