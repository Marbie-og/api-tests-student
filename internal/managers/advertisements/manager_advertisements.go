package advertisements

import (
	"encoding/json"
	"io"
	"testing"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/client/http/advertisements"
	"api-tests-template/internal/managers/advertisements/models"
)

// GetAdvertisementsList возвращает типизированный результат поиска.
func GetAdvertisementsList(t *testing.T, token, text string) models.AdvertisementsList {
	t.Helper()
	raw := advertisements.HttpGetAdvertisements(t, token, text)
	defer raw.Body.Close()
	body, err := io.ReadAll(raw.Body)
	require.NoError(t, err, "Не удалось прочитать ответ поиска")
	var resp models.AdvertisementsList
	require.NoError(t, json.Unmarshal(body, &resp), "Не удалось разобрать ответ поиска")
	return resp
}
