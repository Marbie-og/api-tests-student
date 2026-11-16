package favorites

import (
	"api-tests-template/internal/assertions"
	"api-tests-template/internal/managers/favorites"
	"net/http"
)

func (s *TestSuite) TestMissingAdvertisement() {
	// Валидный UUID, которому не соответствует объявление на стенде.
	const id = "e2f8994a-3341-4d67-bfc1-aad4b81efcfe"
	s.Run("добавление несуществующего объявления", func() {
		result := favorites.AddError(s.T(), s.reader.Token, id, http.StatusNotFound)
		assertions.APIError(s.T(), result, "not_found", "Advertisement not found")
	})
	s.Run("удаление несуществующей записи", func() {
		result := favorites.DeleteError(s.T(), s.reader.Token, id, http.StatusNotFound)
		assertions.APIError(s.T(), result, "not_found", "Favorite entry not found")
	})
}
