package filmdto

type FilmRequest struct {
	ID            int    `json:"id" gorm:"primary_key:auto_increment"`
	Title         string `json:"title" form:"title"  gorm:"type: varchar(255)"`
	Thumbnailfilm string `json:"thumbnailfilm" form:"image" gorm:"type: varchar(255)"`
	Year          string `json:"year" form:"year" gorm:"type: text" validate:"required" `
	CategoryID    int    `json:"category_id" form:"category_id" gorm:"type: int"`
	Description   string `json:"description" form:"description" gorm:"type: varchar(255)"`
	LinkFilm      string `json:"linkfilm"  gorm:"type: text" form:"linkfilm"`
}

type UpdateFilmRequest struct {
	Title         string `json:"title" form:"title"`
	Thumbnailfilm string `json:"image" form:"image"`
	Year          string `json:"year" form:"year"`
	CategoryID    int    `json:"category_id" form:"category_id" gorm:"type: int"`
	Description   string `json:"description" form:"description"`
	LinkFilm      string `json:"linkfilm"  gorm:"type: text" form:"linkfilm"`
}
