package favorites

import (
	"api-tests-template/internal/assertions"
	"api-tests-template/internal/fixtures"
	"api-tests-template/internal/managers/favorites"
	"net/http"
)

func (s *TestSuite) TestUnauthorized() {
	t := s.T()
	ad := fixtures.CreateAdvertisement(t, s.owner.Token, fixtures.AdvertisementRequest())
	favorites.Add(t, s.reader.Token, ad.ID)
	cases := []struct{ name, token, message string }{
		{"без токена", "", "Authorization header required"},
		{"неверный токен", "incorrect_token", "Invalid or expired token"},
	}
	for _, tc := range cases {
		s.Run(tc.name, func() {
			s.Run("GET", func() {
				assertions.APIError(s.T(), favorites.ListError(s.T(), tc.token, http.StatusUnauthorized), "unauthorized", tc.message)
			})
			s.Run("POST", func() {
				assertions.APIError(s.T(), favorites.AddError(s.T(), tc.token, ad.ID, http.StatusUnauthorized), "unauthorized", tc.message)
			})
			s.Run("DELETE", func() {
				assertions.APIError(s.T(), favorites.DeleteError(s.T(), tc.token, ad.ID, http.StatusUnauthorized), "unauthorized", tc.message)
			})
		})
	}
	assertions.FavoritePresent(t, favorites.List(t, s.reader.Token), ad.ID, true)
}
