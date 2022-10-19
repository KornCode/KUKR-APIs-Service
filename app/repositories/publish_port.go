package repository

import "gorm.io/gorm"

type Publish struct {
	gorm.Model
	Category    int    `gorm:"index:idx_name,unique; not null" json:"category"`
	Text        string `json:"text"`
	Bibid       string `gorm:"index:idx_name,unique; size:10; not null" json:"bibid"`
	ContentType string `gorm:"size:100" json:"content_type"`
	TitleTh     string `json:"title_th"`
	TitleEn     string `json:"tile_en"`
	PubYear     int    `gorm:"index:idx_name,unique; not null" json:"pub_year"`
	AuthorTh    string `json:"author_th"`
	AuthorEn    string `json:"author_en"`
	Link        string `json:"link"`
}

type PublishPaginate struct {
	TotalRows int       `json:"total_rows"`
	Sort      string    `json:"sort"`
	Rows      []Publish `json:"rows"`
}

type PublishCrud struct {
	ID           uint
	RowsAffected int64
}

type PublishRepository interface {
	CreateOne(publish Publish) (*PublishCrud, error)
	UpdateOneByID(id uint, publish *Publish) (*PublishCrud, error)
	DeleteOneByID(id uint) (*PublishCrud, error)

	GetByCategoryAndPubYear(category, pub_year int) ([]Publish, error)
	GetByBibid(bibid string) (*Publish, error)
	GetPaginateByOptions(page, limit int, optionals map[string]interface{}) (*PublishPaginate, error)

	CreateBatchUpdateOnConflict(publishes []Publish) error
}
