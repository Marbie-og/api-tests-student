package advertisementcreate

import (
	"time"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/assertions"
	"api-tests-template/internal/fixtures"
	"api-tests-template/internal/managers/advertisements"
	"api-tests-template/internal/managers/advertisements/models"
)

func (s *TestSuite) TestCreateWithAllFields() {
	s.checkCreatedAdvertisement(fixtures.AdvertisementRequest())
}

// checkCreatedAdvertisement проверяет карточку, поиск и список фотографий после создания.
func (s *TestSuite) checkCreatedAdvertisement(request models.CreateRequest) {
	token := s.login.Token
	created := fixtures.CreateAdvertisement(s.T(), token, request)

	s.Run("Читаем карточку и проверяем сохранённые поля", func() {
		saved := advertisements.Get(s.T(), created.ID)
		assertions.AdvertisementFields(s.T(), saved, request, created.ID, s.login.User.Id)
		require.Equal(s.T(), s.login.User.Id, saved.User.ID, "Неверный пользователь в карточке")
	})

	s.Run("Находим объявление по полному названию", func() {
		deadline := time.Now().Add(10 * time.Second)
		for time.Now().Before(deadline) {
			list := advertisements.GetAdvertisementsList(s.T(), token, request.Title)
			if assertions.ContainsAdvertisement(list.Items, created.ID) {
				return
			}
			time.Sleep(200 * time.Millisecond)
		}
		s.T().Fatalf("Объявление %s не появилось в поиске по %q за 10 секунд", created.ID, request.Title)
	})

	s.Run("Проверяем загруженные фотографии через API", func() {
		photos := advertisements.Photos(s.T(), token, created.ID)
		require.Len(s.T(), photos, len(request.Photos), "Неверное число загруженных фотографий")
		require.Len(s.T(), created.Photos, len(photos), "Число фотографий в ответах создания и чтения различается")
		for i, photo := range photos {
			require.NotEmpty(s.T(), photo.ID, "Нет ID фотографии")
			require.Equal(s.T(), created.Photos[i].ID, photo.ID, "Получена другая фотография")
			require.NotEmpty(s.T(), photo.URL, "Нет URL загруженной фотографии")
			require.Equal(s.T(), created.Photos[i].URL, photo.URL, "URL фотографии изменился")
			require.Equal(s.T(), i, photo.SortOrder, "Неверный порядок фотографий")
			require.NotEmpty(s.T(), photo.CreatedAt, "Нет даты загрузки фотографии")
		}
	})
}
