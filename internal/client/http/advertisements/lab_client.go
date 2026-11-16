package advertisements

import (
	"api-tests-template/internal/constants/path"
	apiRunner "api-tests-template/internal/helpers/api-runner"
	"api-tests-template/internal/managers/advertisements/models"
	"bytes"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

// HttpPostAdvertisement собирает multipart-форму; OmitField нужен для проверок отсутствующего поля.
func HttpPostAdvertisement(t *testing.T, token string, request models.CreateRequest) *http.Response {
	t.Helper()
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	fields := []struct{ name, value string }{
		{"title", request.Title},
		{"description", request.Description},
		{"price", strconv.FormatFloat(request.Price, 'f', -1, 64)},
		{"quantity", strconv.Itoa(request.Quantity)},
	}
	for _, field := range fields {
		if field.name != request.OmitField {
			require.NoError(t, writer.WriteField(field.name, field.value), "Не удалось записать %s", field.name)
		}
	}
	for _, photoPath := range request.Photos {
		photo, err := os.ReadFile(filepath.Clean(photoPath))
		require.NoError(t, err, "Не удалось прочитать %s", photoPath)
		part, err := writer.CreateFormFile("photos", filepath.Base(photoPath))
		require.NoError(t, err, "Не удалось создать часть photos")
		_, err = part.Write(photo)
		require.NoError(t, err, "Не удалось записать фотографию")
	}
	require.NoError(t, writer.Close(), "Не удалось завершить multipart-форму")
	return apiRunner.GetRunner().Auth(token).Create().Post(path.AdvertisementPath).
		ContentType(writer.FormDataContentType()).Body(body.String()).Expect(t).End().Response
}

// HttpGetAdvertisement получает объявление по id
func HttpGetAdvertisement(t *testing.T, id string) *http.Response {
	t.Helper()
	return apiRunner.GetRunner().Create().Get(path.AdvertisementPath).Query("id", id).Expect(t).End().Response
}

// HttpDeleteAdvertisement удаляет объявление
func HttpDeleteAdvertisement(t *testing.T, token, id string) *http.Response {
	t.Helper()
	return apiRunner.GetRunner().Auth(token).Create().Delete(path.AdvertisementPath).Query("id", id).Expect(t).End().Response
}

// HttpGetPhotos запрашивает фотографии объявления
func HttpGetPhotos(t *testing.T, token, id string) *http.Response {
	t.Helper()
	return apiRunner.GetRunner().Auth(token).Create().Get(fmt.Sprintf(path.AdvertisementPhotosPath, id)).Expect(t).End().Response
}
