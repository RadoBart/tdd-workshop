package api

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type ServiceSuite struct {
	suite.Suite
}

func (s *ServiceSuite) SetupTest() {

}

func (s *ServiceSuite) TearDownTest() {

}

func (s *ServiceSuite) TestFirstTest() {

}

// endregion <<Configs>>
func TestReleaseServiceSuite(t *testing.T) {
	suite.Run(t, &ServiceSuite{})
}
