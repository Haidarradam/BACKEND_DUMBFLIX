package filmdto

import "dumbflix/models"

type FilmResponse struct {
	ID            int                     `json:"id" gorm:"primary_key:auto_increment"`
	Title         string                  `json:"title" from:"title"  gorm:"type: varchar(255)"`
	Thumbnailfilm string                  `json:"thumbnailfilm" gorm:"type:varchar(255)"`
	Year          string                  `json:"year" form:"year" gorm:"type: text"`
	Category      models.CategoryResponse `json:"category"`
	CategoryID    int                     `json:"-"`
	Description   string                  `json:"description" gorm:"type: varchar(255)"`
	LinkFilm      string                  `json:"linkfilm"  gorm:"type: text" form:"linkfilm"`
}
