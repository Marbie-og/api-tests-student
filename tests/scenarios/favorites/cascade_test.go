package favorites

import (
	"api-tests-template/internal/assertions"
	"api-tests-template/internal/fixtures"
	"api-tests-template/internal/managers/advertisements"
	"api-tests-template/internal/managers/favorites"

	"github.com/stretchr/testify/require"
)

func (s *TestSuite) TestDeleteAdvertisementRemovesFavorite() {
	t := s.T()
	ad := fixtures.CreateAdvertisement(t, s.owner.Token, fixtures.AdvertisementRequest())
	before := favorites.List(t, s.reader.Token)
	favorites.Add(t, s.reader.Token, ad.ID)
	assertions.FavoritePresent(t, favorites.List(t, s.reader.Token), ad.ID, true)
	t.Log("Шаг: владелец удаляет само объявление")
	advertisements.Delete(t, s.owner.Token, ad.ID)
	after := favorites.List(t, s.reader.Token)
	assertions.FavoritePresent(t, after, ad.ID, false)
	require.Equal(t, before.Total, after.Total, "Каскадное удаление не обновило счётчик")
	t.Log("Результат: удалённое объявление исчезло из избранного")
}
