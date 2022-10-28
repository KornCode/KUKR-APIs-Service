package publishserv

import "time"

type Publish struct {
	ID          uint      `json:"id"`
	Category    int       `json:"category"`
	Text        string    `json:"text"`
	Bibid       string    `json:"bibid"`
	ContentType string    `json:"content_type"`
	TitleTh     string    `json:"title_th"`
	TitleEn     string    `json:"title_en"`
	PubYear     int       `json:"pub_year"`
	AuthorTh    string    `json:"author_th"`
	AuthorEn    string    `json:"author_en"`
	Link        string    `json:"link"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type PublishPaginate struct {
	Page       int       `json:"page"`
	Limit      int       `json:"limit"`
	TotalRows  int       `json:"total_rows"`
	TotalPages int       `json:"total_pages"`
	Sort       string    `json:"sort"`
	Rows       []Publish `json:"rows"`
}

type PublishService interface {
	CreateOne(publish *Publish) (uint, error)
	UpdateOneByPK(pk uint, publish *Publish) error
	DeleteOneByPK(pk uint) error

	GetByCategoryAndPubYear(id, pubyear int) ([]Publish, error)
	GetByBibid(bibid string) (*Publish, error)
	GetPaginateByOptions(page, limit int, optionals map[string]interface{}) (*PublishPaginate, error)

	SyncDataSource(pub_year int) error
}
