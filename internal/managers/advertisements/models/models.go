package models

// CreateRequest содержит поля multipart-формы и пути к файлам фотографий.
type CreateRequest struct {
	Title       string
	Description string
	Price       float64
	Quantity    int
	Photos      []string
	OmitField   string
}
