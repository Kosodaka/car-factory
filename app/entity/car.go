package entity

type Car struct {
	Brand string `json:"brand" db:"brand"`
	Color string `json:"color" db:"color"`
	Form  string `json:"form" db:"form"`
}
