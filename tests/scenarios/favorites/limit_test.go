package favorites

import (
	"api-tests-template/internal/assertions"
	"api-tests-template/internal/fixtures"
	"api-tests-template/internal/managers/favorites"
	"net/http"

	"github.com/stretchr/testify/require"
)

const favoritesLimit = 50

func (s *TestSuite) TestFavoritesLimit() {
	t := s.T()
	require.Zero(t, favorites.List(t, s.reader.Token).Total,
		"Для проверки лимита нужен отдельный аккаунт FAVORITES_LOGIN с пустым избранным")
	t.Log("Шаг: заполняем избранное до лимита 50 объявлений")
	for i := 0; i < favoritesLimit; i++ {
		ad := fixtures.CreateAdvertisement(t, s.owner.Token, fixtures.AdvertisementRequest())
		favorites.Add(t, s.reader.Token, ad.ID)
	}
	full := favorites.List(t, s.reader.Token)
	require.Equal(t, favoritesLimit, full.Total, "Лимит не достигнут")
	require.Len(t, full.Items, favoritesLimit, "Запрошенная страница должна содержать все 50 записей")
	extra := fixtures.CreateAdvertisement(t, s.owner.Token, fixtures.AdvertisementRequest())
	t.Log("Шаг: пытаемся добавить 51-е объявление")
	result := favorites.AddError(t, s.reader.Token, extra.ID, http.StatusBadRequest)
	assertions.APIError(t, result, "limit_reached", "Favorites limit of 50 reached")
	after := favorites.List(t, s.reader.Token)
	require.Equal(t, favoritesLimit, after.Total, "После отказа изменился счётчик")
	assertions.FavoritePresent(t, after, extra.ID, false)
}
