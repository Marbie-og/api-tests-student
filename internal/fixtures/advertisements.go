package fixtures

import (
	"api-tests-template/internal/managers/advertisements"
	"api-tests-template/internal/managers/advertisements/models"
	"api-tests-template/internal/utils"
	"testing"
)

const PhotoPath = "../../../testdata/advertisement.png"

// AdvertisementRequest готовит уникальные данные со всеми полями формы.
func AdvertisementRequest() models.CreateRequest {
	return models.CreateRequest{
		Title:       "Лаба Go " + utils.RandomString(12),
		Description: "Проверка создания объявления со всеми заполненными полями",
		Price:       12000, Quantity: 7, Photos: []string{PhotoPath},
	}
}

// CreateAdvertisement регистрирует уборку сразу после создания, до проверок сценария.
func CreateAdvertisement(t *testing.T, token string, request models.CreateRequest) models.Advertisement {
	t.Helper()
	t.Logf("Подготовка: создаём объявление %q", request.Title)
	created := advertisements.Create(t, token, request)
	t.Cleanup(func() {
		t.Logf("Уборка: удаляем объявление %s", created.ID)
		advertisements.Cleanup(t, token, created.ID)
	})
	return created
}
