package publishhand

import (
	"github.com/go-playground/validator/v10"
)

type errorResponse struct {
	Field string
	Tag   string
	Value string
}

var validate = validator.New()

func validateStruct(tovalidate interface{}) []*errorResponse {
	var errs []*errorResponse
	if err := validate.Struct(tovalidate); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			element := errorResponse{
				Field: err.Field(),
				Tag:   err.Tag(),
				Value: err.Param(),
			}

			errs = append(errs, &element)
		}
	}

	return errs
}

type jsonPublishCreateOne struct {
	Category    int    `json:"category" validate:"required,numeric"`
	Text        string `json:"text" validate:"required"`
	Bibid       string `json:"bibid" validate:"required,alphanum,max=10"`
	ContentType string `json:"content_type" validate:"required,max=100"`
	TitleTh     string `json:"title_th" validate:"required"`
	TitleEn     string `json:"title_en" validate:"required"`
	PubYear     int    `json:"pub_year" validate:"required,numeric"`
	AuthorTh    string `json:"author_th" validate:"required"`
	AuthorEn    string `json:"author_en" validate:"required"`
	Link        string `json:"link" validate:"required"`
}

type jsonPublishUpdateOneByPK struct {
	ID          uint   `json:"id" validate:"required,numeric"`
	Category    int    `json:"category" validate:"required,numeric"`
	Text        string `json:"text" validate:"required"`
	Bibid       string `json:"bibid" validate:"required,alphanum,max=10"`
	ContentType string `json:"content_type" validate:"required,max=100"`
	TitleTh     string `json:"title_th" validate:"required"`
	TitleEn     string `json:"title_en" validate:"required"`
	PubYear     int    `json:"pub_year" validate:"required,numeric"`
	AuthorTh    string `json:"author_th" validate:"required"`
	AuthorEn    string `json:"author_en" validate:"required"`
	Link        string `json:"link" validate:"required"`
}

type jsonPublishDeleteOneByPK struct {
	ID uint `json:"id" validate:"required,numeric"`
}

type queryPublishCategoryAndPubYear struct {
	Category int `query:"category" validate:"required,numeric"`
	PubYear  int `query:"pub_year" validate:"required,numeric"`
}

type queryPublishBibid struct {
	Bibid string `query:"bibid" validate:"alphanum"`
}

type queryPublishPaginateByOptions struct {
	Page     int `query:"page" validate:"required,numeric"`
	Limit    int `query:"limit" validate:"required,numeric"`
	Category int `query:"category" validate:"required,numeric"`
	PubYear  int `query:"pub_year" validate:"required,numeric"`
}

type jsonPublishPubYear struct {
	PubYear int `json:"pub_year" validate:"required,numeric"`
}
