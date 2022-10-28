package publishserv

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	publishrpt "github.com/KornCode/KUKR-APIs-Service/app/repositories/publish"
)

type publishApiReadBody struct {
	Text  string `json:"text"`
	Year  string `json:"year"`
	Page  string `json:"page"`
	Limit string `json:"limit"`
	Total string `json:"total"`
	Show  string `json:"show"`
	Data  []struct {
		Bibid       string `json:"bibid"`
		ContentType string `json:"content_type"`
		TitleTh     string `json:"title_th"`
		TitleEn     string `json:"title_en"`
		PubYear     string `json:"pubyear"`
		AuthorTh    string `json:"author_th"`
		AuthorEn    string `json:"author_en"`
		Link        string `json:"link"`
	} `json:"data"`
}

func publishSynApiFetchAll(category, pub_year int, publishes *[]publishrpt.Publish) {
	api_uri := os.Getenv("PUBLISH_SOURCE_API_URI")

	var page int = 1
	for {
		url_vals := url.Values{}
		url_vals.Add("id", strconv.Itoa(category))
		url_vals.Add("is", strconv.Itoa(pub_year))
		url_vals.Add("pg", strconv.Itoa(page))

		req, err := http.NewRequest("GET", api_uri, nil)
		if err != nil {
			break
		}

		req.URL.RawQuery = url_vals.Encode()

		response, err := http.Get(req.URL.String())
		if err != nil {
			break
		}
		defer response.Body.Close()

		response_body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			break
		}

		response_read := publishApiReadBody{}
		if err := json.Unmarshal(response_body, &response_read); err != nil {
			break
		}
		if len(response_read.Data) == 0 {
			break
		}

		for _, resp_data := range response_read.Data {
			*publishes = append(*publishes, publishrpt.Publish{
				Category:    category,
				Text:        response_read.Text,
				Bibid:       resp_data.Bibid,
				ContentType: resp_data.ContentType,
				TitleTh:     resp_data.TitleTh,
				TitleEn:     resp_data.TitleEn,
				PubYear:     pub_year,
				AuthorTh:    resp_data.AuthorTh,
				AuthorEn:    resp_data.AuthorEn,
				Link:        resp_data.Link,
			})
		}

		show_pages := strings.Split(response_read.Show, "-")

		end_page, err := strconv.Atoi(show_pages[1])
		if err != nil {
			break
		}
		total_page, err := strconv.Atoi(response_read.Total)
		if err != nil {
			break
		}

		page = page + 1
		if end_page >= total_page {
			break
		}
	}
}
