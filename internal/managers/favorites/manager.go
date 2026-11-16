package favorites

import (
	client "api-tests-template/internal/client/http/favorites"
	"api-tests-template/internal/managers/favorites/models"
	"api-tests-template/internal/managers/response"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

// Add добавляет объявление и проверяет контракт успешного ответа.
func Add(t *testing.T, token, id string) models.Favorite {
	t.Helper()
	var result models.Favorite
	response.Decode(t, client.HttpAdd(t, token, id), http.StatusCreated, &result)
	require.Equal(t, id, result.Advertisement.ID, "В избранное добавлено другое объявление")
	require.NotEmpty(t, result.FavoriteID, "Нет ID записи избранного")
	require.NotEmpty(t, result.AddedAt, "Нет даты добавления")
	return result
}

// List возвращает избранное текущего пользователя.
func List(t *testing.T, token string) models.ListResponse {
	t.Helper()
	var result models.ListResponse
	response.Decode(t, client.HttpList(t, token), http.StatusOK, &result)
	return result
}

// Delete удаляет запись из избранного.
func Delete(t *testing.T, token, id string) {
	t.Helper()
	require.Empty(t, response.Read(t, client.HttpDelete(t, token, id), http.StatusNoContent), "Ответ 204 должен быть пустым")
}

// AddError возвращает ошибку добавления.
func AddError(t *testing.T, token, id string, status int) response.Error {
	t.Helper()
	return response.ReadError(t, client.HttpAdd(t, token, id), status)
}

// DeleteError возвращает ошибку удаления.
func DeleteError(t *testing.T, token, id string, status int) response.Error {
	t.Helper()
	return response.ReadError(t, client.HttpDelete(t, token, id), status)
}

// ListError возвращает ошибку чтения избранного.
func ListError(t *testing.T, token string, status int) response.Error {
	t.Helper()
	return response.ReadError(t, client.HttpList(t, token), status)
}
