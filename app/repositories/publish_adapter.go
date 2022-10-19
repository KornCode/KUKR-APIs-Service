package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v9"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type publishRepositoryDB struct {
	sqlDB   *gorm.DB
	rdCache *redis.Client
}

func NewPublishRepositoryDB(sqlDB *gorm.DB, rdCache *redis.Client) publishRepositoryDB {
	return publishRepositoryDB{sqlDB, rdCache}
}

func (r publishRepositoryDB) CreateOne(publish Publish) (*PublishCrud, error) {
	sql_tx := r.sqlDB.Create(&publish)
	if sql_tx.Error != nil {
		return nil, sql_tx.Error
	}

	sql_result := PublishCrud{
		ID:           publish.ID,
		RowsAffected: sql_tx.RowsAffected,
	}

	return &sql_result, nil
}

func (r publishRepositoryDB) UpdateOneByID(id uint, publish *Publish) (*PublishCrud, error) {
	sql_tx := r.sqlDB.Where("id = ?", id).Updates(&publish)
	if sql_tx.Error != nil {
		return nil, sql_tx.Error
	}

	sql_result := PublishCrud{
		ID:           id,
		RowsAffected: sql_tx.RowsAffected,
	}

	return &sql_result, nil
}

func (r publishRepositoryDB) DeleteOneByID(id uint) (*PublishCrud, error) {
	sql_tx := r.sqlDB.Delete(&Publish{}, id)
	if sql_tx.Error != nil {
		return nil, sql_tx.Error
	}

	sql_result := PublishCrud{
		ID:           id,
		RowsAffected: sql_tx.RowsAffected,
	}

	return &sql_result, nil
}

func (r publishRepositoryDB) GetByCategoryAndPubYear(category, pub_year int) ([]Publish, error) {
	rd_key := fmt.Sprintf("publishRepositoryDB::GetByCategoryAndPubYear::category=%d,pub_year=%d", category, pub_year)

	rd_json, err := r.rdCache.Get(context.Background(), rd_key).Result()
	if err == nil {
		rd_results := []Publish{}
		if err = json.Unmarshal([]byte(rd_json), &rd_results); err == nil {
			return rd_results, nil
		}
	}

	query_conds := map[string]interface{}{
		"category": category,
		"pub_year": pub_year,
	}

	sql_results := []Publish{}
	if err := r.sqlDB.
		Where(query_conds).Order("id ASC").Find(&sql_results).Error; err != nil {
		return nil, err
	}

	sql_results_bytes, err := json.Marshal(&sql_results)
	if len(sql_results) != 0 && err == nil {
		r.rdCache.Set(context.Background(), rd_key, string(sql_results_bytes), time.Minute*30)
	}

	return sql_results, nil
}

func (r publishRepositoryDB) GetByBibid(bibid string) (*Publish, error) {
	rd_key := fmt.Sprintf("publishRepositoryDB::GetByBibid::bibid=%s", bibid)

	rd_json, err := r.rdCache.Get(context.Background(), rd_key).Result()
	if err == nil {
		rd_result := Publish{}
		if err = json.Unmarshal([]byte(rd_json), &rd_result); err == nil {
			return &rd_result, nil
		}
	}

	query_conds := map[string]interface{}{
		"bibid": bibid,
	}

	sql_result := Publish{}
	if err := r.sqlDB.
		Where(query_conds).Take(&sql_result).Error; err != nil {
		return nil, err
	}

	sql_result_bytes, err := json.Marshal(&sql_result)
	if err == nil {
		r.rdCache.Set(context.Background(), rd_key, string(sql_result_bytes), time.Minute*30)
	}

	return &sql_result, nil
}

func (r publishRepositoryDB) GetPaginateByOptions(page, limit int, options map[string]interface{}) (*PublishPaginate, error) {
	query_conds := map[string]interface{}{
		"category": options["category"],
		"pub_year": options["pub_year"],
	}

	var rows []Publish
	var total_rows int64

	var g errgroup.Group
	g.Go(func() error {
		offset := (page - 1) * limit
		if err := r.sqlDB.
			Where(query_conds).Limit(limit).Offset(offset).Order("id ASC").Find(&rows).Error; err != nil {
			return err
		}

		return nil
	})
	g.Go(func() error {
		if err := r.sqlDB.
			Model(&Publish{}).Where(query_conds).Count(&total_rows).Error; err != nil {
			return err
		}

		return nil
	})
	if err := g.Wait(); err != nil {
		return nil, err
	}

	sql_result := PublishPaginate{
		TotalRows: int(total_rows),
		Sort:      "id ASC",
		Rows:      rows,
	}

	return &sql_result, nil
}

func (r publishRepositoryDB) CreateBatchUpdateOnConflict(publishes []Publish) error {
	rows := []Publish{}
	for _, publish := range publishes {
		rows = append(rows, Publish{
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
		})
	}

	ctx, ctx_cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctx_cancel()

	if err := r.sqlDB.WithContext(ctx).Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: clause.PrimaryKey},
		},
		DoUpdates: clause.AssignmentColumns([]string{
			"category", "text", "bibid", "content_type", "title_th", "title_en", "pub_year", "author_th", "author_en", "link",
		}),
	}).Create(&rows).Error; err != nil {
		return err
	}

	return nil
}
