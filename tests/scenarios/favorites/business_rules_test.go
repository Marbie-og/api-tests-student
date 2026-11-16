package favorites

import (
	"api-tests-template/internal/assertions"
	"api-tests-template/internal/fixtures"
	"api-tests-template/internal/managers/favorites"
	"net/http"

	"github.com/stretchr/testify/require"
)

func (s *TestSuite) TestOwnAdvertisement() {
	t := s.T()
	ad := fixtures.CreateAdvertisement(t, s.owner.Token, fixtures.AdvertisementRequest())
	result := favorites.AddError(t, s.owner.Token, ad.ID, http.StatusForbidden)
	assertions.APIError(t, result, "forbidden", "Cannot add own advertisement to favorites")
	assertions.FavoritePresent(t, favorites.List(t, s.owner.Token), ad.ID, false)
}

func (s *TestSuite) TestDuplicateFavorite() {
	t := s.T()
	ad := fixtures.CreateAdvertisement(t, s.owner.Token, fixtures.AdvertisementRequest())
	favorites.Add(t, s.reader.Token, ad.ID)
	before := favorites.List(t, s.reader.Token)
	result := favorites.AddError(t, s.reader.Token, ad.ID, http.StatusConflict)
	assertions.APIError(t, result, "already_exists", "Advertisement is already in favorites")
	after := favorites.List(t, s.reader.Token)
	require.Equal(t, before.Total, after.Total, "Повторное добавление изменило счётчик")
	count := 0
	for _, item := range after.Items {
		if item.Advertisement.ID == ad.ID {
			count++
		}
	}
	require.Equal(t, 1, count, "В избранном появились дубликаты")
}
