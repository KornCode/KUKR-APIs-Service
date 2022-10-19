package database

import "gorm.io/gorm"

type publishMySQL struct {
	gorm.Model
	Category    int `gorm:"index:idx_name,unique; not null"`
	Text        string
	Bibid       string `gorm:"index:idx_name,unique; size:10; not null"`
	ContentType string `gorm:"size:100"`
	TitleTh     string
	TitleEn     string
	PubYear     int `gorm:"index:idx_name,unique; not null"`
	AuthorTh    string
	AuthorEn    string
	Link        string
}

func (p publishMySQL) TableName() string {
	return "publishes"
}
