package advertisements

import (
	client "api-tests-template/internal/client/http/advertisements"
	"api-tests-template/internal/managers/advertisements/models"
	"api-tests-template/internal/managers/response"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

// Create создаёт объявление и возвращает его поля.
func Create(t *testing.T, token string, request models.CreateRequest) models.Advertisement {
	t.Helper()
	var result models.Advertisement
	response.Decode(t, client.HttpPostAdvertisement(t, token, request), http.StatusCreated, &result)
	require.NotEmpty(t, result.ID, "Создание не вернуло ID объявления")
	return result
}

// CreateError проверяет отказ в создании. При неожиданном успехе заранее регистрирует удаление обьявления
func CreateError(t *testing.T, token, ownerToken string, request models.CreateRequest, status int) response.Error {
	t.Helper()
	raw := client.HttpPostAdvertisement(t, token, request)
	if raw.StatusCode == http.StatusCreated {
		var created models.Advertisement
		response.Decode(t, raw, http.StatusCreated, &created)
		if created.ID != "" {
			t.Cleanup(func() { Cleanup(t, ownerToken, created.ID) })
		}
		t.Fatalf("Ожидался отказ HTTP %d, но объявление было создано", status)
	}
	return response.ReadError(t, raw, status)
}

// Get читает сохранённое объявление
func Get(t *testing.T, id string) models.Advertisement {
	t.Helper()
	var result models.Advertisement
	response.Decode(t, client.HttpGetAdvertisement(t, id), http.StatusOK, &result)
	return result
}

// Delete проверяет успешное удаление и пустое тело ответа 204 StatusNoContent
func Delete(t *testing.T, token, id string) {
	t.Helper()
	body := response.Read(t, client.HttpDeleteAdvertisement(t, token, id), http.StatusNoContent)
	require.Empty(t, body, "Ответ 204 должен быть пустым")
}

// Cleanup удаляет данные теста; при повторном удалении проверяет отсутствие через GET.
func Cleanup(t *testing.T, token, id string) {
	t.Helper()
	raw := client.HttpDeleteAdvertisement(t, token, id)
	status := raw.StatusCode
	body := response.Read(t, raw, status)
	if status == http.StatusNoContent {
		return
	}
	if status == http.StatusForbidden || status == http.StatusNotFound {
		check := client.HttpGetAdvertisement(t, id)
		checkStatus := check.StatusCode
		response.Read(t, check, checkStatus)
		if checkStatus == http.StatusNotFound {
			return
		}
	}
	t.Errorf("Не удалось удалить тестовое объявление %s: HTTP %d, %s", id, status, body)
}

// Photos возвращает список загруженных фотографий
func Photos(t *testing.T, token, id string) []models.Photo {
	t.Helper()
	var result []models.Photo
	response.Decode(t, client.HttpGetPhotos(t, token, id), http.StatusOK, &result)
	return result
}
