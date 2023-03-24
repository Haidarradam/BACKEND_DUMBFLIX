package models

type Film struct {
	ID            int              `json:"id" gorm:"primary_key:auto_increment"`
	Title         string           `json:"title" gorm:"type: varchar(255)"`
	Thumbnailfilm string           `json:"image" gorm:"type:varchar(255)"`
	Year          string           `json:"year" gorm:"type: varchar(255)" `
	Category      CategoryResponse `json:"category"`
	CategoryID    int              `json:"-"`
	Description   string           `json:"description" gorm:"type: varchar(255)"`
	LinkFilm      string           `json:"linkfilm"  gorm:"type: text" form:"linkfilm"`
}

type FilmResponse struct {
	ID            int              `json:"id"`
	Title         string           `json:"title"`
	Thumbnailfilm string           `json:"thumbnailfilm"`
	Year          string           `json:"year"`
	Category      CategoryResponse `json:"category"`
	CategoryID    int              `json:"-"`
	Description   string           `json:"description"`
	LinkFilm      string           `json:"linkfilm"`
}

type FilmInCategory struct {
	ID            int              `json:"id"`
	Title         string           `json:"title"`
	Thumbnailfilm string           `json:"thumbnailfilm"`
	Year          string           `json:"year"`
	Category      CategoryResponse `json:"category"`
	CategoryID    int              `json:"-"`
	Description   string           `json:"description"`
	LinkFilm      string           `json:"linkfilm"`
}

func (FilmResponse) TableName() string {
	return "films"
}

func (FilmInCategory) TableName() string {
	return "films"
}
