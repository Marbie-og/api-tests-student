package advertisementcreate

import (
	"api-tests-template/internal/fixtures"
	"strings"
)

func (s *TestSuite) TestCreateValidBoundaries() {
	s.Run("минимальная цена и количество", func() {
		request := fixtures.AdvertisementRequest()
		request.Price = 0
		request.Quantity = 1
		s.checkCreatedAdvertisement(request)
	})
	s.Run("максимальные значения и три фотографии", func() {
		request := fixtures.AdvertisementRequest()
		request.Title += strings.Repeat("я", 50-len([]rune(request.Title)))
		request.Description = strings.Repeat("я", 500)
		request.Price = 1000000
		request.Quantity = 100
		request.Photos = []string{fixtures.PhotoPath, fixtures.PhotoPath, fixtures.PhotoPath}
		s.checkCreatedAdvertisement(request)
	})
}
