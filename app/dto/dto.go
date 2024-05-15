package dto

type CarDto struct {
	ID    int    `json:"id" db:"id"`
	Brand string `json:"brand" db:"brand"`
	Color string `json:"color" db:"color"`
	Form  string `json:"form" db:"form"`
}
