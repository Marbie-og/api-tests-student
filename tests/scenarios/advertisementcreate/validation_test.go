package advertisementcreate

import (
	"api-tests-template/internal/assertions"
	"api-tests-template/internal/fixtures"
	"api-tests-template/internal/managers/advertisements"
	"api-tests-template/internal/managers/advertisements/models"
	"net/http"
	"strings"
)

func (s *TestSuite) TestCreateValidation() {
	cases := []struct {
		name    string
		change  func(*models.CreateRequest)
		message string
	}{
		{"нет title", func(r *models.CreateRequest) { r.OmitField = "title" }, "title is required"},
		{"пустой title", func(r *models.CreateRequest) { r.Title = "" }, "title is required"},
		{"title 51 символ", func(r *models.CreateRequest) { r.Title = strings.Repeat("я", 51) }, "title must not exceed 50 characters provided 51"},
		{"нет description", func(r *models.CreateRequest) { r.OmitField = "description" }, "description is required"},
		{"пустой description", func(r *models.CreateRequest) { r.Description = "" }, "description is required"},
		{"description 501 символ", func(r *models.CreateRequest) { r.Description = strings.Repeat("я", 501) }, "description must be between 1 and 500 characters"},
		{"цена -1", func(r *models.CreateRequest) { r.Price = -1 }, "price must be between 0 and 1000000 inclusive"},
		{"цена 1000001", func(r *models.CreateRequest) { r.Price = 1000001 }, "price must be between 0 and 1000000 inclusive"},
		{"количество 0", func(r *models.CreateRequest) { r.Quantity = 0 }, "quantity must be between 1 and 100 inclusive"},
		{"количество -1", func(r *models.CreateRequest) { r.Quantity = -1 }, "quantity must be between 1 and 100 inclusive"},
		{"количество 101", func(r *models.CreateRequest) { r.Quantity = 101 }, "quantity must be between 1 and 100 inclusive"},
		{"нет фотографий", func(r *models.CreateRequest) { r.Photos = nil }, "at least 1 photo is required"},
		{"четыре фотографии", func(r *models.CreateRequest) {
			r.Photos = []string{fixtures.PhotoPath, fixtures.PhotoPath, fixtures.PhotoPath, fixtures.PhotoPath}
		}, "maximum 3 photos allowed"},
		{"текст вместо фотографии", func(r *models.CreateRequest) { r.Photos = []string{"../../../testdata/not-image.txt"} }, "photo 1: file type not allowed: extension \".txt\""},
	}
	for _, tc := range cases {
		s.Run(tc.name, func() {
			request := fixtures.AdvertisementRequest()
			tc.change(&request)
			s.T().Log("Шаг: создание с невалидными параметрами")
			result := advertisements.CreateError(s.T(), s.login.Token, s.login.Token, request, http.StatusBadRequest)
			assertions.APIError(s.T(), result, "invalid_request", tc.message)
		})
	}
}
