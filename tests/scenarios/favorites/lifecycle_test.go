package favorites

import (
	"api-tests-template/internal/assertions"
	"api-tests-template/internal/fixtures"
	"api-tests-template/internal/managers/advertisements"
	"api-tests-template/internal/managers/favorites"

	"github.com/stretchr/testify/require"
)

func (s *TestSuite) TestAddAndRemove() {
	t := s.T()
	ad := fixtures.CreateAdvertisement(t, s.owner.Token, fixtures.AdvertisementRequest())
	before := favorites.List(t, s.reader.Token)
	assertions.FavoritePresent(t, before, ad.ID, false)
	t.Log("Шаг: добавляем чужое объявление в избранное")
	favorites.Add(t, s.reader.Token, ad.ID)
	after := favorites.List(t, s.reader.Token)
	assertions.FavoritePresent(t, after, ad.ID, true)
	require.Equal(t, before.Total+1, after.Total, "Счётчик избранного не увеличился")
	t.Log("Шаг: удаляем запись из избранного")
	favorites.Delete(t, s.reader.Token, ad.ID)
	after = favorites.List(t, s.reader.Token)
	assertions.FavoritePresent(t, after, ad.ID, false)
	require.Equal(t, before.Total, after.Total, "Счётчик избранного не восстановился")
	require.Equal(t, ad.ID, advertisements.Get(t, ad.ID).ID, "Удаление из избранного не должно удалять объявление")
	t.Log("Результат: полный цикл избранного пройден")
}
