package assertions

import (
	ads "api-tests-template/internal/managers/advertisements/models"
	favs "api-tests-template/internal/managers/favorites/models"
	"api-tests-template/internal/managers/response"
	"testing"

	"github.com/stretchr/testify/require"
)

// AdvertisementFields сравнивает сохранённое объявление с отправленными данными.
func AdvertisementFields(t *testing.T, actual ads.Advertisement, expected ads.CreateRequest, id, userID string) {
	t.Helper()
	require.Equal(t, id, actual.ID, "ID объявления не совпадает")
	require.Equal(t, expected.Title, actual.Title, "Название не сохранилось")
	require.Equal(t, expected.Description, actual.Description, "Описание не сохранилось")
	require.Equal(t, expected.Price, actual.Price, "Цена не сохранилась")
	require.Equal(t, expected.Quantity, actual.Quantity, "Количество не сохранилось")
	require.Equal(t, userID, actual.UserID, "Неверный владелец")
	require.NotEmpty(t, actual.CreatedAt, "Нет даты создания")
	require.NotEmpty(t, actual.UpdatedAt, "Нет даты обновления")
	require.Len(t, actual.Photos, len(expected.Photos), "Сохранились не все фотографии")
}

// APIError сравнивает код и текст ошибки
func APIError(t *testing.T, actual response.Error, code, message string) {
	t.Helper()
	require.Equal(t, code, actual.Error, "Неверный код ошибки API")
	require.Equal(t, message, actual.Message, "Неверное сообщение об ошибке")
	t.Logf("Проверено: %s — %s", code, message)
}

// ContainsAdvertisement проверяет выдачу по ID
func ContainsAdvertisement(items []ads.Advertisement, id string) bool {
	for _, item := range items {
		if item.ID == id {
			return true
		}
	}
	return false
}

// FavoritePresent проверяет наличие конкретного объявления в избранном.
func FavoritePresent(t *testing.T, list favs.ListResponse, id string, expected bool) {
	t.Helper()
	found := false
	for _, item := range list.Items {
		if item.Advertisement.ID == id {
			found = true
			break
		}
	}
	require.Equal(t, expected, found, "Наличие объявления %s в избранном не соответствует ожиданию", id)
}
