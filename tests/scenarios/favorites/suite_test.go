package favorites

import (
	"api-tests-template/internal/managers/auth"
	"api-tests-template/internal/managers/auth/models"
	base "api-tests-template/tests"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	owner  models.LoginOkResponse
	reader models.LoginOkResponse
}

func TestSuiteRun(t *testing.T) { suite.Run(t, &TestSuite{}) }

func (s *TestSuite) SetupSuite() {
	base.SetupSuite()
	base.Precondition("Авторизация владельца объявлений и отдельного пользователя избранного")
	s.owner = auth.Login(s.T(), os.Getenv("TEST_LOGIN"), os.Getenv("TEST_PASSWORD"))
	s.reader = auth.Login(s.T(), os.Getenv("FAVORITES_LOGIN"), os.Getenv("FAVORITES_PASSWORD"))
	require.NotEqual(s.T(), s.owner.User.Id, s.reader.User.Id, "Для избранного нужны два разных аккаунта")
}

func (s *TestSuite) TearDownSuite() { base.TearDownSuite() }
