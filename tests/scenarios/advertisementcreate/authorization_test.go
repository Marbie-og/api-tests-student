package advertisementcreate

import (
	"api-tests-template/internal/assertions"
	"api-tests-template/internal/fixtures"
	"api-tests-template/internal/managers/advertisements"
	"net/http"
)

func (s *TestSuite) TestCreateUnauthorized() {
	cases := []struct{ name, token, message string }{
		{"без токена", "", "Authorization header required"},
		{"неверный токен", "incorrect_token", "Invalid or expired token"},
	}
	for _, tc := range cases {
		s.Run(tc.name, func() {
			s.T().Log("Шаг: создание без действительной авторизации")
			result := advertisements.CreateError(s.T(), tc.token, s.login.Token, fixtures.AdvertisementRequest(), http.StatusUnauthorized)
			assertions.APIError(s.T(), result, "unauthorized", tc.message)
		})
	}
}
