package service

import (
	"strings"
	"sync"

	"github.com/KornCode/KUKR-APIs-Service/app/errs"
	repository "github.com/KornCode/KUKR-APIs-Service/app/repositories"
	"github.com/KornCode/KUKR-APIs-Service/pkg/logs"
)

type publishService struct {
	publishRepository repository.PublishRepository
}

func NewPublishService(publishRepository repository.PublishRepository) publishService {
	return publishService{publishRepository}
}

func (s publishService) CreateOne(publish *Publish) (uint, error) {
	repo_row := repository.Publish{
		Category:    publish.Category,
		Text:        publish.Text,
		Bibid:       publish.Bibid,
		ContentType: publish.ContentType,
		TitleTh:     publish.TitleTh,
		TitleEn:     publish.TitleEn,
		PubYear:     publish.PubYear,
		AuthorTh:    publish.AuthorTh,
		AuthorEn:    publish.AuthorEn,
		Link:        publish.Link,
	}

	repo_result, err := s.publishRepository.CreateOne(repo_row)
	if err != nil {
		if strings.Contains(err.Error(), "1062") {
			return 0, errs.Conflict(err.Error())
		}

		return 0, errs.UnexpectedError()
	}
	if repo_result.RowsAffected == 0 {
		return 0, errs.NotFound("id: not found")
	}

	return repo_result.ID, nil
}

func (s publishService) UpdateOneByPK(pk uint, publish *Publish) error {
	repo_row := repository.Publish{
		Category:    publish.Category,
		Text:        publish.Text,
		Bibid:       publish.Bibid,
		ContentType: publish.ContentType,
		TitleTh:     publish.TitleTh,
		TitleEn:     publish.TitleEn,
		PubYear:     publish.PubYear,
		AuthorTh:    publish.AuthorTh,
		AuthorEn:    publish.AuthorEn,
		Link:        publish.Link,
	}

	repo_result, err := s.publishRepository.UpdateOneByID(pk, &repo_row)
	if err != nil {
		return errs.UnexpectedError()
	}
	if repo_result.RowsAffected == 0 {
		return errs.NotFound("id: not found")
	}

	return nil
}

func (s publishService) DeleteOneByPK(pk uint) error {
	repo_result, err := s.publishRepository.DeleteOneByID(pk)
	if err != nil {
		return errs.UnexpectedError()
	}
	if repo_result.RowsAffected == 0 {
		return errs.NotFound("id: not found")
	}

	return nil
}

func (s publishService) GetByCategoryAndPubYear(category, pubyear int) ([]Publish, error) {
	repo_results, err := s.publishRepository.GetByCategoryAndPubYear(category, pubyear)
	if err != nil {
		logs.Error(err)

		return nil, errs.UnexpectedError()
	}
	if len(repo_results) == 0 {
		return nil, errs.NotFound("query empty")
	}

	serv_results := []Publish{}
	for _, repo_result := range repo_results {
		serv_results = append(serv_results, Publish{
			ID:          repo_result.ID,
			Category:    repo_result.Category,
			Text:        repo_result.Text,
			Bibid:       repo_result.Bibid,
			ContentType: repo_result.ContentType,
			TitleTh:     repo_result.TitleTh,
			TitleEn:     repo_result.TitleEn,
			PubYear:     repo_result.PubYear,
			AuthorTh:    repo_result.AuthorTh,
			AuthorEn:    repo_result.AuthorEn,
			Link:        repo_result.Link,
			CreatedAt:   repo_result.CreatedAt,
			UpdatedAt:   repo_result.UpdatedAt,
		})
	}

	return serv_results, nil
}

func (s publishService) GetByBibid(bibid string) (*Publish, error) {
	repo_result, err := s.publishRepository.GetByBibid(bibid)
	if err != nil {
		logs.Error(err)
		if err.Error() == "record not found" {
			return nil, errs.NotFound("query empty")
		}

		return nil, errs.UnexpectedError()
	}

	serv_result := Publish{
		ID:          repo_result.ID,
		Category:    repo_result.Category,
		Text:        repo_result.Text,
		Bibid:       repo_result.Bibid,
		ContentType: repo_result.ContentType,
		TitleTh:     repo_result.TitleTh,
		TitleEn:     repo_result.TitleEn,
		PubYear:     repo_result.PubYear,
		AuthorTh:    repo_result.AuthorTh,
		AuthorEn:    repo_result.AuthorEn,
		Link:        repo_result.Link,
		CreatedAt:   repo_result.CreatedAt,
		UpdatedAt:   repo_result.UpdatedAt,
	}

	return &serv_result, nil
}

func (s publishService) GetPaginateByOptions(page, limit int, options map[string]interface{}) (*PublishPaginate, error) {
	repo_result, err := s.publishRepository.GetPaginateByOptions(
		page, limit, map[string]interface{}{
			"category": options["category"],
			"pub_year": options["pub_year"],
		},
	)
	if err != nil {
		logs.Error(err)

		return nil, errs.UnexpectedError()
	}
	if len(repo_result.Rows) == 0 {
		return nil, errs.NotFound("query empty")
	}

	serv_result_rows := []Publish{}
	for _, repo_result_row := range repo_result.Rows {
		serv_result_rows = append(serv_result_rows, Publish{
			ID:          repo_result_row.ID,
			Category:    repo_result_row.Category,
			Text:        repo_result_row.Text,
			Bibid:       repo_result_row.Bibid,
			ContentType: repo_result_row.ContentType,
			TitleTh:     repo_result_row.TitleTh,
			TitleEn:     repo_result_row.TitleEn,
			PubYear:     repo_result_row.PubYear,
			AuthorTh:    repo_result_row.AuthorTh,
			AuthorEn:    repo_result_row.AuthorEn,
			Link:        repo_result_row.Link,
			CreatedAt:   repo_result_row.CreatedAt,
			UpdatedAt:   repo_result_row.UpdatedAt,
		})
	}

	total_pages := calcPaginateTotalPages(repo_result.TotalRows, limit)
	serv_result := PublishPaginate{
		Page:       page,
		Limit:      limit,
		TotalRows:  repo_result.TotalRows,
		TotalPages: total_pages,
		Sort:       repo_result.Sort,
		Rows:       serv_result_rows,
	}

	return &serv_result, nil
}

func (s publishService) SyncDataSource(pub_year int) error {
	repo_rows := []repository.Publish{}

	var wg sync.WaitGroup
	for _, category := range [...]int{1, 2, 3, 4} {
		wg.Add(1)
		go func(category int) {
			defer wg.Done()
			publishSynApiFetchAll(category, pub_year, &repo_rows)
		}(category)
	}
	wg.Wait()

	if len(repo_rows) == 0 {
		return errs.NotFound("fetch empty")
	}

	if err := s.publishRepository.CreateBatchUpdateOnConflict(repo_rows); err != nil {
		logs.Error(err)

		return errs.UnexpectedError()
	}

	return nil
}
