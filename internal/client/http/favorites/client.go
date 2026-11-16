package favorites

import (
	"api-tests-template/internal/constants/path"
	apiRunner "api-tests-template/internal/helpers/api-runner"
	"api-tests-template/internal/managers/favorites/models"
	"fmt"
	"net/http"
	"testing"
)

// HttpAdd добавляет объявление в избранное.
func HttpAdd(t *testing.T, token, id string) *http.Response {
	t.Helper()
	return apiRunner.GetRunner().Auth(token).Create().Post(path.FavoritesPath).
		ContentType("application/json").JSON(models.AddRequest{AdvertisementID: id}).Expect(t).End().Response
}

// HttpList получает до 50 записей избранного, это максимальная вместимость списка.
func HttpList(t *testing.T, token string) *http.Response {
	t.Helper()
	return apiRunner.GetRunner().Auth(token).Create().Get(path.FavoritesPath).Query("page_size", "50").Expect(t).End().Response
}

// HttpDelete удаляет запись по ID объявления, а не по favorite_id.
func HttpDelete(t *testing.T, token, id string) *http.Response {
	t.Helper()
	return apiRunner.GetRunner().Auth(token).Create().Delete(fmt.Sprintf(path.FavoritePath, id)).Expect(t).End().Response
}
