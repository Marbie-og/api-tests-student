package advertisementcreate

import (
	"api-tests-template/internal/managers/auth"
	"api-tests-template/internal/managers/auth/models"
	base "api-tests-template/tests"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	login models.LoginOkResponse
}

func TestSuiteRun(t *testing.T) { suite.Run(t, &TestSuite{}) }

func (s *TestSuite) SetupSuite() {
	base.SetupSuite()
	base.Precondition("Авторизация владельца тестовых объявлений")
	s.login = auth.Login(s.T(), os.Getenv("TEST_LOGIN"), os.Getenv("TEST_PASSWORD"))
}

func (s *TestSuite) TearDownSuite() { base.TearDownSuite() }
