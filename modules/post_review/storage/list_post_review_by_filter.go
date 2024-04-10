package postreviewstorage

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
	postreviewiomodel "paradise-booking/modules/post_review/iomodel"
)

func (s *postReviewStorage) ListPostReviewByFilter(ctx context.Context, paging *common.Paging, filter *postreviewiomodel.Filter) ([]*entities.PostReview, error) {
	db := s.db

	var data []*entities.PostReview

	db = db.Table(entities.PostReview{}.TableName()).Order("id desc")
	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	if filter.TopicID != 0 {
		db = db.Where("topic = ?", filter.TopicID)
	}

	if err := db.Offset((paging.Page - 1) * paging.Limit).Limit(paging.Limit).Find(&data).Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	return data, nil
}
