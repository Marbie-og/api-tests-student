package response

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

type Error struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

// Read проверяет HTTP-статус, читает и закрывает тело ответа.
func Read(t *testing.T, raw *http.Response, status int) []byte {
	t.Helper()
	require.NotNil(t, raw, "Сервер не вернул HTTP-ответ")
	body, err := io.ReadAll(raw.Body)
	closeErr := raw.Body.Close()
	require.NoError(t, err, "Не удалось прочитать ответ")
	require.NoError(t, closeErr, "Не удалось закрыть ответ")
	require.Equal(t, status, raw.StatusCode, "Неожиданный HTTP-статус: %s", body)
	return body
}

// Decode разбирает JSON в переданную модель.
func Decode(t *testing.T, raw *http.Response, status int, result any) {
	t.Helper()
	body := Read(t, raw, status)
	require.NoError(t, json.Unmarshal(body, result), "Ответ не соответствует модели: %s", body)
}

// ReadError возвращает типизированную ошибку API.
func ReadError(t *testing.T, raw *http.Response, status int) Error {
	t.Helper()
	var result Error
	Decode(t, raw, status, &result)
	return result
}
