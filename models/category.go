package models

type Category struct {
	ID   int              `json:"id"`
	Name string           `json:"name" form:"name" validate:"required"`
	Film []FilmInCategory `json:"film"`
}

type CategoryResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name" form:"name" validate:"required"`
}

func (CategoryResponse) TableName() string {
	return "categories"
}
